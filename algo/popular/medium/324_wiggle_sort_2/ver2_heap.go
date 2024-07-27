package main

import "sort"

func wiggleSortV2(nums []int) {
	/*
		METHOD: Используется метод "Heap Sort" с модификацией для нахождения k-го наибольшего элемента.
		После нахождения k-го наибольшего элемента, массив разделяется на две части:
		элементы больше или равные k-му наибольшему и остальные. Затем элементы чередуются.

		TIME COMPLEXITY: O(n log n) в худшем случае, где n - количество элементов в массиве.
		SPACE COMPLEXITY: O(n) для хранения кучи.
	*/
	var heap []int

	// Построение максимальной кучи
	for _, num := range nums {
		heap = append(heap, num)
		heapUp(heap, len(heap)-1) // Всплытие нового элемента в куче
	}

	k := len(nums) / 2

	var topKLargestElement []int

	// Извлечение k наибольших элементов
	for i := 0; i < k; i++ {
		topKLargestElement = append(topKLargestElement, heap[0])
		heap[0] = heap[len(heap)-1]
		heap = heap[:len(heap)-1]
		heapDown(heap, 0, len(heap)-1) // Погружение корневого элемента в куче
	}

	// Сортировка оставшихся элементов в порядке убывания
	sort.Slice(heap, func(i, j int) bool {
		return heap[i] > heap[j]
	})

	j, k := 0, 0

	// Чередование элементов
	for i := 0; i < len(nums); i++ {
		if i%2 == 1 {
			nums[i] = topKLargestElement[j]
			j++
		} else {
			nums[i] = heap[k]
			k++
		}
	}
}

// Всплытие элемента в куче
func heapUp(heap []int, pos int) {
	// Вычисляем индекс родительского элемента
	parent := (pos - 1) / 2

	// Проверяем, что родительский элемент существует и меньше текущего элемента
	if parent >= 0 && heap[parent] < heap[pos] {
		// Меняем местами родительский и текущий элементы
		heap[parent], heap[pos] = heap[pos], heap[parent]
		// Рекурсивно вызываем heapUp для родительского элемента
		heapUp(heap, parent)
	}
}

// Погружение элемента в куче
func heapDown(heap []int, pos, limit int) {
	// Вычисляем индексы левого и правого потомков
	l, r := 2*pos+1, 2*pos+2
	// Инициализируем переменную для хранения индекса наибольшего элемента
	larger := pos

	// Проверяем, что левый потомок находится в пределах кучи и больше текущего элемента
	if l <= limit && heap[l] > heap[larger] {
		larger = l
	}

	// Проверяем, что правый потомок находится в пределах кучи и больше текущего элемента или левого потомка
	if r <= limit && heap[r] > heap[larger] {
		larger = r
	}

	// Если наибольший элемент не является текущим, меняем их местами и продолжаем погружение
	if larger != pos {
		heap[larger], heap[pos] = heap[pos], heap[larger]
		heapDown(heap, larger, limit)
	}
}
