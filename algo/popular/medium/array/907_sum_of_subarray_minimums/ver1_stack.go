package _07_sum_of_subarray_minimums

func sumSubarrayMinsV1(arr []int) int {
	/*
		METHOD: Two pointers & stack
		TIME COMPLEXITY: O(n), где n - размер входного массива, потому что мы проходим по массиву два раза.
		SPACE COMPLEXITY: O(n), потому что мы создаем стек для хранения индексов элементов. Стек может содержать до n элементов в худшем случае.
	*/

	// Инициализируем стек и переменные
	stack, res, mod := []int{}, 0, 1000000007

	// Проходим по массиву
	for j := 0; j < len(arr); j++ {
		// Пока стек не пуст и текущий элемент меньше элемента на вершине стека,
		// то удаляем элементы из стека
		for len(stack) > 0 && arr[j] < arr[stack[len(stack)-1]] {
			// Вычисляем количество подмассивов, которые заканчиваются на текущий элемент
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			right := j
			count := (mid - left) * (right - mid) % mod
			res += (arr[mid] * count) % mod
			res %= mod
		}

		// Добавляем текущий элемент в стек
		stack = append(stack, j)
	}

	// Обрабатываем оставшиеся элементы в стеке
	right := len(arr)
	for len(stack) > 0 {
		mid := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left := -1
		if len(stack) > 0 {
			left = stack[len(stack)-1]
		}
		count := (mid - left) * (right - mid) % mod
		res += (arr[mid] * count) % mod
		res %= mod
	}

	return res
}
