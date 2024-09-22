package config

import (
	"flag"
	"fmt"
	"log"
)

// Config - структура, хранящая параметры, переданные пользователем при запуске приложения.
type Config struct {
	Pattern  string
	Filename string
	After    int
	Before   int
	Context  int
	Count    bool
	Line     bool
	Invert   bool
	Fixed    bool
	IgnCase  bool
}

// Validate проверяет переданные параметры на совместимость. Возвращает ошибку, если параметры
// противоречат друг-другу либо не могут быть интерпретированы.
func (c *Config) Validate() error {
	switch {
	case c.After > 0 && (c.Count || c.Before > 0 || c.Context > 0 || c.Invert):
		return fmt.Errorf("flag -A is incombatible with any of the -v, -B, -C, or -c flags")
	case c.Before > 0 && (c.Count || c.After > 0 || c.Context > 0 || c.Invert):
		return fmt.Errorf("flag -B is incombatible with any of the -v, A, -C, or -c flags")
	case c.Count && (c.After > 0 || c.Before > 0 || c.Context > 0 || c.Invert || c.Line):
		return fmt.Errorf("flag -c is incombatible with any of the -n, -v, -A, -B, or -C flags")
	case c.Context > 0 && (c.After > 0 || c.Before > 0 || c.Invert || c.Count):
		return fmt.Errorf("flag -C is incombatible with any of the -v, -A, -B or -c flags")
	case c.Invert && (c.After > 0 || c.Before > 0 || c.Context > 0 || c.Count):
		return fmt.Errorf("flag -v is incombatible with any of the -A, -B, -C, or -c flags")
	}
	return nil
}

// MustLoad превращает аргументы командной строки в структуру Config. Если параметров недостаточно, приложение
// завершается, информация об ошибке записывается в stderr.
func MustLoad() Config {
	var cfg Config

	flag.StringVar(&cfg.Filename, "f", "", "name of the file to process")
	flag.StringVar(&cfg.Pattern, "p", "", "pattern to match")

	flag.IntVar(&cfg.After, "A", 0, "to print strings after the match")
	flag.IntVar(&cfg.Before, "B", 0, "to print strings before the match")
	flag.IntVar(&cfg.Context, "C", 0, "to print strings around the match")

	flag.BoolVar(&cfg.IgnCase, "i", false, "to ignore case")
	flag.BoolVar(&cfg.Invert, "v", false, "exclude instead")
	flag.BoolVar(&cfg.Line, "n", false, "to print number of string")
	flag.BoolVar(&cfg.Count, "c", false, "number of strings to print")
	flag.BoolVar(&cfg.Fixed, "F", false, "to use exact pattern, not regular expression")

	flag.Parse()

	if cfg.Filename == "" || cfg.Pattern == "" {
		log.Fatal("Both pattern and filename must be specified")
	}

	return cfg
}
