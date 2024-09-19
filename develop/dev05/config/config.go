package config

import (
	"flag"
	"log"
)

type Config struct {
	Expression string
	Filename   string
	Count      int
	Line       int
	Invert     bool
	Fixed      bool
	After      bool
	Before     bool
	Context    bool
	IgnCase    bool
}

func MustLoad() Config {
	var cfg Config

	flag.IntVar(&cfg.Line, "n", 0, "to print number of string")
	flag.IntVar(&cfg.Count, "c", 0, "number of strings to print")

	flag.BoolVar(&cfg.After, "A", false, "to print strings after the match")
	flag.BoolVar(&cfg.Before, "B", false, "to print strings before the match")
	flag.BoolVar(&cfg.Context, "C", false, "to print strings around the match")
	flag.BoolVar(&cfg.IgnCase, "i", false, "to ignore case")
	flag.BoolVar(&cfg.Fixed, "F", false, "to use exact pattern, not regular expression")
	flag.BoolVar(&cfg.Invert, "v", false, "exclude instead")

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		log.Fatalf("not enough data to execute command")
	}

	cfg.Filename = args[0]
	cfg.Expression = args[1]

	return cfg
}
