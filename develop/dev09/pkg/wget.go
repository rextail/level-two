package pkg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Wgeter struct {
	url      string
	fileName string
}

func InitWgeter(url string, filename string) (*Wgeter, error) {
	if url == "" {
		return nil, errors.New("empty text")
	}
	return &Wgeter{url, filename}, nil
}

func (w *Wgeter) Start() error {
	fileName, err := createFile(w.url, w.fileName)
	if err != nil {
		return err
	}
	w.fileName = fileName

	resp, err := getResponse(w.url)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(w.fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Println("Downloaded a file " + w.fileName + " with size " + strconv.FormatInt(size, 10))
	return nil
}
