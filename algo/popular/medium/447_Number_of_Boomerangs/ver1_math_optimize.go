//go:build ignore

package main

import "fmt"

// numberOfBoomerangs решает задачу с использованием математической оптимизации
// Подход: используем квадрат расстояния вместо самого расстояния + хеш-таблица
// Сложность: O(n²) по времени, O(n) по памяти
func numberOfBoomerangs(points [][]int) int {
	/*
		APPROACH: Hash table.
		Для каждой точки i в списке точек:
			- Создаем хеш-таблицу для подсчета количества точек на каждом расстоянии от точки i.
			- Вычисляем расстояния от точки i до всех остальных точек и сохраняем их в хеш-таблице.
			- Используем квадрат расстояния вместо самого расстояния для избежания операции извлечения корня.

		Для каждого расстояния с count точек:
			- Количество бумерангов = count × (count - 1). Это потому что мы можем выбрать любые 2 точки из count точек на одинаковом расстоянии, и порядок имеет значение.
			- Добавляем количество бумерангов для текущей точки i к общему результату.

		Возвращаем общее количество бумерангов.

		Примеры:
		points = [[0,0],[1,0],[2,0]]
		Output: 2
		Explanation: The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]].

		TIME COMPLEXITY: O(n^2), где n - количество точек. Мы проходим по всем точкам и для каждой точки
		вычисляем расстояния до всех остальных точек.

		SPACE COMPLEXITY: O(n), где n - количество точек. В худшем случае все точки могут быть на одном расстоянии
		от точки i, поэтому нам нужно хранить информацию о каждом расстоянии в хеш-таблице.
	*/
	if len(points) < 3 {
		return 0
	}

	result := 0

	// Для каждой точки i рассматриваем её как центр бумеранга
	for i := 0; i < len(points); i++ {
		// Хеш-таблица для подсчета количества точек на каждом расстоянии от точки i
		distanceCount := make(map[int]int)

		// Вычисляем расстояния от точки i до всех остальных точек
		for j := 0; j < len(points); j++ {
			if i != j {
				// Используем квадрат расстояния для избежания операции извлечения корня
				// distance² = (x₁ - x₂)² + (y₁ - y₂)²
				distanceSquared := getDistanceSquared(points[i], points[j])
				distanceCount[distanceSquared]++
			}
		}

		// Для каждого расстояния с count точек:
		// количество бумерангов = count × (count - 1)
		// Это потому что мы можем выбрать любые 2 точки из count точек
		// на одинаковом расстоянии, и порядок имеет значение
		for _, count := range distanceCount {
			if count >= 2 {
				result += count * (count - 1)
			}
		}
	}

	return result
}

// getDistanceSquared вычисляет квадрат расстояния между двумя точками
// Математическая оптимизация: избегаем операции извлечения корня
func getDistanceSquared(point1, point2 []int) int {
	dx := point1[0] - point2[0]
	dy := point1[1] - point2[1]
	return dx*dx + dy*dy
}

func main() {
	// Тестовые случаи
	testCases := [][][]int{
		{{0, 0}, {1, 0}, {2, 0}}, // Expected: 2
		{{1, 1}, {2, 2}, {3, 3}}, // Expected: 2
		{{1, 1}},                 // Expected: 0
		{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}}, // Expected: 20
	}

	expected := []int{2, 2, 0, 20}

	for i, points := range testCases {
		result := numberOfBoomerangs(points)
		fmt.Printf("Test %d: points = %v\n", i+1, points)
		fmt.Printf("Result: %d, Expected: %d\n", result, expected[i])
		if result == expected[i] {
			fmt.Println("✅ PASS")
		} else {
			fmt.Println("❌ FAIL")
		}
		fmt.Println()
	}
}
