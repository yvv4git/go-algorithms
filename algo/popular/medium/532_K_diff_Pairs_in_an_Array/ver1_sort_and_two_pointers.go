package main

import (
	"fmt"
	"sort"
)

// findPairs находит количество уникальных k-diff пар в массиве
// Использует алгоритм сортировки + два указателя
func findPairs(nums []int, k int) int {
	/*
		МЕТОД: Сортировка + два указателя
		ОПИСАНИЕ:
		- Сортируем входной массив для упорядочивания элементов
		- Используем два указателя (left и right) для поиска пар с разницей k
		- Левый указатель фиксирует меньший элемент пары, правый - больший
		- Пропускаем дубликаты для избежания повторного подсчета одинаковых пар
		- Когда находим пару с нужной разницей, увеличиваем счетчик и сдвигаем оба указателя
		- Если разница меньше k, сдвигаем правый указатель вправо
		- Если разница больше k, сдвигаем левый указатель вправо

		TIME COMPLEXITY: O(n log n), где n - количество элементов в массиве
		Основное время тратится на сортировку массива (O(n log n)), проход двумя указателями занимает O(n)

		SPACE COMPLEXITY: O(1), если не учитывать пространство для сортировки
		Алгоритм использует только константное дополнительное пространство для указателей и счетчика
		(не считая места для входного массива)
	*/
	if len(nums) < 2 {
		return 0
	}

	// Сортируем массив
	sort.Ints(nums)

	count := 0
	left, right := 0, 1

	for left < len(nums) && right < len(nums) {
		// Пропускаем дубликаты слева
		for left > 0 && left < len(nums) && nums[left] == nums[left-1] {
			left++
		}

		// Если указатели разошлись слишком далеко, сдвигаем левый
		if left >= right {
			right = left + 1
			continue
		}

		diff := nums[right] - nums[left]

		if diff == k {
			count++
			// Пропускаем все дубликаты справа
			for right < len(nums)-1 && nums[right] == nums[right+1] {
				right++
			}
			left++
			right++
		} else if diff < k {
			right++
		} else {
			left++
		}
	}

	return count
}

func main() {
	// Тестовые примеры из задачи
	fmt.Println("Example 1:", findPairs([]int{3, 1, 4, 1, 5}, 2)) // Output: 2
	fmt.Println("Example 2:", findPairs([]int{1, 2, 3, 4, 5}, 1)) // Output: 4
	fmt.Println("Example 3:", findPairs([]int{1, 3, 1, 5, 4}, 0)) // Output: 1

	// Дополнительные тесты
	fmt.Println("Test 4:", findPairs([]int{1, 1, 1, 2, 2}, 1)) // Output: 1
	fmt.Println("Test 5:", findPairs([]int{}, 1))              // Output: 0
	fmt.Println("Test 6:", findPairs([]int{1}, 1))             // Output: 0
}
