package main

import (
	"fmt"
)

// Генерирует все возможные подпоследовательности строки
func generateSubsequences(s string) []string {
	var result []string
	var backtrack func(string, int, string)

	backtrack = func(s string, index int, current string) {
		if index == len(s) {
			if current != "" {
				result = append(result, current)
			}
			return
		}
		// Включаем текущий символ
		backtrack(s, index+1, current+string(s[index]))
		// Не включаем текущий символ
		backtrack(s, index+1, current)
	}

	backtrack(s, 0, "")
	return result
}

// Проверяет, является ли подпоследовательностью
func isSubsequence(sub, s string) bool {
	i := 0
	for j := 0; i < len(sub) && j < len(s); j++ {
		if sub[i] == s[j] {
			i++
		}
	}
	return i == len(sub)
}

// Находит самую длинную необычную подпоследовательность
// Approach: Brute-force
// 1. Генерируем все возможные подпоследовательности для каждой строки
// 2. Для каждой подпоследовательности проверяем, является ли она необычной
//    (встречается только в одной строке)
// 3. Отслеживаем максимальную длину необычных подпоследовательностей
// Time Complexity: O(n * 2^m * n * m), где:
//   - n - количество строк
//   - m - максимальная длина строки
//   - 2^m - количество подпоследовательностей для строки длины m
//   - n*m - проверка подпоследовательности для всех строк
// Space Complexity: O(2^m) - хранение всех подпоследовательностей для одной строки
func findLUSlength(strs []string) int {
	maxLength := -1
	
	for i, s := range strs {
		subsequences := generateSubsequences(s)
		
		for _, sub := range subsequences {
			isUncommon := true
			
			for j, other := range strs {
				if i == j {
					continue
				}
				if isSubsequence(sub, other) {
					isUncommon = false
					break
				}
			}
			
			if isUncommon && len(sub) > maxLength {
				maxLength = len(sub)
			}
		}
	}
	
	return maxLength
}

func main() {
	// Примеры из README
	fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"})) // Ожидается 3
	fmt.Println(findLUSlength([]string{"aaa", "aaa", "aa"}))  // Ожидается -1
}
