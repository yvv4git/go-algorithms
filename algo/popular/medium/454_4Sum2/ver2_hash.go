//go:build ignore

package main

import "fmt"

// fourSumCount решает задачу 4Sum II используя подход с хеш-таблицей
// Временная сложность: O(n²)
// Пространственная сложность: O(n²)
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// Хеш-таблица для хранения сумм из первых двух массивов
	// Ключ - сумма, значение - количество вхождений этой суммы
	sumMap := make(map[int]int)

	// Первый этап: вычисляем все возможные суммы nums1[i] + nums2[j]
	// и сохраняем их в хеш-таблице с подсчетом количества
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			sumMap[sum]++
		}
	}

	count := 0

	// Второй этап: для каждой суммы nums3[k] + nums4[l]
	// ищем в хеш-таблице противоположное значение
	for k := 0; k < len(nums3); k++ {
		for l := 0; l < len(nums4); l++ {
			sum := nums3[k] + nums4[l]
			// Ищем значение, которое в сумме с текущей даст 0
			target := -sum

			// Если такое значение есть в хеш-таблице,
			// добавляем количество его вхождений к результату
			if occurrences, exists := sumMap[target]; exists {
				count += occurrences
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

	// Демонстрация работы алгоритма для первого примера
	fmt.Printf("\nДетальный разбор алгоритма для примера 1:\n")
	demonstrateHashAlgorithm(nums1, nums2, nums3, nums4)
}

// demonstrateHashAlgorithm показывает пошаговую работу алгоритма с хеш-таблицей
func demonstrateHashAlgorithm(nums1, nums2, nums3, nums4 []int) {
	fmt.Printf("Шаг 1: Создаем хеш-таблицу для сумм nums1[i] + nums2[j]\n")

	sumMap := make(map[int]int)

	// Показываем создание хеш-таблицы
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			sumMap[sum]++
			fmt.Printf("  nums1[%d] + nums2[%d] = %d + %d = %d\n",
				i, j, nums1[i], nums2[j], sum)
		}
	}

	fmt.Printf("\nХеш-таблица: %v\n\n", sumMap)

	fmt.Printf("Шаг 2: Ищем комплементарные суммы для nums3[k] + nums4[l]\n")

	count := 0
	for k := 0; k < len(nums3); k++ {
		for l := 0; l < len(nums4); l++ {
			sum := nums3[k] + nums4[l]
			target := -sum

			fmt.Printf("  nums3[%d] + nums4[%d] = %d + %d = %d, ищем %d в хеш-таблице\n",
				k, l, nums3[k], nums4[l], sum, target)

			if occurrences, exists := sumMap[target]; exists {
				count += occurrences
				fmt.Printf("    Найдено! Количество вхождений: %d\n", occurrences)
			} else {
				fmt.Printf("    Не найдено\n")
			}
		}
	}

	fmt.Printf("\nИтоговое количество комбинаций: %d\n", count)
}
