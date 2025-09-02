//go:build ignore

package main

import "fmt"

// Находит максимальную длину подмассива с равным количеством 0 и 1 (Brute Force)
func findMaxLength(nums []int) int {
	/*
		APPROACH: BRUTE FORCE (CHECK ALL SUBARRAYS)
		Алгоритм проверяет все возможные подмассивы для нахождения максимальной длины
		с равным количеством 0 и 1.

		1. **Инициализация**:
		   - Устанавливаем начальное значение maxLen = 0

		2. **Перебор всех подмассивов**:
		   - Внешний цикл: начало подмассива (i от 0 до n-1)
		   - Внутренний цикл: конец подмассива (j от i до n-1)
		   - Для каждого подмассива nums[i..j]:
		     - Считаем количество 0 (zeros) и 1 (ones)
		     - Если zeros == ones, обновляем maxLen

		3. **Возврат результата**:
		   - После проверки всех подмассивов возвращаем maxLen

		Сложность:
		- Время: O(n²) - два вложенных цикла
		- Память: O(1) - постоянное дополнительное пространство

		Примечания:
		1. Метод прост в реализации, но неэффективен для больших массивов
		2. Подходит только для небольших входных данных
		3. Для каждого подмассива заново подсчитываем 0 и 1
	*/
	maxLen := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		zeros, ones := 0, 0
		for j := i; j < n; j++ {
			if nums[j] == 0 {
				zeros++
			} else {
				ones++
			}

			if zeros == ones && (j-i+1) > maxLen {
				maxLen = j - i + 1
			}
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

	// Тест 7: Большой массив (неэффективно для brute force)
	nums7 := []int{0, 1, 1, 0, 1, 1, 1, 0, 0, 0}
	fmt.Printf("Input: %v\nResult: %d (Expected: 8)\n\n", nums7, findMaxLength(nums7))
}
