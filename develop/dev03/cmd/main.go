package main

import (
	"dev03/internal/config"
	"dev03/internal/format"
	"dev03/internal/print"
	"dev03/internal/sort"
	"fmt"
	"log"
)

const path = "develop/dev03/pkg"

func main() {
	cfg := config.MustParseConfig()

	log.Println("initialized config")

	table, err := format.ToTable(cfg.Filename, cfg.InputOpts)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	log.Println("transformed input file to the table")

	res := sort.SortTable(table, cfg.SortOpts)
	if res.Err != nil {
		log.Fatalf("%s", err.Error())
	}

	fmt.Println("sorting finished")

	err = print.WriteResult(cfg.Filename, res)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

}
