package main

import (
	"fmt"
	"sort"
)

// fourSumCount решает задачу 4Sum II используя подход с сортировкой
// Временная сложность: O(n^2 log n) - сортировка и двойной проход по массивам
// Пространственная сложность: O(n^2) - два дополнительных массива для хранения сумм
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// Генерируем все возможные суммы для первых двух массивов
	sums12 := make([]int, 0, len(nums1)*len(nums2))
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sums12 = append(sums12, nums1[i]+nums2[j])
		}
	}

	// Генерируем все возможные суммы для последних двух массивов
	sums34 := make([]int, 0, len(nums3)*len(nums4))
	for k := 0; k < len(nums3); k++ {
		for l := 0; l < len(nums4); l++ {
			sums34 = append(sums34, nums3[k]+nums4[l])
		}
	}

	// Сортируем массивы сумм
	sort.Ints(sums12)
	sort.Ints(sums34)

	// Используем технику двух указателей для поиска пар с суммой 0
	count := 0
	left := 0
	right := len(sums34) - 1

	for left < len(sums12) && right >= 0 {
		sum := sums12[left] + sums34[right]

		if sum == 0 {
			// Найдена пара с суммой 0
			// Подсчитываем количество одинаковых элементов слева
			leftCount := 1
			for left+leftCount < len(sums12) && sums12[left+leftCount] == sums12[left] {
				leftCount++
			}

			// Подсчитываем количество одинаковых элементов справа
			rightCount := 1
			for right-rightCount >= 0 && sums34[right-rightCount] == sums34[right] {
				rightCount++
			}

			// Добавляем произведение количеств к результату
			count += leftCount * rightCount

			// Перемещаем указатели
			left += leftCount
			right -= rightCount
		} else if sum < 0 {
			// Сумма слишком мала, увеличиваем левый указатель
			left++
		} else {
			// Сумма слишком велика, уменьшаем правый указатель
			right--
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

	// Демонстрация работы алгоритма для первого примера
	fmt.Printf("\nДетальный разбор алгоритма с сортировкой для примера 1:\n")
	demonstrateSortAlgorithm(nums1, nums2, nums3, nums4)
}

// demonstrateSortAlgorithm показывает пошаговую работу алгоритма с сортировкой
func demonstrateSortAlgorithm(nums1, nums2, nums3, nums4 []int) {
	fmt.Printf("Шаг 1: Генерируем суммы для первых двух массивов\n")

	sums12 := make([]int, 0, len(nums1)*len(nums2))
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			sums12 = append(sums12, sum)
			fmt.Printf("  nums1[%d] + nums2[%d] = %d + %d = %d\n",
				i, j, nums1[i], nums2[j], sum)
		}
	}
	fmt.Printf("Суммы nums1+nums2 (до сортировки): %v\n", sums12)

	fmt.Printf("\nШаг 2: Генерируем суммы для последних двух массивов\n")

	sums34 := make([]int, 0, len(nums3)*len(nums4))
	for k := 0; k < len(nums3); k++ {
		for l := 0; l < len(nums4); l++ {
			sum := nums3[k] + nums4[l]
			sums34 = append(sums34, sum)
			fmt.Printf("  nums3[%d] + nums4[%d] = %d + %d = %d\n",
				k, l, nums3[k], nums4[l], sum)
		}
	}
	fmt.Printf("Суммы nums3+nums4 (до сортировки): %v\n", sums34)

	// Сортируем массивы
	sort.Ints(sums12)
	sort.Ints(sums34)

	fmt.Printf("\nШаг 3: Сортируем массивы сумм\n")
	fmt.Printf("Суммы nums1+nums2 (после сортировки): %v\n", sums12)
	fmt.Printf("Суммы nums3+nums4 (после сортировки): %v\n", sums34)

	fmt.Printf("\nШаг 4: Используем технику двух указателей\n")

	count := 0
	left := 0
	right := len(sums34) - 1
	step := 1

	for left < len(sums12) && right >= 0 {
		sum := sums12[left] + sums34[right]

		fmt.Printf("  Шаг %d: left=%d, right=%d, sums12[%d]=%d, sums34[%d]=%d, сумма=%d\n",
			step, left, right, left, sums12[left], right, sums34[right], sum)

		if sum == 0 {
			// Подсчитываем одинаковые элементы
			leftCount := 1
			for left+leftCount < len(sums12) && sums12[left+leftCount] == sums12[left] {
				leftCount++
			}

			rightCount := 1
			for right-rightCount >= 0 && sums34[right-rightCount] == sums34[right] {
				rightCount++
			}

			fmt.Printf("    Найдена пара! leftCount=%d, rightCount=%d, добавляем %d\n",
				leftCount, rightCount, leftCount*rightCount)

			count += leftCount * rightCount
			left += leftCount
			right -= rightCount
		} else if sum < 0 {
			fmt.Printf("    Сумма < 0, увеличиваем left\n")
			left++
		} else {
			fmt.Printf("    Сумма > 0, уменьшаем right\n")
			right--
		}
		step++
	}

	fmt.Printf("\nИтоговое количество комбинаций: %d\n", count)
}
