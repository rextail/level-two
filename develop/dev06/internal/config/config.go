package config

import (
	"flag"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type Config struct {
	Fields    []int
	Delimiter string
	Separated bool
}

// parseFields преобразует столбы из строкового вида в срез целых чисел
func (c Config) parseFields(input string) ([]int, error) {
	splitted := strings.Split(input, ",")
	if len(splitted) == 0 {
		return nil, fmt.Errorf("error: empty input")
	}

	fields := make([]int, 0, len(splitted))
	for _, s := range splitted {
		number, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("can't use %s as fields numbers: %w", input, err)
		}
		fields = append(fields, number)
	}
	slices.Sort(fields)
	return fields, nil

}

// MustParseConfig преобразует параметры, указанные при запуске программы, в Config. Если параметры не могут быть
// интерпретированы, приложение завершается.
func MustParseConfig() Config {
	var cfg Config

	var fieldsInput string

	flag.StringVar(&fieldsInput, "f", "", "fields to return after splitting by delimiter")
	flag.StringVar(&cfg.Delimiter, "d", "\t", "delimiter to separate by")
	flag.BoolVar(&cfg.Separated, "s", false, "to give out only rows containing delimiter")

	flag.Parse()
	fields, err := cfg.parseFields(fieldsInput)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	cfg.Fields = fields

	return cfg
}
