package format

import (
	"bufio"
	"dev03/internal/config"
	"fmt"
	"os"
	"strings"
)

// FormatString возвращает строку, отформатированную в соответствие с параметрами
func FormatString(str string, opts config.InputFormatOptions) string {
	if opts.IgnoreEndingSpaces {
		str = strings.TrimRight(str, " ")
	}
	//Здесь могла быть ваша функция форматирования
	return str
}

// ToTable приводит данные из файла к табличному виду, используя пробел в качестве разделителя
func ToTable(filename string, opts config.InputFormatOptions) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		//если не смогли открыть, вернем ошибку
		return nil, fmt.Errorf("can't open file: %v", err)
	}
	defer file.Close()

	var table [][]string

	//воспользуемся сканером для чтения построчно из файла
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() //получаем строку

		str := FormatString(line, opts) //приводим ее к необходимому формату

		splitted := strings.Split(str, " ")

		table = append(table, splitted)
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while reading file %v", err)
	}

	return table, nil
}
