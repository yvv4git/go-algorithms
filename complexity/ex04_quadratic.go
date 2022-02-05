//go:build quadratic

package main

import "github.com/pterm/pterm"

var inputData = [5]int{5, 4, 3, 2, 1}

func main() {
	insertionSort(inputData[:])
	pterm.Info.Println("Insert sort: ", inputData)
}

// Quadratic complexity - O(n^2)
// Сортировка вставками в классическом исполнении имеет квадратичную сложность.
// Обрати внимание, что по входящему массиву данных приходится проходить 2 раза.
func insertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		x := ar[i]
		j := i
		for ; j >= 1 && ar[j-1] > x; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = x
	}
}
