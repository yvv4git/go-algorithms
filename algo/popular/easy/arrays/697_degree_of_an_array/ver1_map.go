package main

func findShortestSubArray(nums []int) int {
	/*
		METHOD: Map / Dict
		Для решения этой задачи мы будем использовать подход, основанный на двух проходах по массиву.
		В первом проходе мы будем подсчитывать частоту каждого элемента и его первое и последнее вхождение в массив.
		Во втором проходе мы найдем элемент с максимальной частотой и длину наименьшего подмассива,
		который содержит все элементы этой частоты.

		TIME COMPLEXITY: O(n), где n - количество элементов в массиве,
		потому что мы проходим по массиву два раза: один раз для подсчета частоты и индексов,
		и второй раз для поиска минимальной длины подмассива.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить информацию о каждом элементе в массиве.
	*/
	// Создаем структуру для хранения информации о каждом элементе
	type ElementInfo struct {
		firstSeen int
		lastSeen  int
		frequency int
	}

	// Создаем словарь для отслеживания информации о каждом элементе
	elementMap := make(map[int]ElementInfo)

	// Первый проход по массиву для подсчета частоты и индексов
	for i, num := range nums {
		if info, exists := elementMap[num]; exists {
			info.lastSeen = i
			info.frequency++
			elementMap[num] = info
		} else {
			elementMap[num] = ElementInfo{firstSeen: i, lastSeen: i, frequency: 1}
		}
	}

	// Инициализируем переменные для отслеживания максимальной частоты и минимальной длины подмассива
	maxFrequency := 0
	minLength := len(nums)

	// Второй проход по массиву для поиска минимальной длины подмассива с максимальной частотой
	for _, info := range elementMap {
		if info.frequency > maxFrequency {
			maxFrequency = info.frequency
			minLength = info.lastSeen - info.firstSeen + 1
		} else if info.frequency == maxFrequency {
			length := info.lastSeen - info.firstSeen + 1
			if length < minLength {
				minLength = length
			}
		}
	}

	// Возвращаем минимальную длину подмассива
	return minLength
}
