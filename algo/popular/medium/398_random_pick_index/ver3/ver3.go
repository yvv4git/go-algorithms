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

// NewSolution конструктор для инициализации объекта Solution
func NewSolution(nums []int) *Solution {
	return &Solution{nums: nums}
}

// Pick возвращает случайный индекс i такой, что nums[i] == target
func (this *Solution) Pick(target int) int {
	rand.Seed(time.Now().UnixNano())
	var indices []int

	// Собираем все индексы, где nums[i] == target
	for i, num := range this.nums {
		if num == target {
			indices = append(indices, i)
		}
	}

	// Возвращаем случайный индекс из собранных
	return indices[rand.Intn(len(indices))]
}

func main() {
	/*
		METHOD: Reservoir Sampling


		TIME COMPLEXITY:


		SPACE COMPLEXITY:

	*/
	rand.Seed(time.Now().UnixNano())

	nums := []int{1, 2, 3, 3, 3}
	solution := NewSolution(nums)

	// Пример вызова метода pick
	fmt.Println(solution.Pick(3)) // Должен вернуть индекс 2, 3 или 4 с равной вероятностью
	fmt.Println(solution.Pick(1)) // Должен вернуть индекс 0
}
