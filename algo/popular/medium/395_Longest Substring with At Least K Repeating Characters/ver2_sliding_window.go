//go:build ignore

package main

func longestSubstring(s string, k int) int {
	/*
		APPROACH: Sliding Window with Unique Characters Count
		Итерируем по всем возможным количествам уникальных символов (1..26).
		Для каждого варианта используем sliding window, поддерживая счетчики:
		- Общее количество уникальных символов в окне
		- Количество символов, встречающихся ≥k раз
		Максимальная длина окна, удовлетворяющего условиям - наш ответ.

		TIME COMPLEXITY: O(N)
		- Внешний цикл выполняется 26 раз (константа)
		- Внутренний цикл обрабатывает каждый символ ровно один раз
		- Общая сложность O(26*N) = O(N)

		SPACE COMPLEXITY: O(1)
		- Используем фиксированный массив на 26 элементов (частоты символов)
		- Не зависит от размера входных данных
	*/
	maxLen := 0
	for numUniqueTarget := 1; numUniqueTarget <= 26; numUniqueTarget++ {
		freq := [26]int{}
		left := 0
		numUnique := 0
		numAtLeastK := 0

		for right := 0; right < len(s); right++ {
			char := s[right] - 'a'
			if freq[char] == 0 {
				numUnique++
			}
			freq[char]++
			if freq[char] == k {
				numAtLeastK++
			}

			for numUnique > numUniqueTarget {
				leftChar := s[left] - 'a'
				if freq[leftChar] == k {
					numAtLeastK--
				}
				freq[leftChar]--
				if freq[leftChar] == 0 {
					numUnique--
				}
				left++
			}

			if numUnique == numUniqueTarget && numAtLeastK == numUnique {
				if windowLen := right - left + 1; windowLen > maxLen {
					maxLen = windowLen
				}
			}
		}
	}
	return maxLen
}

func main() {
	// Тестовые случаи для проверки
	testCases := []struct {
		s      string
		k      int
		expect int
	}{
		{"aaabb", 3, 3},
		{"ababbc", 2, 5},
		{"abcde", 2, 0},
		{"a", 1, 1},
		{"", 1, 0},
	}

	for _, tc := range testCases {
		result := longestSubstring(tc.s, tc.k)
		fmt.Printf("s: %q, k: %d → %d (expected %d)\n", tc.s, tc.k, result, tc.expect)
	}
}
