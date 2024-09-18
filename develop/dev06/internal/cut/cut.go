package cut

import (
	"dev06/internal/config"
	"strings"
)

// ExtractFields возвращает выборку из строк, полученную путем разбиения исходной строки по заданному разделителю.
// Выборка осуществляется по столбцам, указанных при запуске программы.
func ExtractFields(strs []string, cfg config.Config) ([]string, error) {
	var res []string
	for _, str := range strs {
		splitted := strings.Split(str, cfg.Delimiter)
		//для флага -s после разбиения строки по сепаратору, количество частей должно быть больше 1
		if (cfg.Separated && len(splitted) > 1) || !cfg.Separated {
			for _, field := range cfg.Fields {
				// Проверяем, что индекс поля не выходит за границы разбиения
				if field < len(splitted) {
					// Добавляем выбранное поле к результату
					res = append(res, splitted[field])
				}
			}
		}
	}
	return res, nil
}
