package funcs

import (
	"slices"
	"strconv"
	"unicode"
)

// SortByHumanNumeric сортирует входящий массив по следующим правилам:
// числа с суффиксами приводятся к обычным числам и сортируется по тем же правилам, что и обычные числа.
// Строки, которые не удалось преобразовать к числу, сортируются лексикографически и помещаются перед числами
func SortByHumanNumeric(strs []string) {
	suffix := map[rune]int{
		'K': 1 << 10,
		'M': 1 << 20,
		'G': 1 << 30,
	}

	var digits []string
	var otherData []string

	for _, str := range strs {
		runed := []rune(str)

		if len(runed) == 0 {
			otherData = append(otherData, str)
			continue
		}

		for i, r := range runed {
			//если до последнего символа встретился не числовой символ
			if i < len(runed)-1 && !unicode.IsDigit(r) {
				otherData = append(otherData, str)
				break
			}
			//Если до этого все символы были числовыми, и последний символ либо суффикс, либо число
			if i == len(runed)-1 {
				_, ok := suffix[r]
				if ok || unicode.IsDigit(runed[i]) {
					digits = append(digits, str)
				} else {
					otherData = append(otherData, str)
				}
			}
		}
	}

	LexicSort(otherData)

	slices.SortFunc(digits, func(a, b string) int {
		convert := func(s string) int {
			runed := []rune(s)
			length := len(runed)
			if length == 1 && unicode.IsDigit(runed[0]) {
				//если число состоит из одного элемента, которое не является числом, то это суффикс
				return suffix[runed[length-1]]
			}
			// Если строка оканчивается суффиксом
			if !unicode.IsDigit(runed[length-1]) {
				// Проверяем, есть ли числовая часть перед суффиксом
				num, _ := strconv.Atoi(string(runed[:length-1]))

				return num * suffix[runed[length-1]]
			}

			num, _ := strconv.Atoi(s)
			return num
		}

		x := convert(a)
		y := convert(b)

		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
		return 0
	})
}
