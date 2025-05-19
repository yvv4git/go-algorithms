//go:build ignore

package main

func circularArrayLoop(nums []int) bool {
	/*
		APPROACH: Depth-First Search (DFS)
		1. Используем DFS для обхода массива.
		2. На каждом шаге проверяем, не изменилось ли направление движения или не попали в петлю из одного элемента.
		3. Если находим цикл, возвращаем true.
		4. Если находимся в том же направлени
		3. Если находимся в разных направлениях, обновляем направление движения.
		5. Повторяем шаги 2-5, пока не достигнем конца массива.
		6. Если не найдем цикл, возвращаем false.

		TIME COMPLEXITY: O(n)
		- В худшем случае мы проходим по всем элементам массива.

		SPACE COMPLEXITY: O(n)
		- Используем стeк для рекурсии, где глубина стека может достигать n.
	*/
	n := len(nums)
	if n == 0 {
		return false
	}

	visited := make([]bool, n) // Массив для отметки посещенных индексов

	for i := 0; i < n; i++ {
		if visited[i] {
			continue // Пропускаем уже посещенные индексы
		}

		current := i
		direction := nums[current] > 0 // true = вперед, false = назад
		cycle := make(map[int]bool)    // Для отслеживания элементов в текущем цикле

		for {
			if visited[current] {
				break // Если уже посещали этот индекс, выходим
			}

			visited[current] = true
			cycle[current] = true

			next := (current + nums[current]) % n
			if next < 0 {
				next += n // Корректируем отрицательный индекс
			}

			// Проверяем условия:
			// 1. Смена направления (нельзя смешивать + и -)
			// 2. Петля из одного элемента (next == current)
			if (nums[next] > 0) != direction || next == current {
				break
			}

			// Если следующий индекс уже в текущем цикле -> найдена петля
			if cycle[next] {
				return true
			}

			current = next
		}
	}

	return false
}
