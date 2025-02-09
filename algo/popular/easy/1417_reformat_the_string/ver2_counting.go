//go:build ignore

package main

import (
	"fmt"
	"unicode"
)

func reformat(s string) string {
	/*
		METHOD: Counting and interleaving
		1. Считаем количество букв и цифр.
		2. Если разница в количестве больше 1, возвращаем пустую строку.
		3. Создаем результирующую строку, чередуя буквы и цифры.

		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n) (для результата)
	*/
	var letterCount, digitCount int
	for _, char := range s {
		if unicode.IsLetter(char) {
			letterCount++
		} else if unicode.IsDigit(char) {
			digitCount++
		}
	}

	// Проверяем, можно ли переформатировать строку
	if abs(letterCount-digitCount) > 1 {
		return ""
	}

	// Создаем результирующую строку
	result := make([]rune, len(s))
	letterIndex, digitIndex := 0, 1
	if letterCount > digitCount {
		letterIndex, digitIndex = 0, 1
	} else {
		letterIndex, digitIndex = 1, 0
	}

	for _, char := range s {
		if unicode.IsLetter(char) {
			if letterIndex >= len(s) {
				return ""
			}
			result[letterIndex] = char
			letterIndex += 2
		} else if unicode.IsDigit(char) {
			if digitIndex >= len(s) {
				return ""
			}
			result[digitIndex] = char
			digitIndex += 2
		}
	}

	return string(result)
}

// Функция для вычисления абсолютного значения разницы
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Примеры использования
	fmt.Println(reformat("a0b1c2"))     // "a0b1c2" (или другой допустимый вариант)
	fmt.Println(reformat("leetcode"))   // ""
	fmt.Println(reformat("1229857369")) // ""
	fmt.Println(reformat("covid2019"))  // "c2o0v1i9d" (или другой допустимый вариант)
}
