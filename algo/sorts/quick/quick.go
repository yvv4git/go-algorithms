package quick

import (
	"github.com/yvv4git/go-algorithms/algo/sorts"
)

func QuickSort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	QuickSort(ar[:split])
	QuickSort(ar[split:])
}

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
