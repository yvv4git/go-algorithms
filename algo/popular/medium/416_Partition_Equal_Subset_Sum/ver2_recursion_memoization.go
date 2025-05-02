//go:build ignore

package main

import (
	"fmt"
)

// TASK: Дан массив положительных целых чисел nums. Определить, можно ли разделить массив на два подмножества
// с равной суммой элементов.
func canPartition(nums []int) bool {
	/*
	   APPROACH: Recursion with Memoization
	   1. Вычисляем сумму всех элементов массива. Если сумма нечетная, разделение невозможно.
	   2. Цель — найти подмножество с суммой, равной половине общей суммы (target = sum/2).
	   3. Используем рекурсивную функцию, которая проверяет, можно ли составить сумму target,
	      используя элементы массива с индекса i и далее.
	   4. Применяем мемоизацию, чтобы кэшировать результаты для пар (index, target).
	   5. Если удается составить target, возвращаем true.

	   TIME COMPLEXITY: O(n * target)
	   - n — длина массива nums.
	   - target — половина суммы элементов массива.
	   - Для каждой пары (index, target) выполняем вычисления один раз, а результаты кэшируются.
	   - Итоговая сложность O(n * target), так как всего возможно n * target состояний.

	   SPACE COMPLEXITY: O(n * target)
	   - Используем map для мемоизации, где хранятся результаты для пар (index, target).
	   - Размер кэша в худшем случае O(n * target).
	   - Дополнительно стек рекурсии занимает до O(n) памяти.
	   - Итоговая сложность O(n * target).
	*/

	// Вычисляем сумму всех элементов
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// Если сумма нечетная, разделение невозможно
	if sum%2 != 0 {
		return false
	}

	// Целевая сумма для одного подмножества
	target := sum / 2

	// Создаем map для мемоизации: ключ — строка "index:target"
	memo := make(map[string]bool)

	// Рекурсивная функция
	var canPartitionHelper func(index, target int) bool
	canPartitionHelper = func(index, target int) bool {
		// Базовые случаи
		if target == 0 {
			return true
		}
		if index >= len(nums) || target < 0 {
			return false
		}

		// Ключ для мемоизации
		key := fmt.Sprintf("%d:%d", index, target)

		// Проверяем, есть ли результат в кэше
		if val, exists := memo[key]; exists {
			return val
		}

		// Варианты: взять текущее число или пропустить его
		take := canPartitionHelper(index+1, target-nums[index])
		skip := canPartitionHelper(index+1, target)

		// Сохраняем результат в кэш
		memo[key] = take || skip
		return memo[key]
	}

	// Запускаем рекурсию с начальным индексом 0 и целевой суммой
	return canPartitionHelper(0, target)
}

func main() {
	// Тестовые случаи
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 5, 11, 5}, true}, // Можно разделить на [1,5,5] и [11], сумма = 11
		{[]int{1, 2, 3, 5}, false}, // Нельзя разделить на подмножества с равной суммой
		{[]int{2, 2, 3, 5}, false}, // Нельзя разделить на подмножества с равной суммой
		{[]int{1, 1}, true},        // Можно разделить на [1] и [1], сумма = 1
		{[]int{100}, false},        // Нельзя разделить одно число
	}

	// Запускаем тесты
	for i, tc := range testCases {
		result := canPartition(tc.nums) // Исправлено: tc.nums вместо nums
		fmt.Printf("Тест %d: nums=%v, Ожидаемый результат=%v, Полученный результат=%v\n",
			i+1, tc.nums, tc.expected, result)
		if result == tc.expected {
			fmt.Println("Тест пройден успешно!")
		} else {
			fmt.Println("Тест не пройден!")
		}
	}
}
