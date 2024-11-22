//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func commonCharsV3(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	// Преобразуем каждую строку в срез байтов и сортируем его
	for i := range words {
		byteSlice := []byte(words[i])
		sort.Slice(byteSlice, func(j, k int) bool {
			return byteSlice[j] < byteSlice[k]
		})
		words[i] = string(byteSlice)
	}

	// Инициализация результата
	var result []string

	// Используем двух указателей для поиска общих символов
	for i := 0; i < len(words[0]); i++ {
		char := words[0][i]
		common := true
		for j := 1; j < len(words); j++ {
			if !containsChar(words[j], char) {
				common = false
				break
			}
		}
		if common {
			result = append(result, string(char))
		}
	}

	return result
}

func containsChar(word string, char byte) bool {
	for i := 0; i < len(word); i++ {
		if word[i] == char {
			return true
		}
	}
	return false
}

func main() {
	words := []string{"bella", "label", "roller"}
	fmt.Println(commonCharsV3(words)) // Вывод: ["e", "l", "l"]

	words = []string{"cool", "lock", "cook"}
	fmt.Println(commonCharsV3(words)) // Вывод: ["c", "o"]
}
