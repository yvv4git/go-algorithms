package main

import (
	"fmt"
	"strings"
)

func longestSubstring(s string, k int) int {
	/*
		APPROACH: Divide and Conquer
		Рекурсивно находим и исключаем символы, встречающиеся менее k раз,
		используя их как точки разделения строки. Для каждой полученной подстроки
		повторяем процесс, пока не найдем подстроку, где все символы встречаются ≥k раз.
		Максимальная длина такой подстроки является ответом.

		TIME COMPLEXITY: O(N^2)
		- В худшем случае (например, "ababab...ab" при k=2) выполняем O(N^2) операций
		- В среднем случае работает быстрее благодаря разделению строки

		SPACE COMPLEXITY: O(N)
		- Хранение частот символов требует O(N) памяти
		- Глубина рекурсии в худшем случае O(N)
	*/
	if len(s) < k {
		return 0
	}

	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	for char, count := range freq {
		if count < k {
			maxLen := 0
			for _, sub := range strings.Split(s, string(char)) {
				if current := longestSubstring(sub, k); current > maxLen {
					maxLen = current
				}
			}
			return maxLen
		}
	}

	return len(s)
}

func main() {
	// Тестовые случаи
	testCases := []struct {
		s      string
		k      int
		expect int
	}{
		{"aaabb", 3, 3},
		{"ababbc", 2, 5},
		{"abcde", 2, 0},
		{"", 1, 0},
		{"a", 1, 1},
	}

	for _, tc := range testCases {
		result := longestSubstring(tc.s, tc.k)
		fmt.Printf("Input: %q, k=%d | Output: %d | Expected: %d | Match: %t\n",
			tc.s, tc.k, result, tc.expect, result == tc.expect)
	}
}
