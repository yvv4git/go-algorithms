package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Solution структура для хранения исходного и перемешанного массивов
type Solution struct {
	orig, shuf []int      // Исходный и перемешанный массивы
	r          *rand.Rand // Генератор случайных чисел
}

// Constructor инициализирует объект Solution с заданным массивом
// METHOD: Инициализация объекта с копированием исходного массива и созданием генератора случайных чисел
// TIME COMPLEXITY: O(n) - копирование массива
// SPACE COMPLEXITY: O(n) - хранение двух копий массива
func Constructor(arr []int) Solution {
	s := Solution{}
	s.orig = make([]int, len(arr))
	s.shuf = make([]int, len(arr))
	copy(s.orig, arr)                                     // Копирование исходного массива
	s.r = rand.New(rand.NewSource(time.Now().UnixNano())) // Инициализация генератора случайных чисел
	return s
}

// Reset возвращает исходный массив
// METHOD: Возвращение исходного массива
// TIME COMPLEXITY: O(1) - возвращение ссылки на исходный массив
// SPACE COMPLEXITY: O(1) - дополнительная память не используется
func (s *Solution) Reset() []int {
	return s.orig
}

// Shuffle перемешивает массив случайным образом и возвращает его
// METHOD: Перемешивание массива с использованием алгоритма Фишера-Йетса
// TIME COMPLEXITY: O(n) - один проход по массиву
// SPACE COMPLEXITY: O(1) - дополнительная память не используется
func (s *Solution) Shuffle() []int {
	copy(s.shuf, s.orig) // Копирование исходного массива в перемешанный
	for i := range s.shuf {
		j := i + s.r.Intn(len(s.shuf)-i)            // Генерация случайного индекса
		s.shuf[i], s.shuf[j] = s.shuf[j], s.shuf[i] // Обмен элементов
	}
	return s.shuf
}

func main() {
	/*
		METHOD: Fisher-Yates (aka Knuth shuffle)
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
