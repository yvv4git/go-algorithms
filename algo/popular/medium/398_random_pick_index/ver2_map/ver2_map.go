package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Solution структура для хранения карты индексов
type Solution struct {
	indexMap map[int][]int
}

// Constructor инициализирует объект Solution с заданным массивом nums
func Constructor(nums []int) Solution {
	indexMap := make(map[int][]int)
	for i, num := range nums {
		indexMap[num] = append(indexMap[num], i)
	}
	return Solution{indexMap: indexMap}
}

// Pick возвращает случайный индекс i такой, что nums[i] == target
func (this *Solution) Pick(target int) int {
	indices := this.indexMap[target]
	return indices[rand.Intn(len(indices))]
}

func main() {
	/*
		METHOD: Pre-indexing
		- Используется для ускорения поиска индексов элементов в массиве.
		- Создается карта (map), где ключ — это значение элемента, а значение — список индексов, где этот элемент встречается.
		- При вызове метода Pick, мы просто обращаемся к этой карте, чтобы получить список индексов и выбрать случайный индекс из него.

		TIME COMPLEXITY:
		- Конструктор: O(n), где n - длина массива nums. Требуется один проход по массиву для создания карты индексов.
		- Метод Pick: O(1) в среднем, так как доступ к элементу в карте и выбор случайного индекса из списка выполняются за константное время.

		SPACE COMPLEXITY:
		- Конструктор: O(n), где n - длина массива nums. Создается карта, в которой хранятся все индексы элементов.
		- Метод Pick: O(1), так как используется фиксированное количество дополнительной памяти.
	*/
	rand.Seed(time.Now().UnixNano()) // Инициализируем генератор случайных чисел

	nums := []int{1, 2, 3, 3, 3}
	solution := Constructor(nums)

	// Пример вызова метода pick
	fmt.Println(solution.Pick(3)) // Должен вернуть индекс 2, 3 или 4 с равной вероятностью
	fmt.Println(solution.Pick(1)) // Должен вернуть индекс 0
}
