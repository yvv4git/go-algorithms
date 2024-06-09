package _67_two_sum_2_Input_array_Is_sorted

func twoSumV2(numbers []int, target int) []int {
	/*
		METHOD: Binary search
		TIME COMPLEXITY: O(n log n), где n - количество элементов в массиве.
		SPACE COMPLEXITY: O(1), поскольку мы не используем дополнительное пространство, которое зависит от размера входных данных.
	*/
	for i := 0; i < len(numbers); i++ {
		j := binarySearch(numbers, i+1, len(numbers)-1, target-numbers[i])
		if j != -1 {
			return []int{i + 1, j + 1}
		}
	}

	return []int{-1, -1}
}

// Функция binarySearch выполняет бинарный поиск в массиве numbers.
// Она начинает поиск с индекса start и заканчивает его на индексе end.
// Если число target найдено в массиве, функция возвращает его индекс.
// Если число target не найдено, функция возвращает -1.
func binarySearch(numbers []int, start, end, target int) int {
	// Цикл продолжается, пока start не больше end.
	for start <= end {
		// Вычисляем средний индекс.
		mid := start + (end-start)/2

		// Если число по среднему индексу равно target,
		// то мы нашли target и возвращаем его индекс.
		if numbers[mid] == target {
			return mid
		} else if numbers[mid] < target {
			// Если число по среднему индексу меньше target,
			// то мы знаем, что target не может быть в левой половине массива,
			// поэтому мы начинаем поиск с индекса mid + 1.
			start = mid + 1
		} else {
			// Если число по среднему индексу больше target,
			// то мы знаем, что target не может быть в правой половине массива,
			// поэтому мы заканчиваем поиск с индекса mid - 1.
			end = mid - 1
		}
	}

	// Если мы дошли до этой точки, то target не найден в массиве,
	// поэтому мы возвращаем -1.
	return -1
}
