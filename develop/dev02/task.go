package dev02

import (
	"errors"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var ErrorIncorrectInput = errors.New("input string is not correct")

func Unzip(str string) (string, error) {
	runed := []rune(str)

	res := make([]rune, 0, len(runed))

	var last rune

	fill := func(val rune, times int) {
		for i := 0; i < times; i++ {
			res = append(res, val)
		}
	}

	for i := 0; i < len(runed); i++ {
		if runed[i] == '\\' {
			i++
			if i == len(runed) {
				//если символ последний, то для операции недостаточно данных
				return "", ErrorIncorrectInput
			}
			if i+1 < len(runed) && unicode.IsDigit(runed[i+1]) {
				//1 случай, когда у нас задано число повторений [\45] -> repeat 4 for 5 times
				repeats, _ := strconv.Atoi(string(runed[i+1]))
				fill(runed[i], repeats)
				i++
			} else {
				//2 случай - когда у нас не задано число повторений, принимаем за 1 [\4] -> repeat 4 one time
				res = append(res, runed[i])
			}
			continue
		}
		if !unicode.IsDigit(runed[i]) {
			last = runed[i]
			res = append(res, runed[i])
		} else {
			if last == 0 {
				//если предыдущее число - цифра
				return "", ErrorIncorrectInput
			}
			repeats, _ := strconv.Atoi(string(runed[i])) //получаем число из руны
			fill(last, repeats-1)
			last = 0
		}
	}
	return string(res), nil
}
