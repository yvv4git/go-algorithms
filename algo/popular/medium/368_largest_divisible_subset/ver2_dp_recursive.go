package main

import (
	"fmt"
	"sort"
)

// Функция largestDivisibleSubsetV2 принимает массив целых чисел и возвращает наибольшее делимое подмножество.
// Временная сложность: O(n^2), где n - длина массива nums.
// Пространственная сложность: O(n^2), где n - длина массива nums.
func largestDivisibleSubsetV2(nums []int) []int {
	/*
		METHOD: Memoization (динамическое программирование с использованием мемоизации)
		Метод мемоизации (memoization) удобно выбирать в тех случаях, когда алгоритм решает задачу с использованием рекурсии
		и при этом многократно вычисляет одни и те же подзадачи.
		Мемоизация позволяет сохранить результаты этих подзадач и использовать их повторно, что значительно сокращает время выполнения алгоритма.

		TIME COMPLEXITY: O(n^2), где n - длина массива nums.
		- Сортировка массива занимает O(n log n) времени.
		- Для каждого элемента мы вызываем функцию dp, которая может вызываться рекурсивно для каждого предыдущего элемента, что даёт O(n^2) в худшем случае.
		- Общая временная сложность составляет O(n log n) + O(n^2), что асимптотически эквивалентно O(n^2).

		SPACE COMPLEXITY: O(n^2), где n - длина массива nums.
		- Мы используем карту m для мемоизации результатов, которая может содержать до n записей, каждая из которых содержит массив длины до n.
		- Таким образом, общая пространственная сложность составляет O(n^2).
	*/
	m := map[int][]int{} // Карта для мемоизации результатов
	sort.Ints(nums)      // Сортируем массив для упрощения проверки делимости

	result := []int{}
	for i := range nums {
		cur := dp(i, nums, m) // Рекурсивно находим наибольшее делимое подмножество для текущего элемента
		if len(cur) > len(result) {
			result = append([]int{}, cur...) // Копируем результат, если он больше текущего наибольшего подмножества
		}
	}
	return result
}

// Функция dp рекурсивно находит наибольшее делимое подмножество для элемента с индексом i.
func dp(i int, nums []int, m map[int][]int) []int {
	if res, ok := m[i]; ok {
		return res // Возвращаем мемоизированный результат, если он существует
	}

	result := []int{}
	// Проверяем все предыдущие элементы
	for k := 0; k < i; k++ {
		if nums[i]%nums[k] == 0 { // Если текущий элемент делится на предыдущий
			cur := dp(k, nums, m) // Рекурсивно находим подмножество для предыдущего элемента
			if len(cur) > len(result) {
				result = append([]int{}, cur...) // Копируем результат, если он больше текущего подмножества
			}
		}
	}

	result = append(result, nums[i]) // Добавляем текущий элемент к результату
	m[i] = result                    // Мемоизируем результат
	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 8}
	fmt.Println(largestDivisibleSubsetV2(nums)) // Вывод: [1, 2, 4, 8]
}
