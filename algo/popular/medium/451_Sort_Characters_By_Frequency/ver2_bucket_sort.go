//go:build ignore

package main

import (
	"fmt"
	"strings"
)

// Approach: Bucket Sort
// frequencySort сортирует символы строки по убыванию частоты с использованием bucket sort
// Временная сложность: O(n) - линейная сложность
// Пространственная сложность: O(n) - для хранения ведер и результата
func frequencySort(s string) string {
	// Подсчет частот символов
	freq := make(map[rune]int)
	maxFreq := 0
	for _, char := range s {
		freq[char]++
		if freq[char] > maxFreq {
			maxFreq = freq[char]
		}
	}

	// Создание ведер (buckets)
	buckets := make([][]rune, maxFreq+1)
	for char, count := range freq {
		buckets[count] = append(buckets[count], char)
	}

	// Построение результата
	var builder strings.Builder
	for i := maxFreq; i > 0; i-- {
		for _, char := range buckets[i] {
			builder.WriteString(strings.Repeat(string(char), i))
		}
	}

	return builder.String()
}

func main() {
	// Пример использования
	examples := []string{
		"tree",
		"cccaaa",
		"Aabb",
	}

	for _, example := range examples {
		fmt.Printf("Input: %q\n", example)
		fmt.Printf("Output: %q\n\n", frequencySort(example))
	}
}
