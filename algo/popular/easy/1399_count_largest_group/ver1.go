package main

import (
	"fmt"
	"strconv"
)

func countLargestGroup(n int) int {
	/*
		METHOD: Grouping
		- Группируем числа от 1 до n по сумме их цифр.
		- Находим максимальный размер группы.
		- Считаем количество групп с максимальным размером.

		TIME COMPLEXITY: O(n * m)
		- n: количество чисел от 1 до n.
		- m: количество цифр в числе (в худшем случае log10(n)).

		SPACE COMPLEXITY: O(n)
		- Хранение групп в карте, где каждая группа может содержать до n чисел.
	*/

	// Создаем карту для хранения групп
	groups := make(map[int][]int)

	// Заполняем группы
	for num := 1; num <= n; num++ {
		sum := digitSum(num)                   // Вычисляем сумму цифр числа
		groups[sum] = append(groups[sum], num) // Добавляем число в соответствующую группу
	}

	// Находим максимальный размер группы
	maxSize := 0
	for _, group := range groups {
		if len(group) > maxSize {
			maxSize = len(group)
		}
	}

	// Считаем количество групп с максимальным размером
	count := 0
	for _, group := range groups {
		if len(group) == maxSize {
			count++
		}
	}

	return count
}

// Функция для вычисления суммы цифр числа
func digitSum(num int) int {
	sum := 0
	for _, digit := range strconv.Itoa(num) { // Преобразуем число в строку и проходим по каждой цифре
		val, _ := strconv.Atoi(string(digit)) // Преобразуем цифру обратно в число
		sum += val                            // Добавляем к сумме
	}
	return sum
}

func main() {
	// Пример использования
	n := 13
	fmt.Println(countLargestGroup(n)) // Вывод: 4
}
