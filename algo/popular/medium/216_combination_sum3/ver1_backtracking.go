package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	/*
		METHOD: Backtracking
		Для решения этой задачи мы будем использовать рекурсивный подход с методом обратного отслеживания (backtracking).
		Метод обратного отслеживания - это алгоритмический подход, который используется для решения задач,
		которые требуют генерации всех возможных комбинаций или перестановок.

		TIME COMPLEXITY: O(9^k), где k - количество чисел в комбинации, так как нам нужно проверить все возможные комбинации из 9 чисел.

		SPACE COMPLEXITY: O(9^k), так как в худшем случае мы можем поместить в ответ все возможные комбинации.
	*/
	var result [][]int
	backtrack(&result, []int{}, k, n, 1)
	return result
}

func backtrack(result *[][]int, current []int, k int, n int, start int) {
	// Если сумма текущей комбинации равна n и количество чисел равно k,
	// то добавляем комбинацию в результат
	if n == 0 && len(current) == k {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	// Если сумма текущей комбинации больше n или количество чисел больше k,
	// то прекращаем обход
	if n < 0 || len(current) > k {
		return
	}

	// Проходим по всем возможным числам от start до 9
	for i := start; i <= 9; i++ {
		// Добавляем число в текущую комбинацию
		current = append(current, i)
		// Рекурсивно вызываем backtrack для оставшейся суммы и следующего числа
		backtrack(result, current, k, n-i, i+1)
		// Удаляем число из текущей комбинации для проверки следующей комбинации
		current = current[:len(current)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(3, 7)) // Вывод: [[1 2 4]]
}
