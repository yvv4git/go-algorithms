package __longest_substring_without_repeating_characters

// Функция для поиска самой длинной подстроки без повторяющихся символов
func lengthOfLongestSubstringV1(s string) int {
	/*
		METHOD: Bruteforce
		TIME COMPLEXITY: O(n^3)
		SPACE COMPLEXITY: O(min(n,m)), где n - длина строки, а m - размер алфавита

		В этом коде мы используем метод брутфорса для решения задачи.
		Мы проходим по всем возможным подстрокам строки и проверяем, является ли каждая подстрока уникальной.
		Если подстрока уникальна и ее длина больше максимальной длины, то мы обновляем максимальную длину.
		В конце мы возвращаем максимальную длину.

		Time complexity
		Временная сложность для этого алгоритма составляет O(n^3), где n - длина строки.
		Это связано с двумя циклами, которые проходят по всем возможным подстрокам строки, и функцией isUnique,
		которая также проходит по всем символам подстроки.

		Space complexity
		Пространственная сложность составляет O(min(n, m)), где m - размер алфавита.
		Это связано с хранением символов в map для проверки уникальности.
		В худшем случае, когда все символы уникальны, размер map будет равен размеру алфавита.
	*/
	maxLength := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substring := s[i:j]
			if isUnique(substring) && len(substring) > maxLength {
				maxLength = len(substring)
			}
		}
	}

	return maxLength
}

// Функция для проверки уникальности символов в строке
func isUnique(s string) bool {
	m := make(map[rune]bool)
	for _, char := range s {
		if m[char] {
			return false
		}
		m[char] = true
	}
	return true
}
