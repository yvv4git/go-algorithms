package _400_construct_k_palindrome_strings

// Функция для подсчета количества вхождений символов в строку
func countChars(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	return counts
}

// Функция для проверки возможности составить k палиндромов из строки s
// Временная сложность: O(n), где n - длина строки s
// Пространственная сложность: O(1), так как в худшем случае мы можем хранить 26 символов (букв английского алфавита)
func canConstructV2(s string, k int) bool {
	// Подсчитываем количество вхождений символов в строке
	counts := countChars(s)

	// Подсчитываем количество символов, которые встречаются нечетное количество раз
	oddCount := 0
	for _, count := range counts {
		if count%2 == 1 {
			oddCount++
		}
	}

	// Если количество нечетных символов меньше или равно k, то можно составить ровно k палиндромов
	return oddCount <= k
}
