package main

import "strconv"

// Функция isAdditiveNumberV3 проверяет, является ли числовое строковое представление числа additive.
// Временная сложность: O(n^2), где n - длина строки num.
// Пространственная сложность: O(n), где n - длина строки num.
func isAdditiveNumberV3(num string) bool {
	/*
		METHOD: Recursion
		TIME COMPLEXITY: O(n^2)
		SPACE COMPLEXITY:  O(n)
	*/
	// Инициализация левой границы первого числа.
	firstL := 0
	// Цикл по всем возможным правым границам первого числа.
	for firstR := 0; firstR < len(num)/2; firstR++ {
		// Инициализация левой границы второго числа.
		secondL := firstR + 1
		// Цикл по всем возможным правым границам второго числа.
		for secondR := secondL; secondR < len(num)-1; secondR++ {
			// Вызов вспомогательной функции для проверки, является ли число additive.
			if helper(num, firstL, firstR, secondL, secondR, true) {
				// Если число additive, возвращаем true.
				return true
			}
		}
	}

	// Если число не является additive, возвращаем false.
	return false
}

// Вспомогательная функция helper проверяет, является ли число additive.
// Временная сложность: O(n), где n - длина строки num.
// Пространственная сложность: O(n), где n - длина строки num.
func helper(num string, firstL int, firstR int, secondL int, secondR int, firstIteration bool) bool {
	// Если это не первая итерация и правая граница второго числа - последний символ строки, возвращаем true.
	if !firstIteration && secondR == len(num)-1 {
		return true
	}

	// Если первое число содержит лидирующие нули или второе число содержит лидирующие нули, возвращаем false.
	if firstL != firstR && num[firstL] == '0' || secondL != secondR && num[secondL] == '0' {
		return false
	}

	// Преобразование строки в числа.
	first, _ := strconv.Atoi(num[firstL : firstR+1])
	second, _ := strconv.Atoi(num[secondL : secondR+1])
	// Вычисление суммы первого и второго числа.
	third := first + second

	// Преобразование суммы в строку.
	thirdStr := strconv.Itoa(third)
	// Вычисление границ третьего числа.
	thirdL := secondR + 1
	thirdR := thirdL + len(thirdStr) - 1
	// Если третье число не соответствует ожидаемому, возвращаем false.
	if thirdR >= len(num) || num[thirdL:thirdR+1] != thirdStr {
		return false
	}

	// Рекурсивный вызов вспомогательной функции для проверки оставшейся части строки.
	return helper(num, secondL, secondR, thirdL, thirdR, false)
}
