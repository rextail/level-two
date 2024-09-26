package pkg

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func forkCommand(args []string) (string, error) {
	if len(args) < 2 {
		return "", errors.New("shell: fork command needs an amount of processes")
	}

	fork, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	pid := os.Getpid()
	ppid := os.Getppid()
	builder.WriteString(fmt.Sprintf("pid: %d, ppid: %d, forks: %d\n", pid, ppid, fork))

	children := []int{}

	for i := 0; i < fork; i++ {
		childENV := []string{fmt.Sprintf("CHILD_ID=%d", i)}
		pwd, err := os.Getwd()
		if err != nil {
			return "", err
		}

		// Создаем аргументы для дочернего процесса
		childArgs := append(os.Args, fmt.Sprintf("#child_%d_of_%d", i, pid))

		// Вызываем syscall.ForkExec для создания нового процесса
		childPID, err := syscall.ForkExec(childArgs[0], childArgs, &syscall.ProcAttr{
			Dir: pwd,
			Env: append(os.Environ(), childENV...),
			Sys: &syscall.SysProcAttr{
				Setsid: true,
			},
			Files: []uintptr{0, 1, 2}, // перенаправляем stdin, stdout, stderr
		})

		if err != nil {
			return "", fmt.Errorf("failed to fork child: %w", err)
		}

		builder.WriteString(fmt.Sprintf("parent %d forked child %d\n", pid, childPID))
		children = append(children, childPID)
	}

	// Сохраняем идентификаторы дочерних процессов в окружении
	if len(children) > 0 {
		os.Setenv("CHILDREN", fmt.Sprintf("%v", children))
	}

	return builder.String(), nil
}
(os.Args[0], append(os.Args, fmt.Sprintf("#child_%d_of_%d", i, pid)), &unix.ProcAttr{
				Dir: pwd,
				Env: append(os.Environ(), childENV...),
				Sys: &unix.SysProcAttr{
					Setsid: true,
				},
				Files: []uintptr{0, 1, 2}, // print message to the same pty
			})
			if err != nil {
				return "", err
			}

			builder.WriteString(fmt.Sprintf("parent %d fork %d\n", pid, childPID))
			if childPID != 0 {
				children = append(children, childPID)
			}
		}

		// print children
		builder.WriteString(fmt.Sprintf("parent: PID=%d children=%v\n", pid, children))
		if len(children) == 0 && fork != 0 {
			return "", errors.New("shell: no child available, exit")
		}

		// set env
		for _, childID := range children {
			if c := os.Getenv("CHILDREN"); c != "" {
				os.Setenv("CHILDREN", fmt.Sprintf("%s,%d", c, childID))
			} else {
				os.Setenv("CHILDREN", fmt.Sprintf("%d", childID))
			}
		}
	}
	return builder.String(), nil
}
