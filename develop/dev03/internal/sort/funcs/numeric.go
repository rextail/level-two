package funcs

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

// NumericSort сортирует входящий массив по следующим правилам:
// данные, которые являются числами, сортируются соответственно и помещаются перед строками, начинающихся с чисел.
// Последовательность строк, начинающихся с чисел, сортируется по лидирующим числам, а затем и по строковой части,
// если численные части равны. Затем размещается последовательность строк, которая не начинается с числа,
// отсортированная в лексикографическом порядке
func NumericSort(strs []string) {
	//разделим данные на три вида:
	var digits []string            //строки, в которой только числа
	var letters []string           //строки, состоящая только из символов
	var digitsWithLetters []string //строки, начинающаяся с чисел

	//Распределим строки по группам
	for _, str := range strs {
		runed := []rune(str)

		if len(str) == 0 {
			letters = append(letters, str)
		}

		if len(str) >= 2 {
			if !unicode.IsDigit(runed[0]) {
				//Если первый символ не число, то считаем за обычный набор букв
				letters = append(letters, str)
				continue
			} else {
				nonDigit := false
				for _, r := range runed {
					//Если первый символ число, пройдемся по всей строке в поиске символа, не являющегося числом
					if !unicode.IsDigit(r) {
						nonDigit = true
						digitsWithLetters = append(digitsWithLetters, str)
						break
					}
				}
				//если строка не содержит ничего кроме чисел
				if !nonDigit {
					digits = append(digits, str)
				}
			}
		}
	}
	sortDigits(digits)
	sortDigitsWithLetters(digitsWithLetters)
	LexicSort(letters)

	index := 0

	for _, d := range digits {
		strs[index] = d
		index++
	}
	for _, d := range digitsWithLetters {
		strs[index] = d
		index++
	}
	for _, l := range letters {
		strs[index] = l
		index++
	}

}

func sortDigits(strs []string) {
	//воспользуемся пакетом slices и определим функцию, в которой преобразуем строки к числам и сравним
	slices.SortFunc(strs, func(a, b string) int {
		x, _ := strconv.Atoi(a)
		y, _ := strconv.Atoi(b)
		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
		return 0
	})
}

func sortDigitsWithLetters(strs []string) {
	slices.SortFunc(strs, func(a, b string) int {
		re := regexp.MustCompile(`^(\d+)(.*)$`)

		matchesA := re.FindStringSubmatch(a)
		matchesB := re.FindStringSubmatch(b)

		if len(matchesA) < 3 || len(matchesB) < 3 {
			// Если строка не соответствует формату, возвращаем лексикографическое сравнение
			return strings.Compare(a, b)
		}

		aDigitPart, aStringPart := matchesA[1], matchesA[2]
		bDigitPart, bStringPart := matchesB[1], matchesB[2]

		x, _ := strconv.Atoi(aDigitPart)
		y, _ := strconv.Atoi(bDigitPart)

		if x < y {
			return -1
		}
		if x > y {
			return 1
		}
		return strings.Compare(aStringPart, bStringPart)
	})
}
