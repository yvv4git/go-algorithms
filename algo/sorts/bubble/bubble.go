package bubble

import "github.com/yvv4git/go-algorithms/algo/sorts"

// BubbleSort - used as bubble sort
func BubbleSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				sorts.Swap(ar, i, j)
			}
		}
	}
}

// BubbleSort2 - вариант, где мы сравниваем соседние пары и гоним пузырек наверх.
func BubbleSort2(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := len(ar) - 1; j > i; j-- {
			if ar[j-1] > ar[j] {
				sorts.Swap(ar, j-1, j)
			}
		}
	}
}

// BubbleSort3 - вариант с тонущим пузырьком (sinking sort), когда тяжелый пузырек опускается вниз.
func BubbleSort3(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				sorts.Swap(ar, j-1, j)
			}
		}
	}
}
