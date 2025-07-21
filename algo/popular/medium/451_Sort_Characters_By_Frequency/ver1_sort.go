//go:build ignore

package main

import (
	"fmt"
	"sort"
	"strings"
)

// Approach: Sorting
// frequencySort сортирует символы строки по убыванию частоты
// Временная сложность: O(n log n) из-за сортировки
// Пространственная сложность: O(n) для хранения частот и результата
func frequencySort(s string) string {
	// Подсчет частот символов
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	// Создание slice символов для сортировки
	chars := make([]rune, 0, len(freq))
	for char := range freq {
		chars = append(chars, char)
	}

	// Сортировка символов по убыванию частоты
	sort.Slice(chars, func(i, j int) bool {
		return freq[chars[i]] > freq[chars[j]]
	})

	// Построение результирующей строки
	var builder strings.Builder
	for _, char := range chars {
		builder.WriteString(strings.Repeat(string(char), freq[char]))
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
