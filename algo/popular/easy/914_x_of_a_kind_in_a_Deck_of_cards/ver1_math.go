package main

import (
	"fmt"
)

// greatestCommonDivisor вычисляет наибольший общий делитель двух чисел
// Используется алгоритм Евклида, который работает за O(log(min(a, b)))
func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// hasGroupsSizeX проверяет, можно ли разделить колоду на группы с одинаковым размером
func hasGroupsSizeX(deck []int) bool {
	/*
		METHOD: Math
		- Используем математическую концепцию наибольшего общего делителя (НОД).
		- Подсчитываем частоту каждого значения карты и проверяем, можно ли разделить эти частоты на группы с одинаковым размером.
		- Если НОД всех частот больше или равен 2, то колоду можно разделить на группы.

		TIME COMPLEXITY: O(n)
		- Подсчет частот: O(n), где n — количество карт в колоде.
		- Вычисление НОД: O(k * log(max(counts))), где k — количество уникальных значений карт (максимум n).
		- Общая сложность: O(n).

		SPACE COMPLEXITY: O(n)
		- Используется мапа для хранения частот значений карт: O(n) в худшем случае (если все карты уникальны).
		- Дополнительная память для переменных: O(1).
		- Общая сложность: O(n).
	*/

	// Создаем мапу для подсчета частоты каждого значения карты
	countMap := make(map[int]int)
	for _, card := range deck {
		countMap[card]++
	}

	// Находим НОД всех частот
	var gcd int
	for _, count := range countMap {
		if gcd == 0 {
			gcd = count
		} else {
			gcd = greatestCommonDivisor(gcd, count)
			// Если НОД стал 1, то разделить на группы невозможно
			if gcd == 1 {
				return false
			}
		}
	}

	// Если НОД >= 2, то колоду можно разделить на группы
	return gcd >= 2
}

func main() {
	// Примеры тестов
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))       // true
	fmt.Println(hasGroupsSizeX([]int{1, 1, 2, 2, 2, 2}))             // true
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))       // false
	fmt.Println(hasGroupsSizeX([]int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2})) // true
}
