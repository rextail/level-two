package pkg

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"strings"
)

func psCommand() (string, error) {
	processes, err := ps.Processes()
	if err != nil {
		return "", err
	}
	var builder strings.Builder
	builder.WriteString("PID\t|\tCOMMAND\n")
	builder.WriteString("---------------\n")
	for _, proc := range processes {
		builder.WriteString(
			fmt.Sprintf("%v\t|\t%v\n", proc.Pid(), proc.Executable()),
		)
	}
	builder.WriteString("---------------\n")
	return builder.String(), nil
}
