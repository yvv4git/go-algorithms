package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SearchResult - структура для представления результата поиска с релевантностью
type SearchResult struct {
	ID        int
	Relevance int
}

// Функция для нахождения k-го наименьшего элемента в массиве
func quickselect(results []SearchResult, k int) SearchResult {
	/*
		METHOD: Quickselect
		Объяснение метода: Quickselect — это алгоритм выбора, основанный на принципе "разделяй и властвуй".
		Он использует случайный опорный элемент для разделения массива на части и рекурсивно находит k-й наименьший элемент.

		TIME COMPLEXITY: В среднем O(n), в худшем случае O(n^2)
		Объяснение временной сложности: В среднем алгоритм работает за линейное время, так как каждый раз массив делится на две примерно равные части.
		В худшем случае, когда опорный элемент выбирается неудачно, временная сложность может достигать O(n^2).

		SPACE COMPLEXITY: O(n)
		Объяснение пространственной сложности: Алгоритм использует дополнительную память для хранения разделенных частей массива.
	*/
	if len(results) == 1 {
		return results[0]
	}

	// Выбираем случайный опорный элемент
	rand.Seed(time.Now().UnixNano())
	pivotIndex := rand.Intn(len(results))
	pivot := results[pivotIndex]

	// Разделяем массив на элементы меньше, равные и больше опорного
	var lows, highs, pivots []SearchResult
	for _, result := range results {
		if result.Relevance < pivot.Relevance {
			lows = append(lows, result)
		} else if result.Relevance > pivot.Relevance {
			highs = append(highs, result)
		} else {
			pivots = append(pivots, result)
		}
	}

	// Рекурсивно находим k-й наименьший элемент
	if k <= len(lows) {
		return quickselect(lows, k)
	} else if k <= len(lows)+len(pivots) {
		return pivots[0]
	} else {
		return quickselect(highs, k-len(lows)-len(pivots))
	}
}

func main() {
	/*
		В этом примере мы будем находить топ-k результатов по релевантности.
	*/
	// Пример массива результатов поиска
	results := []SearchResult{
		{ID: 1, Relevance: 5},
		{ID: 2, Relevance: 3},
		{ID: 3, Relevance: 8},
		{ID: 4, Relevance: 2},
		{ID: 5, Relevance: 7},
	}

	k := 3 // Найти топ-3 результата

	// Находим k-й наименьший элемент
	kthSmallest := quickselect(results, k)

	// Фильтруем результаты, чтобы получить топ-k результатов
	var topKResults []SearchResult
	for _, result := range results {
		if result.Relevance >= kthSmallest.Relevance {
			topKResults = append(topKResults, result)
		}
	}

	// Сортируем топ-k результатов по убыванию релевантности
	for i := 0; i < len(topKResults)-1; i++ {
		for j := i + 1; j < len(topKResults); j++ {
			if topKResults[i].Relevance < topKResults[j].Relevance {
				topKResults[i], topKResults[j] = topKResults[j], topKResults[i]
			}
		}
	}

	// Выводим топ-k результатов
	fmt.Printf("Top %d results:\n", k)
	for _, result := range topKResults {
		fmt.Printf("ID: %d, Relevance: %d\n", result.ID, result.Relevance)
	}
}
