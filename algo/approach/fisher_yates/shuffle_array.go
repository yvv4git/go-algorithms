package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для перемешивания массива с использованием алгоритма Fisher-Yates
func shuffleArray(arr []int) []int {
	/*
		METHOD: Fisher-Yates shuffle (Knuth shuffle)
		TIME COMPLEXITY: O(n) - один проход по массиву длины n
		SPACE COMPLEXITY: O(1) - дополнительная память не используется (перемешивание выполняется in-place)
	*/
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)           // Генерация случайного индекса
		arr[i], arr[j] = arr[j], arr[i] // Обмен элементов
	}
	return arr
}

// Задача заключается в том, чтобы перемешать массив с использованием алгоритма Fisher-Yates(Knuth shuffle).
func main() {
	// Исходный массив
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Перемешивание массива
	shuffledNums := shuffleArray(nums)

	// Вывод перемешанного массива
	fmt.Println("Перемешанный массив:", shuffledNums)
}
