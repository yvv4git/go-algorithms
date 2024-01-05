package ver1_dp

// Функция, которая находит самое короткое палиндромное расширение
func shortestPalindrome(s string) string {
	/*
		Method: Dynamic programming
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность этого алгоритма - O(n), где n - длина строки.
		Это связано с тем, что мы проходим по строке дважды: один раз для создания обратной строки,
		а второй раз для вычисления наибольшего префикса, который также является суффиксом.

		Space complexity
		Пространственная сложность - O(n), так как мы храним две новые строки размера n.
		Одна строка - это исходная строка, а вторая - это обратная строка.
	*/
	// Создаем новую строку, которая является обратной исходной строке
	revS := reverse(s)
	// Создаем новую строку, которая является исходной строкой и обратной строкой, разделенными символом '$'
	newS := s + "$" + revS
	// Находим наибольший префикс, который также является суффиксом для новой строки
	lps := computeLPSArray(newS)
	// Находим индекс, на котором начинается самый длинный палиндром
	index := lps[len(lps)-1]
	// Возвращаем самый короткий палиндром, который можно создать, добавив некоторые символы в начало строки
	return revS[:len(s)-index] + s
}

// Функция, которая находит наибольший префикс, который также является суффиксом
func computeLPSArray(s string) []int {
	lps := make([]int, len(s))
	length := 0
	i := 1
	for i < len(s) {
		if s[i] == s[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

// Функция, которая переворачивает строку
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
