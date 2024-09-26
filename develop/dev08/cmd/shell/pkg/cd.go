package pkg

import (
	"os"
	"os/user"
)

func cdCommand(args []string) (string, error) {
	if len(args) == 1 {
		cUser, err := user.Current()
		if err != nil {
			return "", err
		}
		if err := os.Chdir(cUser.HomeDir); err != nil {
			return "", err
		}
	} else {
		if err := os.Chdir(args[1]); err != nil {
			return "", err
		}

	}

	return "Successfully changed dir", nil
}
