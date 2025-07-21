package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// CharFrequency хранит символ и его частоту
type CharFrequency struct {
	char  rune
	count int
}

// PriorityQueue реализует max-heap на основе частот
type PriorityQueue []*CharFrequency

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].count > pq[j].count }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*CharFrequency)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Approach: Priority Queue
// frequencySort сортирует символы строки по убыванию частоты с использованием кучи
// Временная сложность: O(n log k), где k - количество уникальных символов
// Пространственная сложность: O(n) для хранения частот и кучи
func frequencySort(s string) string {
	// Подсчет частот символов
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	// Создание и заполнение кучи
	pq := make(PriorityQueue, 0, len(freq))
	heap.Init(&pq)
	for char, count := range freq {
		heap.Push(&pq, &CharFrequency{char, count})
	}

	// Построение результата
	var builder strings.Builder
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*CharFrequency)
		builder.WriteString(strings.Repeat(string(item.char), item.count))
	}

	return builder.String()
}

func main() {
	// Пример использования
	examples := []string{
		"tree",
		"cccaaa",
		"Aabb",
	}

	for _, example := range examples {
		fmt.Printf("Input: %q\n", example)
		fmt.Printf("Output: %q\n\n", frequencySort(example))
	}
}
