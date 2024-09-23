package extractor

import (
	"bufio"
	"dev05/config"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// Extractor содержит параметры, определяющие заданное выражение
type Extractor struct {
	expression    string
	invert        bool
	isRegular     bool
	caseSensitive bool
}

func New(cfg config.Config) *Extractor {
	return &Extractor{
		expression:    cfg.Pattern,
		invert:        cfg.Invert,
		isRegular:     !cfg.Fixed,
		caseSensitive: !cfg.IgnCase,
	}
}

// ExtractLines возвращает строки по указанным индексам
func (e *Extractor) ExtractLines(data io.Reader, indexes []int) ([]string, error) {

	if len(indexes) == 0 {
		return []string{}, nil
	}

	var lines []string

	i := 0
	count := 0
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		row := scanner.Text()

		if i == indexes[count] && count < len(indexes)-1 {
			lines = append(lines, row)
			count++
		}
		i++
	}
	if scanner.Err() != nil {
		return nil, fmt.Errorf("error reading data: %w", scanner.Err())
	}

	return lines, nil
}

// ExtractIndexes возвращает порядковые номера строк, в которых было найдено совпадение с выражением.
func (e *Extractor) ExtractIndexes(data io.Reader) ([]int, error) {
	var indexes []int

	//приведем к нижнему регистру, если выбран флаг регистронезависимости
	if !e.caseSensitive {
		e.expression = strings.ToLower(e.expression)
	}

	scanner := bufio.NewScanner(data)

	reg := regexp.MustCompile(e.expression)

	//пройдемся построчно тексту, занесем номера строк с совпадениями в indexes
	index := 0

	for scanner.Scan() {
		match := false
		row := scanner.Text()
		if e.isRegular {
			ind := reg.FindStringIndex(row)
			match = ind != nil
		} else {
			match = strings.Contains(row, e.expression)
		}
		if e.invert {
			match = !match
		}
		if match {
			indexes = append(indexes, index)
		}
		index++
	}
	if scanner.Err() != nil {
		return nil, fmt.Errorf("error reading data: %w", scanner.Err())
	}

	return indexes, nil
}
