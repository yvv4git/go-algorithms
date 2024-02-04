package _21_maximum_xor_of_two_numbers_in_an_array

type TrieNode struct {
	Children [2]*TrieNode
}

// Функция для поиска максимального XOR двух чисел в массиве
func findMaximumXORV1(nums []int) int {
	/*
		METHOD: Префиксное дерево
		Time complexity: O(n)
		Space complexity: O(m), где m - количество уникальных чисел в массиве

		Метод Trie (префиксное дерево) используется для эффективного хранения и поиска строк.
		В данном случае, мы используем Trie для хранения двоичных представлений чисел в массиве.
		Каждый узел Trie представляет бит числа. У каждого узла может быть до 2 дочерних узла, по одному для каждого возможного бита (0 и 1).
		Для каждого числа в массиве, мы добавляем его в Trie. Для этого мы проходим по каждому биту числа,
		и для каждого бита мы переходим к соответствующему дочернему узлу. Если дочерний узел не существует, мы создаем его.

		После того, как мы добавили все числа в Trie, мы ищем максимальный XOR двух чисел.
		Для этого мы проходим по каждому биту числа, и для каждого бита мы пытаемся перейти к дочернему узлу, противоположного биту.
		Если такой дочерний узел существует, мы переходим к нему, иначе мы остаемся в текущем узле.

		Максимальный XOR двух чисел будет равен сумме всех бит, которые мы перешли к дочерним узлам противоположного бита.

		Таким образом, мы можем найти максимальный XOR двух чисел в массиве за время O(n), где n - количество чисел в массиве.
	*/
	// Создаем корень Trie
	root := &TrieNode{}
	// Инициализируем максимальный XOR
	maxXor := 0

	// Добавляем числа в Trie
	for _, num := range nums {
		node := root
		xorNode := root
		currXor := 0
		// Проходим по каждому биту числа
		for i := 31; i >= 0; i-- {
			// Получаем i-ый бит числа
			bit := (num >> i) & 1
			// Если у дочернего узла нет бита, то создаем его
			if node.Children[bit] == nil {
				node.Children[bit] = &TrieNode{}
			}
			// Переходим к дочернему узлу
			node = node.Children[bit]

			// Пытаемся перейти в противоположный бит, чтобы получить максимальный XOR
			if xorNode.Children[1-bit] != nil {
				currXor += (1 << i)
				xorNode = xorNode.Children[1-bit]
			} else {
				xorNode = xorNode.Children[bit]
			}
		}
		// Обновляем максимальный XOR
		maxXor = max(maxXor, currXor)
	}

	return maxXor
}

// Функция для нахождения максимума из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
