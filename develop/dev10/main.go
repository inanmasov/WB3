package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// Объявление и парсинг флагов командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "timeout for connection")
	flag.Parse()

	// Получение адреса хоста и порта из аргументов командной строки
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet [options] host port")
		os.Exit(1)
	}
	host := args[0]
	port := args[1]

	// Установка таймаута для подключения к серверу
	dialer := &net.Dialer{
		Timeout: *timeout,
	}

	// Подключение к серверу
	conn, err := dialer.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to", conn.RemoteAddr())

	// Запуск горутины для чтения данных из сокета и вывода их в STDOUT
	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("Connection closed by remote server")
		os.Exit(0)
	}()

	// Копирование данных из STDIN в сокет
	io.Copy(conn, os.Stdin)
	fmt.Println("Connection closed by client")
}
