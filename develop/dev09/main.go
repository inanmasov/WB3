package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Определение флагов командной строки
	urlFlag := flag.String("url", "", "URL сайта для загрузки")
	outputFlag := flag.String("output", "", "Название файла для сохранения")
	flag.Parse()

	// Проверка наличия URL
	if *urlFlag == "" {
		fmt.Println("Необходимо указать URL сайта для загрузки")
		return
	}

	// Выполнение HTTP-запроса для получения содержимого страницы
	response, err := http.Get(*urlFlag)
	if err != nil {
		fmt.Println("Ошибка при выполнении HTTP-запроса:", err)
		return
	}
	defer response.Body.Close()

	// Чтение содержимого страницы
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении содержимого страницы:", err)
		return
	}

	// Определение имени файла для сохранения
	filename := *outputFlag
	if filename == "" {
		// Использование части URL в качестве имени файла
		parts := strings.Split(*urlFlag, "/")
		filename = parts[len(parts)-1]
	}

	// Запись содержимого страницы в файл
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи содержимого в файл:", err)
		return
	}

	fmt.Println("Страница успешно загружена и сохранена в файл", filename)
}
