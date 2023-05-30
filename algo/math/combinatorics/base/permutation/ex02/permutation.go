package ex02

import "fmt"

// Permutation - перестановка элементов
func Permutation(data []int, i int, length int) {
	if length == i {
		fmt.Println(data)
		return
	}

	for j := i; j < length; j++ {
		swap(data, i, j)
		Permutation(data, i+1, length)
		swap(data, i, j)
	}
}

func swap(data []int, x, y int) {
	data[x], data[y] = data[y], data[x]
}
