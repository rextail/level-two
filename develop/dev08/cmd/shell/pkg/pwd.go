package pkg

import (
	"os"
)

func pwdCommand() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return "Current work dir :" + path, err
}
