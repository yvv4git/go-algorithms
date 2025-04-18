package main

import "fmt"

// Функция tribonacci принимает целое число n и возвращает N-е число Трибоначчи.
func tribonacci(n int) int {
	/*
		METHOD: Iterative Approach
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)

		Описание метода:
		- Используем итеративный подход для вычисления N-го числа Трибоначчи.
		- Начинаем с базовых значений T(0) = 0, T(1) = 1, T(2) = 1.
		- В цикле от 3 до n вычисляем следующее число Трибоначчи как сумму трех предыдущих чисел.
		- Обновляем значения переменных a, b и c для хранения последних трех чисел Трибоначчи.
		- После завершения цикла, переменная c содержит значение T(n).

		Временная сложность:
		- O(n) - так как мы выполняем n итераций в цикле.

		Пространственная сложность:
		- O(1) - так как мы используем фиксированное количество переменных (a, b, c) независимо от значения n.
	*/
	// Базовые случаи: если n равно 0, 1 или 2, возвращаем соответствующее значение.
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	// Инициализируем три переменные для хранения первых трех чисел Трибоначчи.
	a, b, c := 0, 1, 1

	// Итерируем от 3 до n, обновляя значения a, b и c.
	for i := 3; i <= n; i++ {
		// Вычисляем следующее число Трибоначчи как сумму трех предыдущих.
		next := a + b + c
		// Обновляем значения a, b и c для следующей итерации.
		a, b, c = b, c, next
	}

	// Возвращаем значение c, которое теперь содержит N-е число Трибоначчи.
	return c
}

func main() {
	// Пример использования функции tribonacci.
	n := 25
	result := tribonacci(n)
	fmt.Printf("T(%d) = %d\n", n, result)
}
