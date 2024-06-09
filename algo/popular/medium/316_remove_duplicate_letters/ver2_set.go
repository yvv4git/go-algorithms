package _16_remove_duplicate_letters

import "strings"

func removeDuplicateLettersV2(s string) string {
	/*
		METHOD: Set
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)

		Time complexity
		Временная сложность этого алгоритма - O(n), где n - длина строки, поскольку мы проходим по строке дважды:
		один раз для подсчета количества вхождений букв, а второй раз для удаления дубликатов.

		Space complexity
		Пространственная сложность - O(1), поскольку не используется дополнительная память, размер которой не зависит от размера входной строки.
	*/
	// Создаем множество для отслеживания уникальных символов
	set := make(map[rune]bool)

	// Создаем строку для результата
	var result strings.Builder

	// Проходим по строке
	for _, char := range s {
		// Если символ еще не встречался, добавляем его в множество и в результат
		if _, ok := set[char]; !ok {
			set[char] = true
			result.WriteRune(char)
		}
	}

	// Возвращаем результат
	return result.String()
}
