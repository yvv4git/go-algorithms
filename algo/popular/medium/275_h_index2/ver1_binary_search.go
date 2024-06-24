package main

import "fmt"

// Функция для вычисления индекса Хирша
func hIndex(citations []int) int {
	/*
		METHOD: Binary search
		Мы используем бинарный поиск для нахождения индекса Хирша, так как массив уже отсортирован в порядке убывания.

		TIME COMPLEXITY: O(log n), так как мы используем бинарный поиск для нахождения индекса Хирша.

		SPACE COMPLEXITY: O(1), так как мы используем только фиксированное количество переменных
		и не используем дополнительную память, зависящую от размера входных данных.
	*/
	n := len(citations)
	left, right := 0, n-1

	// Используем бинарный поиск для нахождения индекса Хирша
	for left <= right {
		mid := left + (right-left)/2
		if citations[mid] == n-mid {
			// Если найдено точное совпадение, возвращаем индекс Хирша
			return n - mid
		} else if citations[mid] > n-mid {
			// Если текущее количество цитирований больше, чем количество оставшихся публикаций,
			// двигаемся влево
			right = mid - 1
		} else {
			// В противном случае двигаемся вправо
			left = mid + 1
		}
	}

	// Если цикл завершился, возвращаем индекс Хирша
	return n - left
}

func main() {
	citations := []int{10, 8, 5, 3, 3}
	fmt.Println("Индекс Хирша:", hIndex(citations))
}
