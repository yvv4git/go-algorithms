package main

import "fmt"

// Находит максимальную длину подмассива с равным количеством 0 и 1 (Prefix Sum без преобразования)
func findMaxLength(nums []int) int {
	/*
		APPROACH: PREFIX SUM WITHOUT 0 TO -1 CONVERSION
		Алгоритм использует хеш-таблицу для хранения первых вхождений разницы между количеством 1 и 0,
		чтобы находить подмассивы с равным количеством 0 и 1.

		1. **Инициализация**:
		   - Создаем map для хранения первых вхождений разницы (diffMap)
		   - Устанавливаем начальное значение diffMap[0] = -1
		   - Инициализируем счетчики count0 и count1

		2. **Итерация по массиву**:
		   - Для каждого элемента обновляем count0 или count1
		   - Вычисляем разницу diff = count1 - count0
		   - Если разница уже встречалась, вычисляем длину подмассива
		   - Если разница новая, сохраняем ее в map

		3. **Возврат результата**:
		   - После обработки всех элементов возвращаем maxLen

		Сложность:
		- Время: O(n) - один проход по массиву
		- Память: O(n) - хранение разниц в map

		Примечания:
		1. Не требует преобразования 0 в -1
		2. Использует разницу между количеством 1 и 0
		3. Подмассив с равным количеством 0 и 1 имеет diff = 0
	*/
	if len(nums) == 0 {
		return 0
	}

	diffMap := make(map[int]int)
	diffMap[0] = -1 // Базовый случай
	maxLen, count0, count1 := 0, 0, 0

	for i, num := range nums {
		if num == 0 {
			count0++
		} else {
			count1++
		}

		diff := count1 - count0

		if prevIndex, exists := diffMap[diff]; exists {
			if i-prevIndex > maxLen {
				maxLen = i - prevIndex
			}
		} else {
			diffMap[diff] = i
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
