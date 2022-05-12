package bubble

import (
	"github.com/yvv4git/go-algorithms/algo/sorts"
)

// BubbleSort1 - used as bubble sort. Здесь самый простой пример реализации. Прям очевидный
func BubbleSort1(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				ar[i], ar[j] = ar[j], ar[i] // меняем местами
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

// BubbleSort4 - еще один вариант сортировки пузырьком
func BubbleSort4(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ { // тут понятно - идем до конца массива
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // меняем местами
			}
		}
	}
}
