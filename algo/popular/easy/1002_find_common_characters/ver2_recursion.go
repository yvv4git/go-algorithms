//go:build ignore

package main

import "fmt"

func commonCharsV2(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	// Инициализация результата
	var result []string

	// Рекурсивная функция для поиска общих символов
	var findCommon func(index int, common map[rune]int)
	findCommon = func(index int, common map[rune]int) {
		if index == len(words) {
			for char, count := range common {
				for i := 0; i < count; i++ {
					result = append(result, string(char))
				}
			}
			return
		}

		currentFreq := make(map[rune]int)
		for _, char := range words[index] {
			currentFreq[char]++
		}

		for char := range common {
			if currentFreq[char] < common[char] {
				common[char] = currentFreq[char]
			}
		}

		findCommon(index+1, common)
	}

	// Инициализация мапы для первой строки
	common := make(map[rune]int)
	for _, char := range words[0] {
		common[char]++
	}

	// Начинаем рекурсию
	findCommon(1, common)

	return result
}

func main() {
	words := []string{"bella", "label", "roller"}
	fmt.Println(commonCharsV2(words)) // Вывод: ["e", "l", "l"]

	words = []string{"cool", "lock", "cook"}
	fmt.Println(commonCharsV2(words)) // Вывод: ["c", "o"]
}
