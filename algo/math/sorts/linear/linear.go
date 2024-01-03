package linear

import "github.com/yvv4git/go-algorithms/algo/math/sorts"

// LinearSortBigOrder - used for sort list of elements with start big order
func LinearSortBigOrder(list []int) {
	for i := 0; i < len(list); i++ {
		max := i
		for j := i + 1; j < len(list); j++ {
			if list[max] < list[j] {
				max = j
			}
		}

		if max != i {
			sorts.Swap(list, i, max)
		}
	}
}

// LinearSortLittleOrder - used for sort list of element with start little order
func LinearSortLittleOrder(list []int) {
	for i := 0; i < len(list); i++ {
		min := i
		for j := i + 1; j < len(list); j++ {
			if list[min] > list[j] {
				min = j
			}
		}

		if min != i {
			sorts.Swap(list, i, min)
		}
	}
}
