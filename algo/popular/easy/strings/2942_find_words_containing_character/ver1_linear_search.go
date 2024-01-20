package main

import "strings"

func findWordsContaining(words []string, x byte) []int {
	/*
		Method: Linear search
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность функции findWordsContaining составляет O(n), где n - количество слов в массиве words.
		Это связано с тем, что функция проходит по каждому слову в массиве ровно один раз.

		Space complexity
		Пространственная сложность функции также составляет O(n),
		так как в худшем случае (когда все слова содержат искомый символ) мы можем заполнить срез indices всеми индексами из исходного массива words.
		Однако, в общем случае, количество элементов в indices будет меньше или равно количеству элементов в words.
	*/
	// Создаем пустой срез для хранения индексов слов, содержащих символ x
	var indices []int // индексы

	// Проходим по всем словам в массиве words
	for i, word := range words {
		// Если слово содержит символ x, добавляем его индекс в срез indices
		if strings.ContainsRune(word, rune(x)) {
			indices = append(indices, i)
		}
	}

	// Возвращаем срез индексов слов, содержащих символ x
	return indices
}
