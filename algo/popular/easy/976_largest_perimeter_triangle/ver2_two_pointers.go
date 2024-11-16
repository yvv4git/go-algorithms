package main

import (
	"sort"
)

func largestPerimeterV2(nums []int) int {
	/*
		METHOD: Two Pointers.
		Данный метод основан на сортировке массива по возрастанию и использовании двух указателей
		для поиска трех последовательных элементов, которые могут образовать треугольник.

		TIME COMPLEXITY: O(n log n)
		Описание: Временная сложность определяется сортировкой массива, которая занимает O(n log n) времени, где n — количество элементов в массиве.
		Проход по массиву для проверки каждой тройки соседних элементов занимает O(n) времени.
		Общая сложность алгоритма составляет O(n log n), так как сортировка доминирует над проходом по массиву.

		SPACE COMPLEXITY: O(1)
		Описание: Пространственная сложность алгоритма составляет O(1), что означает постоянное количество дополнительной памяти.
		Мы не создаем новых массивов или структур данных, которые бы росли с увеличением размера входного массива.
	*/
	// Сортируем массив по возрастанию
	// Временная сложность сортировки: O(n log n)
	sort.Ints(nums)

	// Используем два указателя для поиска трех последовательных элементов, которые могут образовать треугольник
	// Временная сложность прохода по массиву: O(n)
	for i := len(nums) - 1; i >= 2; i-- {
		// Проверяем условие неравенства треугольника
		if nums[i] < nums[i-1]+nums[i-2] {
			// Если условие выполняется, возвращаем периметр
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}

	// Если не найдено ни одной тройки, удовлетворяющей условию, возвращаем 0
	return 0
}

// func main() {
// 	// Примеры использования функции
// 	nums1 := []int{2, 1, 2}
// 	fmt.Println(largestPerimeter(nums1)) // Вывод: 5

// 	nums2 := []int{1, 2, 1}
// 	fmt.Println(largestPerimeter(nums2)) // Вывод: 0
// }
