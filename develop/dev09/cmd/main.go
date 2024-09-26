package main

import (
	"dev08/pkg"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("input more data")
		return
	}

	if len(os.Args) > 3 {
		println("extra data")
		return
	}

	flags := os.Args[1:]
	url := flags[0]
	fileName := ""
	if len(flags) > 1 {
		fileName = flags[1]
	}

	wgeter, err := pkg.InitWgeter(url, fileName)
	if err != nil {
		println(err.Error())
		return
	}

	err = wgeter.Start()
	if err != nil {
		println(err.Error())
	}
}
