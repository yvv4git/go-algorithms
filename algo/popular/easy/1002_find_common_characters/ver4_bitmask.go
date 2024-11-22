//go:build ignore

package main

import (
	"fmt"
)

func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	// Инициализация битовой маски для первой строки
	commonMask := 0
	for _, char := range words[0] {
		commonMask |= 1 << (char - 'a')
	}

	// Обновление битовой маски для каждой следующей строки
	for _, word := range words[1:] {
		currentMask := 0
		for _, char := range word {
			currentMask |= 1 << (char - 'a')
		}
		commonMask &= currentMask
	}

	// Собираем результат
	var result []string
	for i := 0; i < 26; i++ {
		if commonMask&(1<<i) != 0 {
			result = append(result, string('a'+i))
		}
	}

	return result
}

func main() {
	words := []string{"bella", "label", "roller"}
	fmt.Println(commonChars(words)) // Вывод: ["e", "l", "l"]

	words = []string{"cool", "lock", "cook"}
	fmt.Println(commonChars(words)) // Вывод: ["c", "o"]
}
