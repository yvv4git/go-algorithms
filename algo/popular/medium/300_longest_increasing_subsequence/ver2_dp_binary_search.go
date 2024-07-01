package main

func lengthOfLISV2(nums []int) int {
	/*
		METHOD: Dynamic programming & Binary search

		TIME COMPLEXITY: O(n log n), т.к. во-первых мы проходим по всем элементам, что дает O(n), а во вторых,
		нам надо использовать бинарный поиск, что дает O(log n). Соответственно, итоговая временная сложность будет
		равна O(n log n).

		SPACE COMPLEXITY: O(n), это объясняется тем, что в худшем случае массив stack может содержать все элементы из входного массива nums.
	*/
	stack := []int{}

	for _, num := range nums {
		if len(stack) == 0 || stack[len(stack)-1] < num {
			// Если стек пуст или последний элемент стека меньше текущего числа,
			// добавляем текущее число в стек.
			stack = append(stack, num)
		} else {
			// Иначе, находим позицию для замены элемента в стеке с помощью бинарного поиска.
			idx := findNextMaxNum(stack, num)
			stack[idx] = num
		}
	}

	// Длина стека в конце будет длиной самой длинной возрастающей подпоследовательности.
	return len(stack)
}

func findNextMaxNum(arr []int, num int) int {
	left, right := 0, len(arr)-1

	// Бинарный поиск для нахождения позиции, где текущее число должно быть вставлено.
	for left < right {
		mid := (left + right) / 2
		if arr[mid] >= num {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
