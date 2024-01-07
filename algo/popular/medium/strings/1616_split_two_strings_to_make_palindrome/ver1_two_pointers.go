package main

// Функция для проверки, можно ли сделать из двух строк палиндром
func checkPalindromeFormationV1(a string, b string) bool {
	/*
		Method: Two pointers
		Time complexity: O(n)
		Space complexity: O(1)

		Time complexity
		Временная сложность этого алгоритма - O(n), где n - длина строк. Это связано с тем, что мы проходим по строкам только один раз.

		Space complexity
		Пространственная сложность - O(1), так как мы используем небольшое количество дополнительной памяти,
		не зависящей от размера входных данных.
	*/

	// 1. Проверяем, можно ли из строк a+b сделать палиндром.
	// 2. Проверяем, можно ли из строк b+a сделать палиндром
	return check(a, b) || check(b, a)
}

// Функция для проверки, является ли подстрока палиндромом
func check(a string, b string) bool {
	n := len(a)
	start := 0
	end := n - 1

	// Проходим по строкам, начиная с концов, и проверяем, являются ли символы одинаковыми
	for start < end && a[start] == b[end] {
		start++
		end--
	}

	// Проверяем, является ли полученная подстрока палиндромом
	return isPalindrome(a[start:end+1]) || isPalindrome(b[start:end+1])
}

// Функция для проверки, является ли строка палиндромом
func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}
