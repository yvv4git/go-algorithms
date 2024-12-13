//go:build ignore

package main

import (
	"fmt"
)

// Функция для вычисления минимального расстояния между остановками
func distanceBetweenBusStops(distance []int, start int, destination int) int {
	/*
		METHOD: Linear search
		1. Если начальная остановка (start) больше конечной (destination), меняем их местами.
		2. Вычисляем общее расстояние по всему маршруту, суммируя все элементы массива distance.
		3. Вычисляем расстояние по часовой стрелке от start до destination.
		4. Расстояние против часовой стрелки будет равно общему расстоянию минус расстояние по часовой стрелке.
		5. Возвращаем минимальное из двух расстояний (по часовой стрелке и против часовой стрелки).

		TIME COMPLEXITY:
		- O(n): где n — количество остановок (длина массива distance).
		  - Один проход по массиву для вычисления общего расстояния: O(n).
		  - Один проход по массиву для вычисления расстояния по часовой стрелке: O(n).
		  - Итого: O(n) + O(n) = O(n).

		SPACE COMPLEXITY:
		- O(1): используется только постоянное количество дополнительной памяти (переменные totalDistance, clockwiseDistance и counterClockwiseDistance).
	*/
	// Если начальная остановка больше конечной, меняем их местами
	if start > destination {
		start, destination = destination, start
	}

	// Вычисляем общее расстояние по всему маршруту
	totalDistance := 0
	for _, dist := range distance {
		totalDistance += dist
	}

	// Вычисляем расстояние по часовой стрелке от start до destination
	clockwiseDistance := 0
	for i := start; i < destination; i++ {
		clockwiseDistance += distance[i]
	}

	// Расстояние против часовой стрелке будет равно общему расстоянию минус расстояние по часовой стрелке
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
