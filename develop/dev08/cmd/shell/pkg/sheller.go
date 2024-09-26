package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"
)

type Sheller struct {
	Args   []string
	reader io.Reader
	writer io.Writer
}

func InitSheller(reader io.Reader, writer io.Writer) (*Sheller, error) {
	return &Sheller{reader: reader, writer: writer}, nil
}

func (c *Sheller) GetText() string {
	return ""
}

func (s *Sheller) Start() error {
	fmt.Fprintln(s.writer, `MYSHELL. For exit input: \exit`)
	scanner := bufio.NewScanner(s.reader)
	for {
		prefix, err := s.buildPrefix()
		if err != nil {
			return errors.New("myshell: can not check current directory place")
		}
		fmt.Fprint(s.writer, prefix)
		scanner.Scan()
		text := scanner.Text()
		if text == `\exit` {
			break
		}

		args := strings.Fields(text)
		s.Args = args

		var res string
		switch args[0] {
		case "cd":
			res, err = cdCommand(s.Args)
		case "echo":
			res, err = echoCommand(s.Args)
		case "ps":
			res, err = psCommand()
		case "pwd":
			res, err = pwdCommand()
		case "kill":
			res, err = killCommand(s.Args)
		case "fork":
			res, err = forkCommand(s.Args)
		default:
			fmt.Fprintln(s.writer, "shell: unknown command")
		}

		if err != nil {
			fmt.Fprintln(s.writer, err.Error())
			continue
		}
		fmt.Fprintln(s.writer, res)
	}
	if scanner.Err() != nil {
		return errors.New("shell: can not read data")
	}
	if _, err := fmt.Fprintln(s.writer, "shell: success exit. Bye."); err != nil {
		return err
	}
	return nil
}

func (s *Sheller) buildPrefix() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	var postfix string
	userName, err := user.Current()
	if err != nil {
		return "", errors.New("shell: can not get current user info")
	}
	if path == "/home/"+userName.Name {
		postfix = "$ "
	} else {
		postfix = " " + path + " "
	}
	return `rextail@myshell:~` + postfix, nil
}
