//go:build ignore

package main

import "fmt"

// Функция для нахождения медианы двух отсортированных массивов
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	/*
		Задача: Найти медиану двух отсортированных массивов.

		Подход:
		1. Используем рекурсивный метод для сравнения медиан двух массивов.
		2. Отбрасываем половину каждого массива, пока не найдем медиану.

		ТАЙМ-КОМПЛЕКСНОСТЬ: O(log(min(n, m)))
		СПЕЙС-КОМПЛЕКСНОСТЬ: O(1)
	*/
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high := 0, n1
	for low <= high {
		partition1 := (low + high) / 2
		partition2 := (n1+n2+1)/2 - partition1

		maxLeft1 := -1 << 63
		if partition1 > 0 {
			maxLeft1 = nums1[partition1-1]
		}

		minRight1 := 1<<63 - 1
		if partition1 < n1 {
			minRight1 = nums1[partition1]
		}

		maxLeft2 := -1 << 63
		if partition2 > 0 {
			maxLeft2 = nums2[partition2-1]
		}

		minRight2 := 1<<63 - 1
		if partition2 < n2 {
			minRight2 = nums2[partition2]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (n1+n2)%2 == 0 {
				return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2.0
			} else {
				return float64(max(maxLeft1, maxLeft2))
			}
		} else if maxLeft1 > minRight2 {
			high = partition1 - 1
		} else {
			low = partition1 + 1
		}
	}

	return -1
}

// Функция для нахождения максимального из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Функция для нахождения минимального из двух чисел
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Пример: два отсортированных массива
	nums1 := []int{1, 3, 8}
	nums2 := []int{7, 9, 10, 11}

	// Вызываем функцию для нахождения медианы
	result := findMedianSortedArrays(nums1, nums2)

	// Выводим результат
	fmt.Printf("Медиана двух отсортированных массивов: %.2f\n", result)
}
