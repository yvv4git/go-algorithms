package main

// Функция canMeasureWaterV3 решает задачу "Water and Jug Problem"
// с использованием Breadth-First Search (BFS) для поиска всех возможных состояний
// двух ведер и проверки, можно ли достичь целевого объема воды.
func canMeasureWaterV3(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	/*
		METHOD: BFS
		- Используется Breadth-First Search (BFS) для поиска всех возможных состояний двух ведер.
		- Начинаем с начального состояния, когда оба ведра пусты, и добавляем все возможные переходы в очередь.
		- Проверяем каждое состояние, чтобы увидеть, достигает ли оно целевого объема воды.

		TIME COMPLEXITY:
		- Временная сложность зависит от количества уникальных состояний, которые могут быть посещены.
		- В худшем случае, каждое состояние может быть посещено один раз, что приводит к экспоненциальной временной сложности O(n^2), где n - количество операций.

		SPACE COMPLEXITY:
		- Пространственная сложность определяется размером карты visited, которая хранит посещенные состояния.
		- В худшем случае, карта может содержать все возможные состояния, что приводит к пространственной сложности O(n^2), где n - количество уникальных состояний.
	*/
	// Создаем карту для отслеживания посещенных состояний
	visited := make(map[[2]int]bool)
	// Состояние перехода выполняется с использованием пары <количество в ведре1, количество в ведре2>
	// Мы начинаем с обоих ведер пустыми
	var queue [][2]int
	// Инициализируем очередь начальным состоянием, когда оба ведра пусты
	initial := [2]int{0, 0}
	queue = append(queue, initial)
	visited[initial] = true

	// Пока очередь не пуста, продолжаем поиск
	for len(queue) != 0 {
		// Извлекаем текущее состояние из очереди
		curr := queue[0]
		queue = queue[1:]

		// Если сумма количества воды в обоих ведрах равна целевому объему, возвращаем true
		if curr[0]+curr[1] == targetCapacity {
			return true
		}

		// Возможные переходы состояний
		var possibleStates [][2]int
		if curr[0] == 0 {
			// Заполняем ведро1 водой
			possibleStates = append(possibleStates, [2]int{jug1Capacity, curr[1]})
		} else {
			// Опустошаем ведро1
			possibleStates = append(possibleStates, [2]int{0, curr[1]})
			// Переливаем воду из ведра1 в ведро2
			// Доступное пространство в ведре2 - (jug2Capacity - curr[1])
			possibleStates = append(possibleStates, [2]int{
				max(0, curr[0]-(jug2Capacity-curr[1])), // ведро1
				min(jug2Capacity, curr[1]+curr[0]),     // ведро2
			})
		}
		if curr[1] == 0 {
			// Заполняем ведро2 водой
			possibleStates = append(possibleStates, [2]int{curr[0], jug2Capacity})
		} else {
			// Опустошаем ведро2
			possibleStates = append(possibleStates, [2]int{curr[0], 0})
			// Переливаем воду из ведра2 в ведро1
			// Логика такая же, как и выше, но в обратном порядке
			possibleStates = append(possibleStates, [2]int{
				min(jug1Capacity, curr[0]+curr[1]),
				max(0, curr[1]-(jug1Capacity-curr[0])),
			})
		}

		// Добавляем все возможные состояния в очередь, если они еще не были посещены
		for _, next := range possibleStates {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}

	// Если мы не нашли способ достичь целевого объема, возвращаем false
	return false
}

// Вспомогательная функция для нахождения максимума двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Вспомогательная функция для нахождения минимума двух чисел
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
