package main

import (
	"dev05/config"
	"dev05/internal/extractor"
	"dev05/internal/grep"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	cfg := config.MustLoad()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("config validation failed: %v", err)
	}

	extr := extractor.New(cfg)

	file, err := os.Open(cfg.Filename)
	if err != nil {
		log.Fatalf("can't open file: name %s, %v", cfg.Filename, err)
	}
	defer file.Close()

	indexes, err := extr.ExtractIndexes(file)
	if err != nil {
		log.Fatalf("failed to extract indexes: %v", err)
	}

	if cfg.After > 0 || cfg.Before > 0 || cfg.Context > 0 {
		processABC(cfg, file, extr, indexes)
	}

	if cfg.Invert {
		processInvert(file, extr, indexes)
	}

	if cfg.Count {
		fmt.Println(len(indexes))
	}

	if cfg.Line {
		fmt.Println(indexes)
	}
}

// processABC обрабатывает флаги -A, -B, -C и выводит результат
func processABC(cfg config.Config, file *os.File, extr *extractor.Extractor, indexes []int) {
	grepper := grep.New(cfg)
	ranges := grepper.GrepABC(indexes)
	resultIndexes := unpackRanges(ranges)

	setPointerToStart(file)
	result, err := extr.ExtractLines(file, resultIndexes)
	if err != nil {
		log.Fatalf("failed to extract lines with context: %v", err)
	}
	fmt.Println(strings.Join(result, "\n"))
}

// processInvert обрабатывает флаг -v (invert) и выводит строки, не содержащие совпадений
func processInvert(file *os.File, extr *extractor.Extractor, indexes []int) {
	setPointerToStart(file)
	result, err := extr.ExtractLines(file, indexes)
	if err != nil {
		log.Fatalf("failed to extract lines with invert: %v", err)
	}
	fmt.Println(strings.Join(result, "\n"))
}

// unpackRanges разворачивает диапазоны индексов в одномерный массив индексов
func unpackRanges(ranges [][2]int) []int {
	var res []int
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			res = append(res, i)
		}
	}
	return res
}

// setPointerToStart сбрасывает указатель файла на начало
func setPointerToStart(file *os.File) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatalf("failed to seek to the start of file: %v", err)
	}
}
