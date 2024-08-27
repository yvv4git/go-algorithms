package main

import (
	"fmt"
)

// Функция max возвращает максимальное значение из двух целых чисел.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Функция maxRotateFunction вычисляет максимальное значение функции F(k) для заданного массива A.
func maxRotateFunction(nums []int) int {
	/*
		METHOD: Dynamic programming with memoization
		Этот подход можно назвать итеративным с оптимизацией на основе предыдущего состояния или динамическим программированием с использованием мемоизации.
		Для решения задачи используется оптимизированный подход, основанный на наблюдении,
		что каждое следующее значение функции F(k) можно вычислить на основе предыдущего значения F(k-1) без необходимости пересчитывать всю сумму заново.
		Это позволяет значительно сократить количество вычислений.

		Оптимизация на основе предыдущего состояния.
		Оптимизация на основе предыдущего состояния означает, что мы используем уже вычисленное значение F(k-1) для вычисления нового значения F(k),
		что позволяет избежать повторных вычислений и сократить временную сложность

		Динамическое программирование с использованием мемоизации:
		Динамическое программирование — это метод решения сложных задач путём разбиения их на более простые подзадачи.
		Мемоизация — это техника запоминания результатов выполнения функций для предотвращения повторных вычислений.
		В данном случае, мы запоминаем результат предыдущей итерации (F(k-1)) и используем его для вычисления следующей итерации (F(k)),
		что делает этот подход похожим на динамическое программирование с использованием мемоизации.

		TIME COMPLEXITY:
		Временная сложность алгоритма составляет O(n), где n — длина массива nums.
		Это связано с тем, что мы выполняем один проход по массиву для вычисления суммы элементов и один проход для вычисления начального значения F(0).
		Затем мы выполняем еще один проход по массиву для вычисления значений F(k) для всех k от 1 до n-1. Таким образом, общее количество операций пропорционально длине массива.

		SPACE COMPLEXITY:
		Пространственная сложность алгоритма составляет O(1),
		так как мы используем фиксированное количество дополнительной памяти (переменные sum, F, maxF и несколько временных переменных) независимо от размера входного массива.
	*/
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Вычисляем сумму всех элементов массива nums.
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// Вычисляем начальное значение F(0).
	F := 0
	for i, num := range nums {
		F += i * num
	}

	// Инициализируем максимальное значение F(k) начальным значением F(0).
	maxF := F

	// Вычисляем значения F(k) для всех k от 1 до n-1.
	for k := 1; k < n; k++ {
		// Обновляем значение F(k) на основе предыдущего значения F(k-1).
		F = F + sum - n*nums[n-k]
		// Обновляем максимальное значение F(k).
		maxF = max(maxF, F)
	}

	return maxF
}

func main() {
	A := []int{4, 3, 2, 6}
	fmt.Println("Максимальное значение функции F(k):", maxRotateFunction(A))
}
