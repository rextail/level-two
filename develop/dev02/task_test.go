package dev02

import "testing"

func TestUnzip(t *testing.T) {
	testcases := []struct {
		input    string
		expected string
		hasError bool
	}{
		// Примитивная распаковка строк
		{"a4bc2d5e", "aaaabccddddde", false}, // обычная строка с повторениями
		{"abcde", "abcde", false},            // строка без повторений
		{"", "", false},                      // пустая строка
		{"4abcde", "", true},                 // некорректная строка (начинается с цифры)

		// Проверка escape-последовательностей
		{"qwe\\45a", "qwe44444a", false},     // символ 4 повторяется 5 раз
		{"qwe\\\\5", "qwe\\\\\\\\\\", false}, // символ \ повторяется 5 раз
		{"qwe\\4\\5", "qwe45", false},        // символы \4 и \5 добавлены по одному разу
		{"\\a5", "aaaaa", false},             // символ a повторяется 5 раз

		// Ошибочные случаи
		{"\\", "", true},    // строка заканчивается на \
		{"abc\\", "", true}, // строка заканчивается на \
		{"qwe\\", "", true}, // строка заканчивается на \
	}

	for _, tc := range testcases {
		got, err := Unzip(tc.input)
		if tc.hasError {
			if err == nil {
				t.Errorf("expected error for input %s, got none", tc.input)
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error for input %s: %v", tc.input, err)
			}
			if got != tc.expected {
				t.Errorf("for input %s, got %s, expected %s", tc.input, got, tc.expected)
			}
		}
	}
}
