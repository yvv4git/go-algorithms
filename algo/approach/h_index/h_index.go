package main

import (
	"fmt"
	"sort"
)

// Функция для вычисления индекса Хирша
func hIndex(citations []int) int {
	/*
		METHOD: Binary search
		Этот код сначала сортирует массив цитирований в порядке убывания, а затем проходит по массиву, чтобы найти индекс Хирша.
		Если количество цитирований текущей статьи больше или равно текущему индексу (i + 1), то мы увеличиваем значение индекса Хирша.
		Как только мы находим статью, у которой количество цитирований меньше текущего индекса,
		мы останавливаемся и возвращаем текущее значение индекса Хирша.

		TIME COMPLEXITY: O(n log n). Такую сложность получаем потому, что сначала сортируем числа O(n log n).
		Потом проходим по массиву за O(n). В итоге, временная сложность будет O(n log n)

		SPACE COMPLEXITY: O(n), так как в худшем случае может потребовать памяти O(n).
	*/
	// Сортируем массив цитирований в порядке убывания
	sort.Sort(sort.Reverse(sort.IntSlice(citations))) // O(n log n)

	// Находим индекс Хирша
	h := 0
	for i := 0; i < len(citations); i++ { // O(n)
		if citations[i] >= i+1 {
			h = i + 1
		} else {
			break
		}
	}

	return h
}

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println("Индекс Хирша:", hIndex(citations))
}
