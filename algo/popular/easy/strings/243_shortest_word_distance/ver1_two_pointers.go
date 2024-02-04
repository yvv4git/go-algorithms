package main

func shortestDistance(words []string, word1 string, word2 string) int {
	/*
		METHOD: Two pointers
		TIME COMPLEXITY: O(n), где n - количество слов в массиве, потому что мы проходим по массиву всего один раз.
		Space complexity: O(1), так как мы используем не более константного количества дополнительной памяти.
	*/

	// Инициализируем два указателя idx1 и idx2 в -1, что означает, что они еще не были обновлены.
	idx1, idx2 := -1, -1
	// Инициализируем переменную minDistance длиной массива words, чтобы она всегда была больше, чем возможное расстояние между словами.
	minDistance := len(words)

	// Проходим по всем словам в массиве words.
	for i, word := range words {
		// Если текущее слово равно word1, обновляем idx1.
		if word == word1 {
			idx1 = i
			// Если текущее слово равно word2, обновляем idx2.
		} else if word == word2 {
			idx2 = i
		}

		// Если оба указателя обновлены, обновляем минимальное расстояние между ними.
		if idx1 != -1 && idx2 != -1 {
			minDistance = min(minDistance, abs(idx1-idx2))
		}
	}

	// Возвращаем найденное минимальное расстояние.
	return minDistance
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
