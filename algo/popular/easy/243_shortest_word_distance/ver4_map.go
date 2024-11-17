package main

import (
	"math"
)

func shortestDistanceV4(wordsDict []string, word1 string, word2 string) int {
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

	// Находим минимальное расстояние между индексами
	for _, i1 := range indices1 {
		for _, i2 := range indices2 {
			distance := int(math.Abs(float64(i1 - i2)))
			if distance < minDistance {
				minDistance = distance
			}
		}
	}

	return minDistance
}

// func main() {
// 	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
// 	word1 := "coding"
// 	word2 := "practice"

// 	result := shortestDistance(wordsDict, word1, word2)
// 	fmt.Println("Кратчайшее расстояние:", result)
// }
