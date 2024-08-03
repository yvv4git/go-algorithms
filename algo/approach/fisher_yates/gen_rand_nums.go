package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для генерации случайных чисел в заданном диапазоне без повторений
func generateRandomNumbers(min, max int) []int {
	if min >= max {
		panic("min должно быть меньше max")
	}

	// Создаем массив чисел в заданном диапазоне
	arr := make([]int, max-min+1)
	for i := range arr {
		arr[i] = min + i
	}

	// Перемешиваем массив с использованием алгоритма Fisher-Yates
	rand.Seed(time.Now().UnixNano())
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func main() {
	// Генерация случайных чисел в диапазоне от 1 до 10
	randomNumbers := generateRandomNumbers(1, 10)

	// Вывод результата
	fmt.Println("Сгенерированные случайные числа:", randomNumbers)
}
