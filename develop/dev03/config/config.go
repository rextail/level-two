package config

import (
	"flag"
	"fmt"
	"os"
)

// InputFormatOptions параметры, которые относятся к форматированию данных из файла
type InputFormatOptions struct {
	IgnoreEndingSpaces bool
}

// SortingOptions параметры, которые относятся к сортировке данных
type SortingOptions struct {
	Column       int
	Unique       bool
	Reverse      bool
	Numeric      bool
	ByMonth      bool
	HumanNumeric bool
	CheckSorted  bool
}

// Config  хранит в себе все опции, переданные пользователем через командную строку
type Config struct {
	Filename  string
	InputOpts InputFormatOptions
	SortOpts  SortingOptions
}

// MustParseConfig превращает данные, переданные пользователем, в Config. Если данные ошибочные, выполнение прекращается.
func MustParseConfig() *Config {
	cfg := &Config{}

	flag.BoolVar(&cfg.InputOpts.IgnoreEndingSpaces, "b", false, "ignore ending spaces")

	flag.BoolVar(&cfg.SortOpts.Unique, "u", false, "only unique values are included extractor the result")
	flag.IntVar(&cfg.SortOpts.Column, "k", 0, "column to sort")
	flag.BoolVar(&cfg.SortOpts.Reverse, "r", false, "descending sort")
	flag.BoolVar(&cfg.SortOpts.Numeric, "n", false, "sort by numeric value")
	flag.BoolVar(&cfg.SortOpts.HumanNumeric, "h", false, "sort with num suffix (15K, 10M, 1G...)")
	flag.BoolVar(&cfg.SortOpts.ByMonth, "M", false, "sort by short month name (JAN, FEB...)")
	flag.BoolVar(&cfg.SortOpts.CheckSorted, "c", false, "check if input data is already sorted")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Error: file name is not specified")
		os.Exit(1)
	}
	cfg.Filename = args[0]

	return cfg
}

func (c *Config) Validate() error {
	if moreThanOneOpt(c.SortOpts.CheckSorted,
		c.SortOpts.ByMonth,
		c.SortOpts.Numeric,
		c.SortOpts.HumanNumeric) {
		return fmt.Errorf("can't sort by more than one type")
	}
	if c.SortOpts.Column < 0 {
		return fmt.Errorf("column index must be greater than or equal to zero")
	}
	return nil
}

func moreThanOneOpt(flags ...bool) bool {
	count := 0
	for _, f := range flags {
		if f == true {
			count++
		}
		if count > 1 {
			return true
		}
	}
	return false
}
