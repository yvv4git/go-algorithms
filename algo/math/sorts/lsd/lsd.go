package main

import (
	"fmt"
	"math"
)

// Функция для определения количества разрядов в наибольшем числе
func maxDigits(nums []int) int {
	maxDigits := 0
	for _, num := range nums {
		digits := int(math.Log10(float64(num))) + 1
		if digits > maxDigits {
			maxDigits = digits
		}
	}
	return maxDigits
}

// LSD сортировка
func lsdRadixSort(nums []int) []int {
	/*
		METHOD: LSD - Least Significant Digit

		TIME COMPLEXITY: O(d * (n + b)), где d - количество разрядов, n - количество чисел, b - основание системы счисления (в нашем случае 10).
		Это с учетом того, что мы проходим по каждому разряду и выполняем операции подсчета и перемещения чисел.

		SPACE COMPLEXITY: O(n + b), где n - количество чисел, b - основание системы счисления.
		Это с учетом дополнительного пространства, необходимого для хранения временного сортирующего массива и счетчиков.
	*/
	// Определяем количество разрядов в наибольшем числе
	digits := maxDigits(nums)

	// Создаем временный срез для хранения отсортированных чисел
	sorted := make([]int, len(nums))

	// Проходим по каждому разряду
	for place := 1; place <= 10^digits; place *= 10 {
		// Создаем счетчики для каждого возможного значения разряда (от 0 до 9)
		counts := make([]int, 10)

		// Подсчитываем количество чисел с определенным значением разряда
		for _, num := range nums {
			digit := (num / place) % 10
			counts[digit]++
		}

		// Изменяем счетчики, чтобы значения в них представляли индексы в выходном срезе
		for i := 1; i < 10; i++ {
			counts[i] += counts[i-1]
		}

		// Перемещаем числа в выходной срез в соответствии с индексами, полученными в счетчиках
		for i := len(nums) - 1; i >= 0; i-- {
			num := nums[i]
			digit := (num / place) % 10
			sorted[counts[digit]-1] = num
			counts[digit]--
		}

		// Обновляем входной срез для следующей итерации
		copy(nums, sorted)
	}

	return nums
}

func main() {
	nums := []int{329, 457, 657, 839, 436, 720, 355}
	sortedNums := lsdRadixSort(nums)
	fmt.Println(sortedNums)
}
