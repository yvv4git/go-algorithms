package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	/*
		METHOD: Math & bitwise
		TIME COMPLEXITY: O(log n), где n - абсолютное значение делимого.
		Это обусловлено тем, что в каждой итерации мы уменьшаем делимое вдвое, что позволяет нам уменьшать количество итераций.
		SPACE COMPLEXITY: O(1), так как мы используем фиксированное количество переменных, не зависящих от размера входных данных.

		Алгоритм основан на последовательном вычитании удвоенного делителя из делимого.
		Это позволяет уменьшать делимое быстрее, чем просто на 1, и увеличивать результат на соответствующий множитель.
	*/
	// Обработка крайних случаев
	if divisor == 0 {
		panic("Division by zero is not allowed")
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	// Определение знака результата
	sign := -1
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		sign = 1
	}

	// Используем модули для работы с положительными числами
	dividend = abs(dividend)
	divisor = abs(divisor)

	// Инициализация результата
	result := 0

	// Пока делимое больше или равно делителю
	for dividend >= divisor {
		// Инициализация счетчика для подсчета количества вычитаний
		temp := divisor
		multiple := 1
		// Пока делимое больше или равно удвоенному делителю
		for dividend >= (temp << 1) {
			// Умножаем делитель на 2 и увеличиваем счетчик в 2 раза
			temp <<= 1
			multiple <<= 1
		}
		// Вычитаем из делимого полученное значение и увеличиваем результат на соответствующий множитель
		dividend -= temp
		result += multiple
	}

	// Возвращаем результат с учетом знака
	return sign * result
}

// Функция для нахождения абсолютного значения
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(divide(10, 3)) // Вывод: 3
	fmt.Println(divide(7, -3)) // Вывод: -2
}
