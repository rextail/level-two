package funcs

import "slices"

// IsSorted проверяет, отсортирован ли лексикографически заданный массив
func IsSorted(strs []string) bool {
	return slices.IsSorted(strs)
}
