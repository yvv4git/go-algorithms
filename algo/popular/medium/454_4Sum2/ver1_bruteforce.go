//go:build ignore

package main

import "fmt"

// fourSumCount решает задачу 4Sum II используя подход Brute Force
// Временная сложность: O(n^4)
// Пространственная сложность: O(1)
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	n := len(nums1)

	// Четыре вложенных цикла для перебора всех возможных комбинаций индексов
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					// Проверяем, равна ли сумма элементов нулю
					if nums1[i]+nums2[j]+nums3[k]+nums4[l] == 0 {
						count++
					}
				}
			}
		}
	}

	return count
}

func main() {
	// Пример 1
	nums1 := []int{1, 2}
	nums2 := []int{-2, -1}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}

	result1 := fourSumCount(nums1, nums2, nums3, nums4)
	fmt.Printf("Пример 1:\n")
	fmt.Printf("nums1 = %v\n", nums1)
	fmt.Printf("nums2 = %v\n", nums2)
	fmt.Printf("nums3 = %v\n", nums3)
	fmt.Printf("nums4 = %v\n", nums4)
	fmt.Printf("Результат: %d\n", result1)
	fmt.Printf("Ожидаемый результат: 2\n\n")

	// Пример 2
	nums1_2 := []int{0}
	nums2_2 := []int{0}
	nums3_2 := []int{0}
	nums4_2 := []int{0}

	result2 := fourSumCount(nums1_2, nums2_2, nums3_2, nums4_2)
	fmt.Printf("Пример 2:\n")
	fmt.Printf("nums1 = %v\n", nums1_2)
	fmt.Printf("nums2 = %v\n", nums2_2)
	fmt.Printf("nums3 = %v\n", nums3_2)
	fmt.Printf("nums4 = %v\n", nums4_2)
	fmt.Printf("Результат: %d\n", result2)
	fmt.Printf("Ожидаемый результат: 1\n\n")

	// Дополнительный пример для демонстрации
	nums1_3 := []int{-1, -1}
	nums2_3 := []int{-1, 1}
	nums3_3 := []int{-1, 1}
	nums4_3 := []int{1, -1}

	result3 := fourSumCount(nums1_3, nums2_3, nums3_3, nums4_3)
	fmt.Printf("Дополнительный пример:\n")
	fmt.Printf("nums1 = %v\n", nums1_3)
	fmt.Printf("nums2 = %v\n", nums2_3)
	fmt.Printf("nums3 = %v\n", nums3_3)
	fmt.Printf("nums4 = %v\n", nums4_3)
	fmt.Printf("Результат: %d\n", result3)

	// Демонстрация найденных комбинаций для первого примера
	fmt.Printf("\nДетальный разбор для примера 1:\n")
	demonstrateCombinations(nums1, nums2, nums3, nums4)
}

// demonstrateCombinations показывает все найденные комбинации для лучшего понимания
func demonstrateCombinations(nums1, nums2, nums3, nums4 []int) {
	n := len(nums1)
	count := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				for l := 0; l < n; l++ {
					sum := nums1[i] + nums2[j] + nums3[k] + nums4[l]
					if sum == 0 {
						count++
						fmt.Printf("Комбинация %d: (%d, %d, %d, %d) -> %d + %d + %d + %d = %d\n",
							count, i, j, k, l, nums1[i], nums2[j], nums3[k], nums4[l], sum)
					}
				}
			}
		}
	}
}
