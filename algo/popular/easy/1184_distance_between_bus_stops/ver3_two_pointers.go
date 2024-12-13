//go:build ignore

package main

import (
	"fmt"
)

// Функция для вычисления минимального расстояния между остановками
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	/*
		METHOD:
		- Используем два указателя для вычисления расстояний по часовой и против часовой стрелке.
		- Один указатель идет от start до destination (по часовой стрелке).
		- Другой указатель идет от destination до start (против часовой стрелки).

		TIME COMPLEXITY:
		- O(n): где n — количество остановок.
		- Один проход по массиву для вычисления расстояний.

		SPACE COMPLEXITY:
		- O(1): используется только постоянное количество дополнительной памяти.
	*/

	// Если start > destination, меняем их местами
	if start > destination {
		start, destination = destination, start
	}

	// Вычисляем расстояние по часовой стрелке
	clockwiseDistance := 0
	for i := start; i < destination; i++ {
		clockwiseDistance += distance[i]
	}

	// Вычисляем расстояние против часовой стрелки
	counterClockwiseDistance := 0
	// Указатель против часовой стрелки идет от destination до start
	for i := destination; i < len(distance); i++ {
		counterClockwiseDistance += distance[i]
	}
	for i := 0; i < start; i++ {
		counterClockwiseDistance += distance[i]
	}

	// Возвращаем минимальное из двух расстояний
	if clockwiseDistance < counterClockwiseDistance {
		return clockwiseDistance
	}
	return counterClockwiseDistance
}

func main() {
	// Пример входных данных
	distance := []int{1, 2, 3, 4}
	start := 0
	destination := 2

	// Вызов функции и вывод результата
	result := distanceBetweenBusStops(distance, start, destination)
	fmt.Printf("Минимальное расстояние между остановками %d и %d: %d\n", start, destination, result)

	// Другой пример
	distance2 := []int{7, 10, 1, 12, 11, 14, 5, 0}
	start2 := 7
	destination2 := 2

	result2 := distanceBetweenBusStops(distance2, start2, destination2)
	fmt.Printf("Минимальное расстояние между остановками %d и %d: %d\n", start2, destination2, result2)
}
