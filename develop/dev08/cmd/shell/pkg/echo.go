package pkg

import (
	"errors"
)

func echoCommand(args []string) (string, error) {
	if len(args) == 1 {
		return "", errors.New("shell: echo should have some data")
	}
	return args[1], nil
}
