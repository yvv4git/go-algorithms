package quick

import (
	"github.com/yvv4git/go-algorithms/algo/math/sorts"
)

// QuickSort сортирует массив по алгоритму быстрой сортировки.
// Approach: Divide and conquer, Recursion.
// Временная сложность:
//   - Лучший случай: O(n log n)
//   - Средний случай: O(n log n)
//   - Худший случай: O(n²)
//
// Пространственная сложность:
//   - O(log n) в лучшем/среднем случае
//   - O(n) в худшем случае
func QuickSort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	QuickSort(ar[:split])
	QuickSort(ar[split:])
}

// partition разделяет массив на две части относительно опорного элемента (pivot)
// и возвращает индекс разделения. Элементы меньше pivot перемещаются влево,
// элементы больше pivot - вправо.
func partition(ar []int) int {
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		sorts.Swap(ar, left, right)
	}
}
