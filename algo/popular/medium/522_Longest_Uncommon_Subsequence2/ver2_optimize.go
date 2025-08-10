//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// Проверяет, является ли строка подпоследовательностью другой строки
func isSubsequence(sub, s string) bool {
	i := 0
	for j := 0; i < len(sub) && j < len(s); j++ {
		if sub[i] == s[j] {
			i++
		}
	}
	return i == len(sub)
}

// Проверяет, является ли строка необычной (не подпоследовательностью других строк)
func isUncommon(strs []string, idx int) bool {
	for i, s := range strs {
		if i == idx {
			continue
		}
		if isSubsequence(strs[idx], s) {
			return false
		}
	}
	return true
}

// Находит самую длинную необычную подпоследовательность
// Approach: Оптимизированный подход - проверка самых длинных строк
// 1. Сортируем строки по убыванию длины
// 2. Проверяем каждую строку на уникальность (не является подпоследовательностью других)
// 3. Первая найденная необычная строка будет самой длинной
// Time Complexity: O(n^2 * m), где:
//   - n - количество строк
//   - m - максимальная длина строки
//
// Space Complexity: O(1) - не используем дополнительной памяти
func findLUSlength(strs []string) int {
	// Сортируем строки по убыванию длины
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})

	// Проверяем каждую строку на уникальность
	for i := range strs {
		if isUncommon(strs, i) {
			// Проверяем, что строка не дублируется
			isDuplicate := false
			for j := 0; j < i; j++ {
				if strs[j] == strs[i] {
					isDuplicate = true
					break
				}
			}
			if !isDuplicate {
				return len(strs[i])
			}
		}
	}

	return -1
}

func main() {
	// Примеры из README
	fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"})) // Ожидается 3
	fmt.Println(findLUSlength([]string{"aaa", "aaa", "aa"}))  // Ожидается -1
}
