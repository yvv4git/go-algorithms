package main

import (
	"fmt"
)

// Функция sieveOfEratosthenes реализует алгоритм "Решето Эратосфена"
// для нахождения всех простых чисел до заданного целого числа n.
func sieveOfEratosthenes(n int) []bool {
	// Создаем массив nums размером n+1, где каждый элемент изначально равен false.
	// false означает, что число не отмечено как составное.
	nums := make([]bool, n+1)

	// Начинаем с первого простого числа (2) и проходим по всем числам до n.
	for i := 2; i <= n; i++ {
		// Если число i не отмечено как составное, то оно простое.
		if !nums[i] {
			// Отмечаем все кратные числа i как составные.
			// Начинаем с i*i, так как все предыдущие кратные уже были отмечены.
			for j := i * i; j <= n; j += i {
				nums[j] = true
			}
		}
	}

	// Возвращаем массив nums, где false означает простое число.
	return nums
}

func main() {
	// Задаем значение n, до которого будем искать простые числа.
	n := 30

	// Вызываем функцию sieveOfEratosthenes для нахождения простых чисел до n.
	nums := sieveOfEratosthenes(n)

	// Выводим все простые числа до n.
	fmt.Printf("Простые числа до %d:\n", n)
	for i := 2; i <= n; i++ {
		// Если число i не отмечено как составное, то оно простое.
		if !nums[i] {
			fmt.Printf("%d ", i)
		}
	}
}