package main

import (
	"bufio"
	"dev06/internal/config"
	"dev06/internal/cut"
	"fmt"
	"os"
	"strings"
)

func main() {
	cfg := config.MustParseConfig()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			rows := strings.Split(input, "\n")
			fields, err := cut.ExtractFields(rows, cfg)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			fmt.Println(fields)
		}

	}
}
