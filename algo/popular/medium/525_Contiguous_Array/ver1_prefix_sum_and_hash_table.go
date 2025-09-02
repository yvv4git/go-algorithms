//go:build ignore

package main

import "fmt"

// Находит максимальную длину подмассива с равным количеством 0 и 1
func findMaxLength(nums []int) int {
	/*
		APPROACH: PREFIX SUM WITH HASH TABLE (OPTIMAL)
		Алгоритм использует хеш-таблицу для хранения первых вхождений префиксных сумм,
		чтобы эффективно находить подмассивы с равным количеством 0 и 1.

		1. **Инициализация**:
		   - Создаем map для хранения первых вхождений сумм (ключ - сумма, значение - индекс)
		   - Устанавливаем начальное значение sum=0 с индексом -1

		2. **Преобразование элементов**:
		   - 0 заменяем на -1, 1 оставляем как есть
		   - Это позволяет считать сумму подмассива: 0 означает равное количество 0 и 1

		3. **Вычисление префиксных сумм**:
		   - Идем по массиву, накапливая сумму
		   - Если сумма уже встречалась, вычисляем длину подмассива
		   - Если сумма новая, сохраняем ее в map

		Сложность:
		- Время: O(n) - один проход по массиву
		- Память: O(n) - хранение сумм в map

		Пример:
		nums = [0,1,0,1,1,0,0]
		sums = [-1,0,-1,0,1,0,-1]
		Максимальная длина подмассива: 6 ([0,1,0,1,1,0] или [1,0,1,1,0,0])
	*/
	if len(nums) == 0 {
		return 0
	}

	sumMap := make(map[int]int)
	sumMap[0] = -1 // Базовый случай
	maxLen, sum := 0, 0

	for i, num := range nums {
		// Преобразуем 0 в -1, 1 оставляем
		if num == 0 {
			sum--
		} else {
			sum++
		}

		// Если сумма уже встречалась, обновляем maxLen
		if prevIndex, exists := sumMap[sum]; exists {
			if i-prevIndex > maxLen {
				maxLen = i - prevIndex
			}
		} else {
			sumMap[sum] = i
		}
	}

	return maxLen
}

func main() {
	// Тест 1: Простой случай
	nums1 := []int{0, 1}
	fmt.Printf("Input: %v\nResult: %d (Expected: 2)\n\n", nums1, findMaxLength(nums1))

	// Тест 2: Чередующиеся 0 и 1
	nums2 := []int{0, 1, 0}
	fmt.Printf("Input: %v\nResult: %d (Expected: 2)\n\n", nums2, findMaxLength(nums2))

	// Тест 3: Длинный подмассив
	nums3 := []int{0, 1, 0, 1, 0, 1, 0, 0, 1}
	fmt.Printf("Input: %v\nResult: %d (Expected: 6)\n\n", nums3, findMaxLength(nums3))

	// Тест 4: Нет подходящего подмассива
	nums4 := []int{0, 0, 0}
	fmt.Printf("Input: %v\nResult: %d (Expected: 0)\n\n", nums4, findMaxLength(nums4))

	// Тест 5: Пустой массив
	nums5 := []int{}
	fmt.Printf("Input: %v\nResult: %d (Expected: 0)\n\n", nums5, findMaxLength(nums5))

	// Тест 6: Все 1
	nums6 := []int{1, 1, 1}
	fmt.Printf("Input: %v\nResult: %d (Expected: 0)\n\n", nums6, findMaxLength(nums6))

	// Тест 7: Большой массив
	nums7 := []int{0, 1, 1, 0, 1, 1, 1, 0, 0, 0}
	fmt.Printf("Input: %v\nResult: %d (Expected: 8)\n\n", nums7, findMaxLength(nums7))
}
