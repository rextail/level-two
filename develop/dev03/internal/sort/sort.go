package sort

import (
	"dev03/internal/config"
	"dev03/internal/sort/funcs"
	"fmt"
	"slices"
)

/*
	Сама по себе sort, судя по документации, пытается применить лексикографическую сортировку,
	пока мы не попросим ее об обратном. Если мы просим ее отсортировать значения, как если бы это были
	месяца или числа, то данные, которые не являются месяцами или числами, будут считаться набольшими
	и сортироваться в лексикографическом порядке.

	Варианты:
		1. Лексикографическая сортировка при -n false, -M false, -h false
		2. Сортировка при -n true осуществляется следующим образом
			2.1. Сначала сортируем по числам, с которых начинается строка
			2.2. Затем, если лидирующие числа одинаковые, сортируем лексикографически
			2.3. Если строка начинается не с числа, то она считается больше любого начинающегося с числа,
				 но затем все равно сортируется лексикографически
		3.  Строки без числовых значений или суффиксов при -h true воспринимаются как минимальные,
			поэтому они помещаются в начало.
		4. Аналогично с п.3 для сортировки по месяцам, все, что не является месяцем, сортируется лексикографически
			и помещается в начало.
*/

type SortResult struct {
	Column   []string
	IsSorted bool
	Err      error
}

func newSortResult(column []string, isSorted bool, err error) SortResult {
	return SortResult{
		Column:   column,
		IsSorted: isSorted,
		Err:      err,
	}
}

func SortTable(tab [][]string, opts config.SortingOptions) SortResult {
	col := make([]string, 0, len(tab))
	for i, t := range tab {
		if len(t) < opts.Column {
			return newSortResult(nil, false, fmt.Errorf(
				"can't sort for given column: %d, row %d has only %d columns",
				opts.Column, i, len(t),
			))
		}
		col = append(col, t[opts.Column])
	}

	if opts.CheckSorted {
		return newSortResult(nil, funcs.IsSorted(col), nil)
	}

	switch {
	case opts.Numeric:
		funcs.NumericSort(col)
	case opts.HumanNumeric:
		funcs.SortByHumanNumeric(col)
	case opts.ByMonth:
		funcs.SortByMonth(col)
	default:
		funcs.LexicSort(col)
	}

	if opts.Reverse {
		slices.Reverse(col)
	}

	if opts.Unique {
		used := make(map[string]struct{})
		uniqueCol := make([]string, 0, len(col))
		for _, c := range col {
			if _, ok := used[c]; !ok {
				uniqueCol = append(uniqueCol, c)
				used[c] = struct{}{}
			}
		}
		return newSortResult(uniqueCol, false, nil)
	}

	return newSortResult(col, false, nil)
}
