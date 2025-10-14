//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	// Читаем количество элементов
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Читаем массив
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
	}

	// Сортируем массив
	quickSortOptimized(arr, 0, n-1)

	// Выводим результат
	for i, num := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()
}

func quickSortOptimized(arr []int, low, high int) {
	if low < high {
		// Выбираем опорный элемент как медиану из трех
		mid := low + (high-low)/2
		if arr[low] > arr[mid] {
			arr[low], arr[mid] = arr[mid], arr[low]
		}
		if arr[low] > arr[high] {
			arr[low], arr[high] = arr[high], arr[low]
		}
		if arr[mid] > arr[high] {
			arr[mid], arr[high] = arr[high], arr[mid]
		}

		// Помещаем медиану в конец
		arr[mid], arr[high] = arr[high], arr[mid]

		pi := partition(arr, low, high)
		quickSortOptimized(arr, low, pi-1)
		quickSortOptimized(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
