package main

// Функция для поиска опорного индекса
func pivotIndexV2(nums []int) int {
	/*
		METHOD: Prefix sum
		В этом алгоритме мы используем один проход для вычисления префиксных сумм,
		что позволяет нам получить сумму всех элементов до каждого индекса за один проход.
		Затем мы используем эти префиксные суммы для поиска опорного индекса.
		Если сумма элементов слева от текущего индекса равна сумме элементов справа, то этот индекс является опорным.
		Если мы не нашли опорного индекса, возвращаем -1.

		Префиксная сумма (prefix sum) - это математический термин,
		который используется в информатике и математике для быстрого вычисления суммы подмассива любого массива.
		Префиксная сумма для элемента i в массиве A обычно обозначается как S[i]
		и определяется как сумма всех элементов от начала массива до индекса i включительно.
		То есть, S[i] = A[0] + A[1] + ... + A[i].

		Префиксные суммы позволяют быстро вычислять сумму подмассива за константное время,
		независимо от размера подмассива, используя простой арифметический выражение.

		TIME COMPLEXITY: O(n), где n - количество элементов в массиве, потому что мы проходим по массиву два раза.

		SPACE COMPLEXITY: O(1), так как мы используем только несколько переменных для хранения сумм и индекса.
	*/
	// Вычисляем префиксные суммы
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	// Вычисляем общую сумму всех элементов
	total := nums[len(nums)-1]

	// Инициализируем левую сумму
	var left int

	// Проходим по массиву
	for i := range nums {
		// Если индекс больше 0, то обновляем левую сумму
		if i > 0 {
			left = nums[i-1]
		} else {
			// Если индекс равен 0, то левую сумму считаем равной 0
			left = 0
		}

		// Если левую сумму равную общая сумме минус текущий элемент, то это опорный индекс
		if left == total-nums[i] {
			return i
		}
	}

	// Если опорного индекса нет, возвращаем -1
	return -1
}