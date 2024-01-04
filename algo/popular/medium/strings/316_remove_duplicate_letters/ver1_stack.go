package _16_remove_duplicate_letters

// Функция для удаления дубликатов букв
func removeDuplicateLettersV1(s string) string {
	/*
		Method: Stack
		Time complexity: O(n)
		Space complexity: O(1)

		Time complexity
		Временная сложность этого алгоритма - O(n), где n - длина строки, поскольку мы проходим по строке дважды:
		один раз для подсчета количества вхождений букв, а второй раз для удаления дубликатов.

		Space complexity
		Пространственная сложность - O(1), поскольку не используется дополнительная память, размер которой не зависит от размера входной строки.
	*/
	// Создаем массив для подсчета количества вхождений каждой буквы
	count := make([]int, 26)
	// Создаем массив для отслеживания букв, которые уже есть в стеке
	visited := make([]bool, 26)
	// Создаем стек для хранения уникальных букв
	stack := make([]rune, 0)

	// Подсчитываем количество вхождений каждой буквы
	for _, c := range s {
		count[c-'a']++
	}

	// Проходим по строке
	for _, c := range s {
		// Уменьшаем количество вхождений текущей буквы
		count[c-'a']--

		// Если буквы уже есть в стеке, то пропускаем ее
		if visited[c-'a'] {
			continue
		}

		// Пока стек не пуст и последняя буква в стеке больше текущей буквы,
		// и последняя буква в стеке еще встречается позже,
		// то удаляем последнюю букву из стека
		for len(stack) > 0 && c < stack[len(stack)-1] && count[stack[len(stack)-1]-'a'] > 0 {
			visited[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		// Добавляем текущую букву в стек
		stack = append(stack, c)
		visited[c-'a'] = true
	}

	// Преобразуем стек в строку и возвращаем ее
	return string(stack)
}
