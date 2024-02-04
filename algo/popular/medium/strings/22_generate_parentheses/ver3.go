package main

// Stackitem представляет собой элемент стека, который содержит количество открывающих и закрывающих скобок, а также текущую строку скобочных последовательностей.
type stackItem struct {
	open  int    // Количество открывающих скобок
	close int    // Количество закрывающих скобок
	str   string // Текущая строка скобочных последовательностей
}

// GenerateParenthesisV3 генерирует все правильные скобочные последовательности заданной длины.
func generateParenthesisV3(n int) []string {
	/*
		METHOD: Iterative
		TIME COMPLEXITY: O(4^n / sqrt(n))
		Space complexity: O(4^n / sqrt(n))
	*/
	// Результирующий срез для хранения всех правильных скобочных последовательностей.
	result := make([]string, 0)
	// Инициализация стека с единственным элементом, представляющим пустую строку скобок.
	stack := []stackItem{
		{0, 0, ""},
	}

	// Пока стек не пуст, извлекаем элемент и проверяем его на соответствие условиям.
	for len(stack) != 0 {
		curr := stack[len(stack)-1]  // Получаем последний элемент стека.
		stack = stack[:len(stack)-1] // Удаляем элемент из стека.

		// Если количество открывающих и закрывающих скобок равно 2*n, то добавляем текущую строку в результат.
		if curr.close+curr.open == 2*n {
			result = append(result, curr.str)
			continue
		}

		// Если количество открывающих скобок меньше n, добавляем открывающую скобку в стек.
		if curr.open < n {
			stack = append(stack, stackItem{curr.open + 1, curr.close, curr.str + "("})
		}

		// Если количество закрывающих скобок меньше количества открывающих, добавляем закрывающую скобку в стек.
		if curr.close < curr.open {
			stack = append(stack, stackItem{curr.open, curr.close + 1, curr.str + ")"})
		}
	}

	// Возвращаем результирующий срез всех правильных скобочных последовательностей.
	return result
}
