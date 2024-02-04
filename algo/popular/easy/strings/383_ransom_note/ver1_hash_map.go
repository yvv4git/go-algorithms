package _83_ransom_note

func canConstructV1(ransomNote string, magazine string) bool {
	/*
		METHOD: Hash map, Dictionary
		Time complexity: O(n)
		Space complexity: O(1)

		Time complexity
		Временная сложность этого алгоритма - O(n), где n - длина самого длинного из строк (ransomNote или magazine).
		Это связано с тем, что мы проходимся по каждой строке ровно один раз.

		Space complexity
		Пространственная сложность этого алгоритма - O(1), так как мы используем фиксированное количество пространства для хранения счетчиков букв,
		независимо от размера входных строк. В худшем случае это будет 26 для английского алфавита, что является константой.
	*/

	// Создаем мапу для подсчета количества букв в magazine
	letterCounts := make(map[rune]int)

	// Проходимся по magazine и увеличиваем счетчик для каждой буквы
	for _, letter := range magazine {
		letterCounts[letter]++
	}

	// Проходимся по ransomNote
	for _, letter := range ransomNote {
		// Если буквы нет в мапе или их количество меньше 1, то мы не можем составить ransomNote
		if letterCounts[letter] < 1 {
			return false
		}
		// Уменьшаем счетчик буквы
		letterCounts[letter]--
	}

	// Если мы успешно прошли все буквы ransomNote, то он может быть составлен
	return true
}
