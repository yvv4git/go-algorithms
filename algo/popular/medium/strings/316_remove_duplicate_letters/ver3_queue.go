package _16_remove_duplicate_letters

import "strings"

func removeDuplicateLettersV3(s string) string {
	/*
		METHOD: Queue
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность этого алгоритма составляет O(n), где n - количество символов в строке.
		Это связано с тем, что мы проходим по строке дважды: один раз для подсчета количества вхождений каждого символа,
		а второй раз для формирования результирующей строки.

		Space complexity
		Пространственная сложность также составляет O(n), так как в худшем случае мы можем хранить все символы в очереди.
	*/
	// Создаем очередь для отслеживания уникальных символов
	var queue []rune

	// Создаем множество для отслеживания символов, которые уже есть в очереди
	set := make(map[rune]bool)

	// Проходим по строке
	for _, char := range s {
		// Если символ еще не встречался, добавляем его в очередь и в множество
		if _, ok := set[char]; !ok {
			queue = append(queue, char)
			set[char] = true
		}
	}

	// Создаем строку для результата
	var result strings.Builder

	// Проходим по очереди
	for _, char := range queue {
		// Добавляем символ в результат
		result.WriteRune(char)
	}

	// Возвращаем результат
	return result.String()
}
