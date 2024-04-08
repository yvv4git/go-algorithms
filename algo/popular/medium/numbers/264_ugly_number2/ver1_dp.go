package main

import "fmt"

// Функция для нахождения n-го "ugly number"
func nthUglyNumber(n int) int {
	/*
		METHOD: Dynamic programming
		Для решения задачи используется динамическое программирование. Мы создаем массив для хранения "ugly numbers" и три указателя для 2, 3 и 5.
		Затем мы находим следующее "ugly number" путем выбора минимального из возможных следующих "ugly numbers"
		и увеличиваем соответствующий указатель.
		Таким образом, мы строим последовательность "ugly numbers" по мере необходимости.

		TIME COMPLEXITY: O(n), где n - номер "ugly number", который мы ищем. Мы проходим по всем "ugly numbers" от 1 до n, поэтому временная сложность линейная.

		SPACE COMPLEXITY: O(n), где n - номер "ugly number". Мы храним все "ugly numbers" в массиве, поэтому пространственная сложность также линейная.
	*/
	// Инициализируем слайс для хранения "ugly numbers"
	ugly := make([]int, n)
	ugly[0] = 1 // Первое "ugly number" всегда 1

	// Инициализируем три указателя для 2, 3 и 5
	i2, i3, i5 := 0, 0, 0

	// Находим следующие "ugly numbers"
	for i := 1; i < n; i++ {
		// Выбираем минимальное из следующих возможных "ugly numbers"
		nextUgly := min(ugly[i2]*2, ugly[i3]*3, ugly[i5]*5)
		ugly[i] = nextUgly

		// Увеличиваем указатели, если они указывают на числа, которые мы только что добавили
		if nextUgly == ugly[i2]*2 {
			i2++
		}
		if nextUgly == ugly[i3]*3 {
			i3++
		}
		if nextUgly == ugly[i5]*5 {
			i5++
		}
	}

	// Возвращаем последнее "ugly number"
	return ugly[n-1]
}

// Функция для нахождения минимального числа из трех
func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}

func main() {
	fmt.Println(nthUglyNumber(10)) // Выводит: 12
}
