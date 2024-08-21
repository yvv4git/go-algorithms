package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Solution структура для хранения массива чисел
type Solution struct {
	nums []int
}

// Constructor инициализирует объект Solution с заданным массивом nums
func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

// Pick возвращает случайный индекс i такой, что nums[i] == target
func (s *Solution) Pick(target int) int {
	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Счетчик для количества вхождений target
	count := 0
	// Результат (индекс)
	result := -1

	// Проходим по массиву и ищем индексы, где nums[i] == target
	for i, num := range s.nums {
		if num == target {
			count++
			// Используем алгоритм резервуарной выборки
			// С вероятностью 1/count выбираем текущий индекс
			if rand.Intn(count) == 0 {
				result = i
			}
		}
	}

	return result
}

func main() {
	/*
		METHOD: Reservoir Sampling
		- Используется алгоритм резервуарной выборки (Reservoir Sampling) для выбора случайного индекса,
		  соответствующего заданному значению target, с равномерным распределением вероятности.

		TIME COMPLEXITY:
		- Конструктор: O(1) - просто сохраняет ссылку на массив.
		- Метод Pick: O(n), где n - длина массива nums. Требуется один проход по массиву для поиска всех вхождений target.

		SPACE COMPLEXITY:
		- Конструктор: O(n), где n - длина массива nums, так как сохраняет массив в поле структуры.
		- Метод Pick: O(1) - используется фиксированное количество дополнительной памяти (переменные count и result).
	*/
	nums := []int{1, 2, 3, 3, 3}
	solution := Constructor(nums)

	// Пример вызова метода pick
	fmt.Println(solution.Pick(3)) // Должен вернуть индекс 2, 3 или 4 с равной вероятностью
	fmt.Println(solution.Pick(1)) // Должен вернуть индекс 0
}
