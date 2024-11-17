package main

import (
	"math"
	"sort"
)

func shortestDistanceV2(wordsDict []string, word1 string, word2 string) int {
	/*
		METHOD: Binary Search
		TIME COMPLEXITY: O(nlogn)
		SPACE COMPLEXITY: O(n)
	*/
	// Создаем хэш-таблицу для хранения индексов слов
	wordIndices := make(map[string][]int)

	// Заполняем хэш-таблицу индексами слов
	for i, word := range wordsDict {
		wordIndices[word] = append(wordIndices[word], i)
	}

	// Получаем списки индексов для word1 и word2
	indices1 := wordIndices[word1]
	indices2 := wordIndices[word2]

	// Инициализируем минимальное расстояние
	minDistance := math.MaxInt32

	// Для каждого индекса из indices1 находим ближайший индекс из indices2
	for _, i1 := range indices1 {
		// Используем бинарный поиск для нахождения ближайшего индекса
		pos := sort.SearchInts(indices2, i1)
		if pos < len(indices2) {
			minDistance = int(math.Min(float64(minDistance), math.Abs(float64(i1-indices2[pos]))))
		}
		if pos > 0 {
			minDistance = int(math.Min(float64(minDistance), math.Abs(float64(i1-indices2[pos-1]))))
		}
	}

	return minDistance
}

// func main() {
// 	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
// 	word1 := "coding"
// 	word2 := "practice"

// 	result := shortestDistanceV2(wordsDict, word1, word2)
// 	fmt.Println("Кратчайшее расстояние:", result)
// }
