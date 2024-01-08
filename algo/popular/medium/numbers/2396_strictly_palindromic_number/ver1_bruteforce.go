package main

// Функция для перевода числа в произвольную систему счисления
func convertToBase(n int, base int) string {
	digits := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if n < base {
		return string(digits[n])
	}
	return convertToBase(n/base, base) + string(digits[n%base])
}

// Функция для проверки, является ли строка палиндромом
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// Функция для проверки, является ли число строго палиндромным
func isStrictlyPalindromicV1(n int) bool {
	/*
		Method: Bruteforce
		Time complexity: O(n^2), где n - входное число. Это связано с двумя циклами: внешним, который проходит по всем основаниям системы счисления, и внутренним, который проверяет, является ли полученная строка палиндромом.
		Space complexity: O(1), так как в худшем случае мы можем хранить в памяти строку размером n.
	*/

	// Проходим по всем основаниям системы счисления от 2 до n - 2
	for base := 2; base <= n-2; base++ {
		// Конвертируем число в систему счисления base
		converted := convertToBase(n, base)
		// Проверяем, является ли полученная строка палиндромом
		if !isPalindrome(converted) {
			return false
		}
	}

	return true
}
