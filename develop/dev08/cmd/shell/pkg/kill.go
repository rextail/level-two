package pkg

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/process"
)

func killCommand(args []string) (string, error) {
	// проверим, что есть , что убивать
	if len(args) < 2 {
		return "", errors.New("shell: kill command needs name of process")
	}
	processes, err := process.Processes()
	if err != nil {
		return "", err
	}
	for _, process := range processes {
		name, err := process.Name()
		if err != nil {
			return "", err
		}
		if name == args[1] {
			if err := process.Kill(); err != nil {
				return "", errors.New("shell: this process can not be killed")
			}
		}
	}
	return fmt.Sprintf("Process %v successfully killed", args[1]), nil
}
