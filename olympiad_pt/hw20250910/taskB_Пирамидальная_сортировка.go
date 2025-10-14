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

	// Выполняем пирамидальную сортировку
	heapSort(arr)

	// Выводим результат
	for i, num := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()
}

func heapSort(arr []int) {
	n := len(arr)

	// Построение max-heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Извлекаем элементы из кучи по одному
	for i := n - 1; i > 0; i-- {
		// Перемещаем текущий корень в конец
		arr[0], arr[i] = arr[i], arr[0]

		// Вызываем heapify на уменьшенной куче
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i     // Инициализируем наибольший элемент как корень
	left := 2*i + 1  // Левый дочерний элемент
	right := 2*i + 2 // Правый дочерний элемент

	// Если левый дочерний элемент больше корня
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// Если правый дочерний элемент больше, чем самый большой элемент на данный момент
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// Если самый большой элемент не корень
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		// Рекурсивно heapify поддерево
		heapify(arr, n, largest)
	}
}
