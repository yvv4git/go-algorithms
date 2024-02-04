package main

// Функция для проверки, можно ли сделать из двух строк палиндром
func checkPalindromeFormationV2(a string, b string) bool {
	/*
		METHOD: Two pointers & Greedy
		Time complexity: O(n)
		Space complexity: O(n), helper будет создавать новые строки, которые могут быть достаточно большими.


		Алгоритм называется "жадным", если он всегда делает локально оптимальный выбор в каждый момент времени,
		основываясь только на информации, доступной в текущем момент.
		Это означает, что он не рассматривает последствия своих действий на будущее, а всегда выбирает действие, которое,
		по его мнению, даст наибольшее возможное количество вознаграждения в текущем момент времени.
	*/

	// Проверяем, является ли строка a или строка b палиндромом
	// Или, используя вспомогательную функцию helper, проверяем, можно ли сделать из строк a и b палиндром
	return isPalindrome(a) || isPalindrome(b) || helper(a, b) || helper(b, a)
}

// Вспомогательная функция для проверки, можно ли сделать из двух строк палиндром
func helper(a string, b string) bool {
	var i, j int
	// Проходим по строкам, начиная с начала и конца
	for i, j = 0, len(b)-1; i < len(a)/2; i, j = i+1, j-1 {
		// Если символы не равны, прерываем цикл
		if a[i] != b[j] {
			break
		}
	}

	// Если i > j, значит, строки a и b уже являются палиндромами
	// Или, проверяем, является ли среза строки b[i:j+1] или строки a[i:j+1] палиндромом
	return i > j || isPalindrome(b[i:j+1]) || isPalindrome(a[i:j+1])
}

// Функция для проверки, является ли строка палиндромом
//func isPalindrome(s string) bool {
//	for i := 0; i < len(s)/2; i++ {
//		if s[i] != s[len(s)-1-i] {
//			return false
//		}
//	}
//	return true
//}
