package main

import (
	"fmt"
	"strings"
	"unicode"
)

func UnpackingString(st string) string {
	str := []rune(st)

	if len(str) == 0 {
		return ""
	}

	if unicode.IsDigit(str[0]) {
		fmt.Println("некорректная строка")
		return ""
	}

	var strOut string
	// Проходим по вей строке
	for i := 0; i < len(str); {
		// Обрабатываем последний символ
		if i == len(str)-1 {
			strOut += string(str[i])
			break
		}
		// Проверяем на Escape-последовательность
		if CheckEscape(str[i]) {
			// Обрабатываем последнюю последовательность
			if i+1 == len(str)-1 {
				strOut += string(str[i+1])
				break
			} else if unicode.IsDigit(str[i+2]) { // Если Escape-последовательность надо вставить в нескольких экземплярах
				strOut += strings.Repeat(string(str[i+1]), int(str[i+2]-'0'))
				i += 3
				continue
			} else { // Если Escape-последовательность надо вставить в одной экземпляре
				strOut += string(str[i+1])
				i += 2
				continue
			}

		}
		// Случай когда после символа идет число
		if !unicode.IsDigit(str[i]) && unicode.IsDigit(str[i+1]) {
			strOut += strings.Repeat(string(str[i]), int(str[i+1]-'0'))
			i += 2
		} else if !unicode.IsDigit(str[i]) && !unicode.IsDigit(str[i+1]) { // После символа идет символ
			strOut += string(str[i])
			i++
		} else if unicode.IsDigit(str[i]) && unicode.IsDigit(str[i+1]) { // Обработка неккоректной строки
			fmt.Println("некорректная строка")
			return ""
		}
	}
	return strOut
}

// Проверка на Escape-последовательность
func CheckEscape(escape rune) bool {
	if string(escape) == "\\" {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(UnpackingString("a4bc2d5e"))
	fmt.Println(UnpackingString("abcd"))
	fmt.Println(UnpackingString("45"))
	fmt.Println(UnpackingString(""))
	fmt.Println(UnpackingString(`qwe\4\5`))
	fmt.Println(UnpackingString(`qwe\45`))
	fmt.Println(UnpackingString(`qwe\\5`))
}
