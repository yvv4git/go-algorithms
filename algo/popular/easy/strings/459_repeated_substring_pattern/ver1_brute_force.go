package _59_repeated_substring_pattern

import "strings"

func repeatedSubstringPatternV1(s string) bool {
	/*
		METHOD: Brute force
		TIME COMPLEXITY: O(n^2)
		Space complexity: O(n)

		Time complexity
		Временная сложность этого алгоритма составляет O(n^2), где n - длина строки.
		Это связано с тем, что для каждой подстроки мы создаем новую строку, которая может быть достаточно большой.

		Space complexity
		Пространственная сложность этого алгоритма составляет O(n), где n - длина строки.
		Это связано с тем, что мы создаем новую строку для каждой подстроки.
	*/
	length := len(s)

	// Проходим по всем возможным длинам подстрок
	for subLen := 1; subLen <= length/2; subLen++ {
		// Если длина не делится нацело на длину подстроки, то пропускаем эту длину
		if length%subLen != 0 {
			continue
		}

		// Получаем подстроку длины subLen
		sub := s[:subLen]

		// Проверяем, можно ли составить исходную строку, добавив несколько копий этой подстроки
		if strings.Repeat(sub, length/subLen) == s {
			return true
		}
	}

	return false
}
