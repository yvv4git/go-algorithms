package main

import "fmt"

// Функция combine принимает два аргумента: n и k, и возвращает все возможные комбинации из k чисел из диапазона от 1 до n.
func combine(n int, k int) [][]int {
	/*
		METHOD: Backtracking
		TIME COMPLEXITY: O(2^n), так как для каждого числа i у нас есть два варианта действий (включить или не включить), и, следовательно, для n чисел у нас есть 2^n возможных комбинаций.
		SPACE COMPLEXITY: O(C(n, k)), где C - это биномиальный коэффициент, который дает количество сочетаний из n по k.
	*/
	var result [][]int                   // Инициализируем слайс для хранения результатов
	backtrack(n, k, 1, []int{}, &result) // Вызываем функцию backtrack для построения комбинаций
	return result                        // Возвращаем результат
}

// Функция backtrack используется для рекурсивного построения комбинаций.
// Аргументы:
// n - верхняя граница диапазона чисел
// k - количество чисел в комбинации
// start - начальное число для текущей комбинации
// current - текущая комбинация чисел
// result - указатель на слайс, в который будут добавляться комбинации
func backtrack(n int, k int, start int, current []int, result *[][]int) {
	// Если длина текущей комбинации равна k, то мы добавляем ее копию в результат
	if len(current) == k {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	// Перебираем все числа от start до n
	for i := start; i <= n; i++ {
		// Добавляем текущее число в комбинацию
		current = append(current, i)
		// Рекурсивно вызываем backtrack для оставшихся чисел
		backtrack(n, k, i+1, current, result)
		// Удаляем последнее число из комбинации
		current = current[:len(current)-1]
	}
}

func main() {
	n := 4
	k := 2
	combinations := combine(n, k)
	fmt.Println(combinations)
}
