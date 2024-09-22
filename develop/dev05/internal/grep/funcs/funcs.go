package funcs

// RangeFunc возвращает интервал. Логика определения интервала зависит от реализации
type RangeFunc func(index, count int) [2]int
