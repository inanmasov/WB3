package main

import (
	"fmt"
	"sort"
	"strings"
)

// FindAnagrams находит все множества анаграмм по словарю
func FindAnagrams(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

	// Создаем словарь, где ключ - отсортированное слово, значение - массив слов-анаграмм
	anagramMap := make(map[string][]string)
	for _, word := range words {
		sortedWord := sortString(strings.ToLower(word))
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	// Проходим по словарю и добавляем множества анаграмм в результирующую мапу
	for _, anagrams := range anagramMap {
		if len(anagrams) > 1 {
			firstWord := anagrams[0] // Берем первое слово из множества в качестве ключа
			sort.Strings(anagrams)   // Сортируем множество анаграмм
			anagramSets[firstWord] = anagrams
		}
	}

	return anagramSets
}

// sortString сортирует символы в строке
func sortString(s string) string {
	sChars := strings.Split(s, "")
	sort.Strings(sChars)
	return strings.Join(sChars, "")
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "ток", "кот", "ткань", "накат"}
	anagramSets := FindAnagrams(words)
	for key, value := range anagramSets {
		fmt.Println(key, ":", value)
	}
}
