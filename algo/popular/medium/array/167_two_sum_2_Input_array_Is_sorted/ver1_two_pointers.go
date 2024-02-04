package _67_two_sum_2_Input_array_Is_sorted

// Функция для поиска двух чисел в отсортированном массиве, сумма которых равна target
func twoSumV1(numbers []int, target int) []int {
	/*
		METHOD: Two pointers
		Time complexity: O(n), где n - количество элементов в массиве. Это связано с тем, что мы проходим по массиву только один раз.
		Space complexity: O(1), поскольку мы не используем дополнительное пространство, которое зависит от размера входных данных.
		Мы используем только несколько переменных для хранения индексов, но это не зависит от размера входных данных.
	*/

	// Инициализируем два указателя - один на начало массива, другой на конец
	left, right := 0, len(numbers)-1

	// Пока указатели не встретятся
	for left < right {
		// Если сумма чисел равна target, возвращаем индексы
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}

		// Если сумма чисел больше target, двигаем правый указатель
		if numbers[left]+numbers[right] > target {
			right--
		} else {
			// Если сумма чисел меньше target, двигаем левый указатель
			left++
		}
	}

	// Если не нашли пары чисел, возвращаем пустой массив
	return []int{}
}
