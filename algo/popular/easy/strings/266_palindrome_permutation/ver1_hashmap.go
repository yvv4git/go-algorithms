package _66_palindrome_permutation

func canPermutePalindromeV1(s string) bool {
	/*
		METHOD: HashMap, Dict
		Time complexity: O(n), где n - длина строки, поскольку мы проходим по всем символам строки.
		Space complexity: O(n), поскольку в худшем случае мы можем хранить все символы строки в словаре
	*/
	// Создаем словарь для подсчета количества вхождений каждого символа
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}

	// Инициализируем счетчик нечетных символов
	oddCount := 0
	for _, count := range charCount {
		// Если количество вхождений нечетно, увеличиваем счетчик
		if count%2 != 0 {
			oddCount++
		}
	}

	// Если количество нечетных символов больше 1, то строка не может быть перестановкой палиндрома
	return oddCount <= 1
}
