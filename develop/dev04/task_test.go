package dev04

import (
	"reflect"
	"testing"
)

func TestFindAllAnagramms(t *testing.T) {
	words := []string{"Пятак", "пятка", "тяпка", "листок", "столик", "слиток"}
	got := FindAllAnagramms(words)
	expected := map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}, "листок": {"листок", "слиток", "столик"}}
	if reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
