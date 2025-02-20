//go:build ignore

package main

import (
	"fmt"
	"math"
	"sort"
)

// Структура для хранения координат точки
type Point struct {
	x, y int
}

// Функция для вычисления расстояния между двумя точками
func distance(p1, p2 Point) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

// Функция для нахождения минимального расстояния в объединенной части
func stripClosest(strip []Point, d float64) float64 {
	/*
		Задача: Найти минимальное расстояние среди точек, лежащих в полосе ширины d.

		Подход:
		1. Учитываем только точки, которые расположены близко друг к другу по оси Y.
		2. Сравниваем расстояния между каждой парой точек в полосе.
		3. Возвращаем минимальное из этих расстояний.

		ТАЙМ-КОМПЛЕКСНОСТЬ: O(n)
		СПЕЙС-КОМПЛЕКСНОСТЬ: O(n)
	*/
	n := len(strip)
	minDist := d

	// Сортируем точки по y-координате
	sort.Slice(strip, func(i, j int) bool {
		return strip[i].y < strip[j].y
	})

	// Сравниваем только близкие точки
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && float64(strip[j].y-strip[i].y) < minDist; j++ {
			minDist = math.Min(minDist, distance(strip[i], strip[j]))
		}
	}
	return minDist
}

// Функция для поиска минимального расстояния среди точек
func closestUtil(points []Point, left, right int) float64 {
	/*
		Задача: Найти минимальное расстояние среди точек в массиве, используя Divide and Conquer.

		Подход:
		1. Разделяем массив на две части.
		2. Рекурсивно находим минимальное расстояние для каждой части.
		3. Находим минимальное расстояние для точек, которые лежат по обе стороны от раздела.
		4. Объединяем результаты и находим минимальное расстояние для всего массива.

		ТАЙМ-КОМПЛЕКСНОСТЬ: O(n log n)
		СПЕЙС-КОМПЛЕКСНОСТЬ: O(n)
	*/
	if right-left <= 3 {
		// Для маленького количества точек используем наивный подход
		minDist := math.MaxFloat64
		for i := left; i <= right; i++ {
			for j := i + 1; j <= right; j++ {
				minDist = math.Min(minDist, distance(points[i], points[j]))
			}
		}
		return minDist
	}

	// Разделяем массив на две части
	mid := (left + right) / 2
	midPoint := points[mid]

	// Рекурсивный вызов для левой и правой части
	d1 := closestUtil(points, left, mid)
	d2 := closestUtil(points, mid+1, right)

	// Минимальное расстояние для двух частей
	d := math.Min(d1, d2)

	// Находим минимальное расстояние среди точек, которые пересекают раздел
	var strip []Point
	for i := left; i <= right; i++ {
		if math.Abs(float64(points[i].x-midPoint.x)) < d {
			strip = append(strip, points[i])
		}
	}

	// Возвращаем минимальное расстояние
	return math.Min(d, stripClosest(strip, d))
}

// Основная функция для нахождения минимального расстояния между точками
func closest(points []Point) float64 {
	// Сортируем точки по x-координатам
	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x
	})

	// Вызов рекурсивной функции
	return closestUtil(points, 0, len(points)-1)
}

func main() {
	// Пример: массив точек на плоскости
	points := []Point{
		{2, 3},
		{12, 30},
		{40, 50},
		{5, 1},
		{12, 10},
		{3, 4},
	}

	// Вызов функции для нахождения минимального расстояния
	result := closest(points)

	// Выводим результат
	fmt.Printf("Минимальное расстояние между двумя точками: %.4f\n", result)
}
