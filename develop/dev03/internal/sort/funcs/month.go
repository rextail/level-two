package funcs

import (
	"slices"
)

// SortByMonth сортирует входящую строку по следующим правилам:
// все, что не является сокращенной записью месяца, ака JAN, FEB etc. сортируется в лексикографическом
// порядке и помещается перед отсортированными месяцами. Месяцы сортируются по их порядку в году.
func SortByMonth(strs []string) {
	month := map[string]byte{
		"JAN": 1, "FEB": 2, "MAR": 3, "APR": 4,
		"MAY": 5, "JUN": 6, "JUL": 7, "AUG": 8,
		"SEP": 9, "OCT": 10, "NOV": 11, "DEC": 12,
	}
	var monthes []string
	var otherData []string

	for _, str := range strs {
		if _, ok := month[str]; !ok {
			otherData = append(otherData, str)
		} else {
			monthes = append(monthes, str)
		}
	}

	LexicSort(otherData)

	slices.SortFunc(monthes, func(a, b string) int {
		x := month[a]
		y := month[b]
		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
		return 0
	})

	index := 0

	//сначала идут данные, которые не являются месяцами
	for _, o := range otherData {
		strs[index] = o
		index++
	}
	//после чего идут месяца
	for _, m := range monthes {
		strs[index] = m
		index++
	}
}
