package main

import (
	"fmt"
	"strconv"
)

// Функция для вычисления факториала
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// Функция для получения k-й перестановки из n чисел
func getPermutation(n int, k int) string {
	/*
		METHOD: Math
		TIME COMPLEXITY: O(n^2), где n - количество чисел для перестановки. Это обусловлено двойным циклом, который проходит по всем числам и удаляет выбранное число из слайса.
		SPACE COMPLEXITY: O(n), так как мы храним слайсы для чисел и факториалов, а также строку для результата.
	*/
	// Создаем слайс для хранения доступных чисел
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = i + 1
	}

	// Вычисляем факториалы чисел от 0 до n
	factorials := make([]int, n+1)
	for i := 0; i <= n; i++ {
		factorials[i] = factorial(i)
	}

	// Инициализируем результирующую строку
	result := ""
	k-- // Уменьшаем k на 1, чтобы индексы соответствовали нумерации с 0

	// Пока есть числа для выбора
	for n > 0 {
		// Вычисляем индекс текущего числа
		index := k / factorials[n-1]
		// Добавляем в результат выбранное число
		result += strconv.Itoa(numbers[index])
		// Удаляем выбранное число из доступных
		numbers = append(numbers[:index], numbers[index+1:]...)
		// Обновляем k и n
		k %= factorials[n-1]
		n--
	}

	return result
}

func main() {
	// Пример использования
	n := 3
	k := 3
	permutation := getPermutation(n, k)
	fmt.Printf("The %dth permutation sequence of %d numbers is: %s\n", k, n, permutation)
}
