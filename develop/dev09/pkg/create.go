package pkg

import (
	"net/url"
	"os"
	"strings"
)

func createFile(fullUrl, fileName string) (string, error) {
	if fileName == "" {
		fileUrl, err := url.Parse(fullUrl)
		if err != nil {
			return "", err
		}
		path := fileUrl.Path
		segments := strings.Split(path, "/")
		fileName = segments[len(segments)-1]
		if fileName == "" {
			fileName = "wget_" + fileUrl.Host
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return fileName, nil
}
