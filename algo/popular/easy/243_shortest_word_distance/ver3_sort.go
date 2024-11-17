package main

import (
	"math"
	"sort"
)

func shortestDistanceV3(wordsDict []string, word1 string, word2 string) int {
	// Создаем список для хранения индексов
	var indices []int

	// Заполняем список индексами слов
	for i, word := range wordsDict {
		if word == word1 || word == word2 {
			indices = append(indices, i)
		}
	}

	// Сортируем список индексов
	sort.Ints(indices)

	// Инициализируем минимальное расстояние
	minDistance := math.MaxInt32

	// Находим минимальное расстояние между соседними индексами
	for i := 1; i < len(indices); i++ {
		distance := indices[i] - indices[i-1]
		if distance < minDistance {
			minDistance = distance
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
