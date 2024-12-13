//go:build ignore

package main

import (
	"fmt"
)

// Функция для вычисления минимального расстояния между остановками
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	/*
		METHOD: Prefix Sum
		- Используем префиксные суммы (Prefix Sum) для быстрого вычисления расстояний.
		- Префиксная сумма prefixSum[i] — это сумма расстояний от начала маршрута до i-й остановки.

		TIME COMPLEXITY:
		- O(n) для вычисления префиксных сумм.
		- O(1) для вычисления расстояний между остановками.

		SPACE COMPLEXITY:
		- O(n) для хранения массива префиксных сумм.
	*/

	// Если start > destination, меняем их местами
	if start > destination {
		start, destination = destination, start
	}

	// Создаем массив префиксных сумм
	n := len(distance)
	prefixSum := make([]int, n+1) // +1 для удобства
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + distance[i-1]
	}

	// Расстояние по часовой стрелке
	clockwiseDistance := prefixSum[destination] - prefixSum[start]

	// Общее расстояние
	totalDistance := prefixSum[n]

	// Расстояние против часовой стрелки
	counterClockwiseDistance := totalDistance - clockwiseDistance

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
