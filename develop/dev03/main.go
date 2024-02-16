// go run main.go -i input.txt -o output.txt -k 2 -n -r
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Флаги командной строки
	key := flag.String("k", "", "указание колонки для сортировки")
	numeric := flag.Bool("n", false, "сортировать по числовому значению")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	sortByMonth := flag.Bool("M", false, "сортировать по названию месяца")
	ignoreTrailingSpace := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	checkSorted := flag.Bool("c", false, "проверять отсортированы ли данные")
	numericWithSuffix := flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")
	inputFileName := flag.String("i", "", "имя файла для вывода результатов")
	outputFileName := flag.String("o", "", "имя файла для вывода результатов")
	flag.Parse()

	// Открытие файла
	inputFile, err := os.Open(*inputFileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка открытия файла:", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	// Чтение строк из файла
	scanner := bufio.NewScanner(inputFile)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения файла:", err)
		os.Exit(1)
	}

	// Функция сравнения строк в зависимости от установленных флагов
	compare := func(i, j int) bool {
		if *numeric {
			a, _ := strconv.ParseFloat(strings.Fields(lines[i])[getColumnIndex(*key)], 64)
			b, _ := strconv.ParseFloat(strings.Fields(lines[j])[getColumnIndex(*key)], 64)
			if *numericWithSuffix {
				a, _ = parseNumericWithSuffix(strings.Fields(lines[i])[getColumnIndex(*key)])
				b, _ = parseNumericWithSuffix(strings.Fields(lines[j])[getColumnIndex(*key)])
			}
			if *reverse {
				return a > b
			}
			return a < b
		}
		if *sortByMonth {
			a, _ := time.Parse("Jan", strings.Fields(lines[i])[getColumnIndex(*key)])
			b, _ := time.Parse("Jan", strings.Fields(lines[j])[getColumnIndex(*key)])
			if *reverse {
				return a.After(b)
			}
			return a.Before(b)
		}
		if *ignoreTrailingSpace {
			return strings.TrimSpace(lines[i]) < strings.TrimSpace(lines[j])
		}
		if *reverse {
			return lines[i] > lines[j]
		}
		return lines[i] < lines[j]
	}

	// Сортировка строк
	sort.SliceStable(lines, compare)

	// Уникальность строк
	if *unique {
		lines = uniqueLines(lines)
	}

	// Проверка отсортированности строк
	if *checkSorted {
		if !isSorted(lines, compare) {
			fmt.Println("Данные не отсортированы.")
			os.Exit(1)
		}
		fmt.Println("Данные отсортированы.")
		return
	}

	// Открытие файла
	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка создания файла для записи:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// Вывод отсортированных строк
	for _, line := range lines {
		fmt.Fprintln(outputFile, line)
	}
	fmt.Println("Сортировка успешно завершена.")
}

// getColumnIndex возвращает индекс колонки для сортировки
func getColumnIndex(key string) int {
	if key == "" {
		return 0
	}
	columns := strings.Split(key, ",")
	index, _ := strconv.Atoi(columns[0])
	return index - 1
}

// parseNumericWithSuffix парсит числовое значение с учетом суффиксов (например, "1K" = 1024)
func parseNumericWithSuffix(value string) (float64, error) {
	suffixes := map[string]float64{
		"K":  1024,
		"M":  1024 * 1024,
		"G":  1024 * 1024 * 1024,
		"T":  1024 * 1024 * 1024 * 1024,
		"P":  1024 * 1024 * 1024 * 1024 * 1024,
		"E":  1024 * 1024 * 1024 * 1024 * 1024 * 1024,
		"Ki": 1024,
		"Mi": 1024 * 1024,
		"Gi": 1024 * 1024 * 1024,
		"Ti": 1024 * 1024 * 1024 * 1024,
		"Pi": 1024 * 1024 * 1024 * 1024 * 1024,
		"Ei": 1024 * 1024 * 1024 * 1024 * 1024 * 1024,
	}
	for suffix, factor := range suffixes {
		if strings.HasSuffix(value, suffix) {
			numPart := strings.TrimSuffix(value, suffix)
			num, err := strconv.ParseFloat(numPart, 64)
			if err != nil {
				return 0, err
			}
			return num * factor, nil
		}
	}
	return strconv.ParseFloat(value, 64)
}

// uniqueLines возвращает уникальные строки
func uniqueLines(lines []string) []string {
	uniqueMap := make(map[string]struct{})
	for _, line := range lines {
		uniqueMap[line] = struct{}{}
	}
	uniqueLines := make([]string, 0, len(uniqueMap))
	for line := range uniqueMap {
		uniqueLines = append(uniqueLines, line)
	}
	return uniqueLines
}

// isSorted проверяет, отсортированы ли строки
func isSorted(lines []string, compare func(i, j int) bool) bool {
	for i := 1; i < len(lines); i++ {
		if compare(i, i-1) {
			return false
		}
	}
	return true
}
