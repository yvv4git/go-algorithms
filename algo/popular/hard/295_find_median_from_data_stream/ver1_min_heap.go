package main

import (
	"container/heap"
	"fmt"
)

// MinHeap представляет минимальную кучу
type MinHeap []int

// Len возвращает длину кучи
// Временная сложность: O(1)
func (h MinHeap) Len() int { return len(h) }

// Less сравнивает элементы кучи
// Временная сложность: O(1)
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap меняет местами элементы кучи
// Временная сложность: O(1)
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push добавляет элемент в кучу
// Временная сложность: O(log n)
// Пространственная сложность: O(1)
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop удаляет и возвращает наименьший элемент из кучи
// Временная сложность: O(log n)
// Пространственная сложность: O(1)
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxHeap представляет максимальную кучу
type MaxHeap []int

// Len возвращает длину кучи
// Временная сложность: O(1)
func (h MaxHeap) Len() int { return len(h) }

// Less сравнивает элементы кучи
// Временная сложность: O(1)
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

// Swap меняет местами элементы кучи
// Временная сложность: O(1)
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push добавляет элемент в кучу
// Временная сложность: O(log n)
// Пространственная сложность: O(1)
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop удаляет и возвращает наибольший элемент из кучи
// Временная сложность: O(log n)
// Пространственная сложность: O(1)
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MedianFinder структура для поиска медианы
type MedianFinder struct {
	minHeap *MinHeap // Минимальная куча для хранения большей половины элементов
	maxHeap *MaxHeap // Максимальная куча для хранения меньшей половины элементов
}

// Constructor для инициализации MedianFinder
func Constructor() MedianFinder {
	minHeap := &MinHeap{}
	maxHeap := &MaxHeap{}
	heap.Init(minHeap)
	heap.Init(maxHeap)
	return MedianFinder{minHeap: minHeap, maxHeap: maxHeap}
}

// AddNum добавляет число в структуру данных
// Временная сложность: O(log n)
// Пространственная сложность: O(n)
func (this *MedianFinder) AddNum(num int) {
	if this.maxHeap.Len() == 0 || num <= (*this.maxHeap)[0] {
		heap.Push(this.maxHeap, num)
	} else {
		heap.Push(this.minHeap, num)
	}

	// Убеждаемся, что разница в размере двух куч не превышает 1
	if this.maxHeap.Len() > this.minHeap.Len()+1 {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap))
	} else if this.minHeap.Len() > this.maxHeap.Len()+1 {
		heap.Push(this.maxHeap, heap.Pop(this.minHeap))
	}
}

// FindMedian находит медиану текущего набора данных
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() == this.minHeap.Len() {
		return float64((*this.maxHeap)[0]+(*this.minHeap)[0]) / 2.0
	} else if this.maxHeap.Len() > this.minHeap.Len() {
		return float64((*this.maxHeap)[0])
	} else {
		return float64((*this.minHeap)[0])
	}
}

func main() {
	/*
		METHOD: Heap
		Задача решена с использованием двух куч (heap): одной максимальной кучи для хранения меньшей половины элементов
		и одной минимальной кучи для хранения большей половины элементов.
		Этот метод позволяет эффективно поддерживать медиану при добавлении новых элементов.

		TIME COMPLEXITY:
		AddNum. O(log n) — Вставка элемента в кучу требует O(log n) времени из-за необходимости поддержания свойства кучи.
		FindMedian. O(1) — Доступ к корню кучи и вычисление медианы выполняются за постоянное время.

		SPACE COMPLEXITY:
		AddNum. O(n) — В худшем случае, все элементы будут храниться в одной из куч.
		FindMedian. O(1) — Для вычисления медианы не требуется дополнительной памяти, кроме уже существующих элементов в кучах.
	*/
	medianFinder := Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	fmt.Println(medianFinder.FindMedian()) // Вывод: 1.5
	medianFinder.AddNum(3)
	fmt.Println(medianFinder.FindMedian()) // Вывод: 2.0
}
