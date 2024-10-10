//go:build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для нахождения k-го наименьшего элемента в массиве
func quickselectV2(nums []int, k int) int {
	/*
		METHOD: Quickselect
		Объяснение метода: Quickselect — это алгоритм выбора, основанный на принципе "разделяй и властвуй".
		Он использует случайный опорный элемент для разделения массива на части и рекурсивно находит k-й наименьший элемент.

		TIME COMPLEXITY: В среднем O(n), в худшем случае O(n^2)
		Объяснение временной сложности: В среднем алгоритм работает за линейное время, так как каждый раз массив делится на две примерно равные части.
		В худшем случае, когда опорный элемент выбирается неудачно, временная сложность может достигать O(n^2).

		SPACE COMPLEXITY: O(n)
		Объяснение пространственной сложности: Алгоритм использует дополнительную память для хранения разделенных частей массива.
	*/
	if len(nums) == 1 {
		return nums[0]
	}

	// Выбираем случайный опорный элемент
	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(len(nums))
	pivot := nums[pivotIndex]

	// Разделяем массив на элементы меньше, равные и больше опорного
	var lows, highs, pivots []int
	for _, num := range nums {
		if num < pivot {
			lows = append(lows, num)
		} else if num > pivot {
			highs = append(highs, num)
		} else {
			pivots = append(pivots, num)
		}
	}

	// Рекурсивно находим k-й наименьший элемент
	if k <= len(lows) {
		return quickselectV2(lows, k)
	} else if k <= len(lows)+len(pivots) {
		return pivots[0]
	} else {
		return quickselectV2(highs, k-len(lows)-len(pivots))
	}
}

// Функция для нахождения медианы массива
func median(nums []int) float64 {
	n := len(nums)
	if n%2 == 1 {
		// Если количество элементов нечетное, медиана - средний элемент
		return float64(quickselectV2(nums, n/2))
	} else {
		// Если количество элементов четное, медиана - среднее арифметическое двух средних элементов
		left := quickselectV2(nums, n/2-1)
		right := quickselectV2(nums, n/2)
		return float64(left+right) / 2.0
	}
}

// Функция для нахождения квантиля массива
func quantile(nums []int, q float64) float64 {
	k := int(q * float64(len(nums)))
	return float64(quickselectV2(nums, k))
}

func main() {
	/*
		Этот пример демонстрирует, как можно использовать Quickselect для нахождения медианы и квантилей в массиве с линейной временной сложностью в среднем случае.

		Quickselect - это алгоритм выбора, основанный на принципе "разделяй и властвуй".
		Он использует случайный опорный элемент для разделения массива на части и рекурсивно находит k-й наименьший элемент.
	*/
	nums := []int{3, 1, 8, 2, 7, 4, 6, 5, 9, 0}

	// Находим медиану
	med := median(nums)
	fmt.Printf("Median: %.2f\n", med)

	// Находим квантили
	q1 := quantile(nums, 0.25)
	q3 := quantile(nums, 0.75)
	fmt.Printf("Q1: %.2f\n", q1)
	fmt.Printf("Q3: %.2f\n", q3)
}
