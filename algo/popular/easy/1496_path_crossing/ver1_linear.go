package main

type Point struct {
	x, y int
}

func isPathCrossing(path string) bool {
	/*
		METHOD: Linear
		Суть решения в том, что мы храним все посещенные точки в хеш-карте и двигаемся по координатам согласно инструкциям (N,S,E,W),
		а как только попадаем в точку, которая уже есть в карте - значит путь пересекает сам себя.

		TIME COMPLEXITY: O(n), где n - длина строки `path`.
		- Мы проходим по всем символам строки один раз.
		- В худшем случае, когда все символы различны, мы посетим все точки на плоскости.
		- Временная сложность алгоритма линейна относительно длины строки `path`.

		SPACE COMPLEXITY: O(n), где n - длина строки `path`.
		- Мы используем карту `visited` для хранения посещенных точек,
		- которая может содержать до n элементов.
		- Пространственная сложность алгоритма линейна относительно длины строки `path`.
	*/

	// Создаем карту для хранения посещенных точек
	visited := make(map[Point]bool)

	// Начальная точка (0,0)
	current := Point{0, 0}
	visited[current] = true

	// Проходим по всем символам пути
	for _, direction := range path {
		// Обновляем координаты в зависимости от направления
		switch direction {
		case 'N':
			current.y++
		case 'S':
			current.y--
		case 'E':
			current.x++
		case 'W':
			current.x--
		}

		// Проверяем, были ли мы уже в этой точке
		if visited[current] {
			return true
		}

		// Добавляем новую точку в множество посещенных
		visited[current] = true
	}

	return false
}
