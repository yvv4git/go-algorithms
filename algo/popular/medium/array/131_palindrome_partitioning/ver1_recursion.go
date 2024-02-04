package _31_palindrome_partitioning

// Функция partitionV1 разбивает строку на палиндромы
func partitionV1(s string) [][]string {
	/*
		METHOD: Recursion
		TIME COMPLEXITY: O(n * 2^n)
		Space complexity: O(n * 2^n)

		Time complexity
		Временная сложность этого алгоритма составляет O(n * 2^n), где n - длина строки.
		Это связано с тем, что для каждого символа мы можем делать выбор - включать его в текущую часть или начать новую часть.
		В худшем случае мы можем сделать 2^n выборов.

		Space complexity
		Пространственная сложность также составляет O(n * 2^n), поскольку в худшем случае мы можем хранить 2^n разбиений строки.
	*/

	// Базовый случай: пустая строка
	if len(s) == 0 {
		return [][]string{{}}
	}

	var result [][]string
	for i := 1; i <= len(s); i++ {
		// Если префикс является палиндромом
		if isPalindrome(s[:i]) {
			// Рекурсивно разбиваем остаток строки
			for _, partitions := range partitionV1(s[i:]) {
				// Добавляем префикс к каждому разбиению
				result = append(result, append([]string{s[:i]}, partitions...))
			}
		}
	}
	return result
}

// Функция partitionV1 проверяет, является ли строка палиндромом
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
