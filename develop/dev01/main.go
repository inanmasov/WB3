package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Получаем точное время с помощью библиотеки NTP
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при получении точного времени:", err)
		os.Exit(1)
	}

	// Выводим точное время
	fmt.Println("Точное время:", ntpTime.Format(time.RFC3339))
}
