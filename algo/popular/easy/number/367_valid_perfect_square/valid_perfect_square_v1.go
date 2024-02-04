package _67_valid_perfect_square

import "fmt"

func isPerfectSquareV1(num int) bool {
	/*
		METHOD: Binary search.
		TIME COMPLEXITY: O(log n)
		Space complexity:

		Асимптотическая сложность данного алгоритма - O(log n).
		Это связано с использованием бинарного поиска, который имеет логарифмическую сложность O(log n).
		В каждой итерации цикла while, мы делим поиск пространства пополам, поэтому мы можем ёго выполнить за log n шагов.

		Алгоритм isPerfectSquare проверяет, является ли число num полным квадратом.
		1. Если num меньше 2, то оно является полным квадратом, поэтому функция возвращает true.
		Полный квадрат - это целое число, которое является квадратом другого целого числа.
		Для чисел меньше 2, единственным полным квадратом является 1 (11 = 1), а 0 (00 = 0) также является полным квадратом.
		Поэтому, если число меньше 2, то оно является полным квадратом.
		Иными словами, квадратом является целое число, квадратный корень из которого извлекается нацело.
		Например, 9 = 3 * 3.

		2. Если num больше или равно 2, то алгоритм начинает бинарный поиск.
		Он инициализирует две переменные left и right, равные 2 и num/2 соответственно.

		3. Алгоритм входит в цикл, который продолжается, пока left меньше или равно right.

		4. Внутри цикла, алгоритм находит среднее значение между left и right, которое мы назовем mid.

		5. Алгоритм проверяет, является ли квадрат mid равным num. Если да, то num является полным квадратом, поэтому функция возвращает true.

		6. Если квадрат mid больше num, то алгоритм устанавливает right равным mid - 1, чтобы исключить правую половину пространства поиска.

		7. Если квадрат mid меньше num, то алгоритм устанавливает left равным mid + 1, чтобы исключить левую половину пространства поиска.

		8. Если алгоритм не находит полного квадрата num, то он выходит из цикла и возвращает false.

		В результате, этот алгоритм использует бинарный поиск для поиска полного квадрата num,
		что позволяет ему работать быстро, даже для больших чисел.
	*/
	if num < 2 {
		// Числа 1 и 0 являются полными квадратами, так как 1 * 1 = 1, 0 * 0 = 0.
		// Т.е. корень из этих чисел возвращает целое число. Т.е. из них квадратный корень извлекается нацело.
		return true
	}

	left, right := 2, num/2 // Бинарный поиск делит данные на пополам. Это гарантирует, что мы не будем искать число больше num.
	for left <= right {
		mid := left + (right-left)/2
		sqr := mid * mid
		if sqr == num {
			return true
		} else if sqr > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func isPerfectSquareV1Research(num int) bool {
	/*
		METHOD: Binary search.
		TIME COMPLEXITY: O(log n)
		Space complexity:

		Асимптотическая сложность данного алгоритма - O(log n).
		Это связано с использованием бинарного поиска, который имеет логарифмическую сложность O(log n).
		В каждой итерации цикла while, мы делим поиск пространства пополам, поэтому мы можем ёго выполнить за log n шагов.

		Алгоритм isPerfectSquare проверяет, является ли число num полным квадратом.
		1. Если num меньше 2, то оно является полным квадратом, поэтому функция возвращает true.
		Полный квадрат - это целое число, которое является квадратом другого целого числа.
		Для чисел меньше 2, единственным полным квадратом является 1 (11 = 1), а 0 (00 = 0) также является полным квадратом.
		Поэтому, если число меньше 2, то оно является полным квадратом.
		Иными словами, квадратом является целое число, квадратный корень из которого извлекается нацело.
		Например, 9 = 3 * 3.

		2. Если num больше или равно 2, то алгоритм начинает бинарный поиск.
		Он инициализирует две переменные left и right, равные 2 и num/2 соответственно.

		3. Алгоритм входит в цикл, который продолжается, пока left меньше или равно right.

		4. Внутри цикла, алгоритм находит среднее значение между left и right, которое мы назовем mid.

		5. Алгоритм проверяет, является ли квадрат mid равным num. Если да, то num является полным квадратом, поэтому функция возвращает true.

		6. Если квадрат mid больше num, то алгоритм устанавливает right равным mid - 1, чтобы исключить правую половину пространства поиска.

		7. Если квадрат mid меньше num, то алгоритм устанавливает left равным mid + 1, чтобы исключить левую половину пространства поиска.

		8. Если алгоритм не находит полного квадрата num, то он выходит из цикла и возвращает false.

		В результате, этот алгоритм использует бинарный поиск для поиска полного квадрата num,
		что позволяет ему работать быстро, даже для больших чисел.
	*/
	if num < 2 {
		// Числа 1 и 0 являются полными квадратами, так как 1 * 1 = 1, 0 * 0 = 0.
		// Т.е. корень из этих чисел возвращает целое число. Т.е. из них квадратный корень извлекается нацело.
		return true
	}
	fmt.Printf("=>num=%v \n", num)

	left, right := 2, num/2 // Бинарный поиск делит данные на пополам. Это гарантирует, что мы не будем искать число больше num.
	for left <= right {
		mid := left + (right-left)/2
		fmt.Printf("->mid=%v\n", mid) // mid = 5, 3, 4
		sqr := mid * mid
		if sqr == num {
			return true
		} else if sqr > num {
			fmt.Printf("==>right=%v \n", right)
			right = mid - 1
		} else {
			fmt.Printf("==>left=%v \n", left)
			left = mid + 1
		}
	}

	return false
}
