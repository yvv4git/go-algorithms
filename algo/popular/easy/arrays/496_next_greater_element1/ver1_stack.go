package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	/*
		METHOD: Stack

		TIME COMPLEXITY: O(n), где n - количество элементов в массиве nums2,
		так как мы проходим по массиву дважды: один раз для создания словаря nextGreater,
		а второй раз для создания результирующего массива.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем заполнить словарь и стек n элементами.
	*/
	// Создаем словарь для хранения следующего большего элемента для каждого элемента из nums2
	nextGreater := make(map[int]int)
	// Создаем стек для отслеживания элементов
	stack := []int{}

	// Проходим по nums2 справа налево
	for i := len(nums2) - 1; i >= 0; i-- {
		// Пока стек не пуст и элемент на вершине стека меньше текущего, удаляем его из стека
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}

		// Если стек не пуст, то элемент на вершине стека - следующий больший элемент
		if len(stack) > 0 {
			nextGreater[nums2[i]] = stack[len(stack)-1]
		} else {
			// Иначе следующий больший элемент отсутствует
			nextGreater[nums2[i]] = -1
		}

		// Добавляем текущий элемент в стек
		stack = append(stack, nums2[i])
	}

	// Создаем результирующий массив для ответов
	result := make([]int, len(nums1))
	// Заполняем результирующий массив значениями из словаря nextGreater
	for i, num := range nums1 {
		result[i] = nextGreater[num]
	}

	return result
}
