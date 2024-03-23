package main

import (
	"fmt"
	"math"
)

// Функция для разделения массива и получения индекса опорного элемента
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Функция для вставки сортировки
func insertionSort(arr []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		key := arr[i]
		j := i - 1
		for j >= left && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Функция для пирамидальной сортировки
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[i] < arr[left] {
		largest = left
	}

	if right < n && arr[largest] < arr[right] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int) {
	n := len(arr)

	// Построение пирамидального дерева
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Один за другим извлекаем элементы из кучи
	for i := n - 1; i >= 0; i-- {
		// Перемещаем текущий корень в конец
		arr[0], arr[i] = arr[i], arr[0]

		// Вызываем процедуру heapify на уменьшенной куче
		heapify(arr, i, 0)
	}
}

// Функция для определения оптимального размера массива для вставки сортировки
func getThreshold(n int) int {
	return 16
}

// IntroSort
func introSort(arr []int, maxDepth int) {
	n := len(arr)
	if n <= 1 {
		return
	} else if n <= getThreshold(n) {
		insertionSort(arr, 0, n-1)
	} else if maxDepth == 0 {
		heapSort(arr)
	} else {
		p := partition(arr, 0, n-1)
		introSort(arr[:p], maxDepth-1)
		introSort(arr[p+1:], maxDepth-1)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	maxDepth := int(math.Log2(float64(len(arr)))) * 2
	introSort(arr, maxDepth)
	fmt.Println("Sorted array is:", arr)
}
