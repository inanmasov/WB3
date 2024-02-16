package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Флаги командной строки
	fieldsStr := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	onlySeparated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// Разбор запрошенных колонок
	fields := parseFields(*fieldsStr)

	// Сканер для чтения данных из STDIN
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if *onlySeparated && !strings.Contains(line, *delimiter) {
			continue
		}
		columns := strings.Split(line, *delimiter)
		output := make([]string, 0, len(fields))
		for _, field := range fields {
			if field > 0 && field <= len(columns) {
				output = append(output, columns[field-1])
			} else {
				output = append(output, "")
			}
		}
		fmt.Println(strings.Join(output, *delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении данных из STDIN:", err)
		os.Exit(1)
	}
}

// parseFields разбирает строку с указанными полями и возвращает их в виде среза
func parseFields(fieldsStr string) []int {
	fields := strings.Split(fieldsStr, ",")
	result := make([]int, 0, len(fields))
	for _, field := range fields {
		if field == "" {
			continue
		}
		if strings.Contains(field, "-") {
			rangeParts := strings.Split(field, "-")
			start := parseField(rangeParts[0])
			end := parseField(rangeParts[1])
			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else {
			result = append(result, parseField(field))
		}
	}
	return result
}

// parseField преобразует строку в число
func parseField(fieldStr string) int {
	fieldStr = strings.TrimSpace(fieldStr)
	if fieldStr == "" {
		return 0
	}
	field, err := strconv.Atoi(fieldStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при парсинге поля: %v\n", err)
		os.Exit(1)
	}
	return field
}
