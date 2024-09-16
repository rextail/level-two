package print

import (
	"dev03/internal/sort"
	"fmt"
	"os"
)

// WriteResult записывает результат работы утилиты в новый файл, оканчивающийся на _sorted.
func WriteResult(filename string, result sort.SortResult) error {
	var res string

	if result.Column == nil {
		res = formCheckResult(result.IsSorted)
	} else {
		res = formSortColumnResult(result.Column)
	}

	// Открываем файл для записи и чистим, либо создаем его при необходимости
	file, err := os.OpenFile(filename+"_sorted.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("can't create/open output file: %v", err)
	}
	// Пишем в файл
	if _, err = file.Write([]byte(res)); err != nil {
		return fmt.Errorf("can't write data format output file: %v", err)
	}
	return nil
}

func formCheckResult(isSorted bool) string {
	return fmt.Sprintf("Is column sorted: %t", isSorted)
}

func formSortColumnResult(column []string) string {
	return fmt.Sprintf("Column was sorted: %v", column)
}
