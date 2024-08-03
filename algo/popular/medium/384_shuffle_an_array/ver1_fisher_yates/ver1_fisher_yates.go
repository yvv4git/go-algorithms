package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Solution структура для хранения исходного массива и текущего состояния
type Solution struct {
	original []int // Исходный массив
	current  []int // Текущее состояние массива
}

// Constructor инициализирует объект Solution с заданным массивом
// METHOD: Инициализация объекта с копированием исходного массива
// TIME COMPLEXITY: O(n) - копирование массива
// SPACE COMPLEXITY: O(n) - хранение двух копий массива
func Constructor(nums []int) Solution {
	original := make([]int, len(nums))
	copy(original, nums) // Копирование исходного массива
	return Solution{
		original: original,
		current:  nums,
	}
}

// Reset возвращает исходный массив
// METHOD: Восстановление исходного массива путем копирования
// TIME COMPLEXITY: O(n) - копирование массива
// SPACE COMPLEXITY: O(1) - дополнительная память не используется
func (this *Solution) Reset() []int {
	copy(this.current, this.original) // Копирование исходного массива в текущее состояние
	return this.current
}

// Shuffle перемешивает массив случайным образом и возвращает его
// METHOD: Перемешивание массива с использованием алгоритма Фишера-Йетса
// TIME COMPLEXITY: O(n) - один проход по массиву
// SPACE COMPLEXITY: O(1) - дополнительная память не используется
func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	for i := len(this.current) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)                                               // Генерация случайного индекса
		this.current[i], this.current[j] = this.current[j], this.current[i] // Обмен элементов
	}
	return this.current
}

func main() {
	/*
		METHOD: Fisher-Yates shuffle(Фишера-Йетса) алгоритм
		TIME COMPLEXITY: O(n) - вызовы методов Reset и Shuffle
		SPACE COMPLEXITY: O(n) - хранение массива и объекта Solution
	*/
	nums := []int{1, 2, 3, 4, 5}
	solution := Constructor(nums)

	fmt.Println("Original array:", solution.Reset())
	fmt.Println("Shuffled array:", solution.Shuffle())
	fmt.Println("Reset array:", solution.Reset())
	fmt.Println("Shuffled array again:", solution.Shuffle())
}
