package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ") // Выводим приглашение для ввода команды
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}

		// Удаляем символ новой строки из введенной строки
		input = strings.TrimSpace(input)

		// Разбиваем введенную строку на команду и аргументы
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		// Обработка команды
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Необходимо указать путь для cd")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Ошибка смены директории:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Ошибка получения текущей директории:", err)
				continue
			}
			fmt.Println(dir)
		case "echo":
			if len(args) < 2 {
				fmt.Println("Нет аргументов для вывода")
				continue
			}
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Необходимо указать PID для команды kill.")
				continue
			}
			pid, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Ошибка при поиске процесса:", err)
				continue
			}
			proc, err := os.FindProcess(pid)
			if err != nil {
				fmt.Println("Ошибка при поиске процесса:", err)
				continue
			}
			err = proc.Signal(os.Kill)
			if err != nil {
				fmt.Println("Ошибка при завершении процесса:", err)
			}
		case "ps":
			cmd := exec.Command("tasklist")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при выполнении команды tasklist:", err)
			}
		case "exit":
			fmt.Println("Выход из программы")
			return
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка выполнения команды:", err)
			}
		}
	}
}
