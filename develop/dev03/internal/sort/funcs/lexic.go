package funcs

import "slices"

// LexicSort сортирует входящий массив в лексикографическом порядке
func LexicSort(strs []string) {
	slices.Sort(strs)
}
