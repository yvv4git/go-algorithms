package selection

import (
	"github.com/yvv4git/go-algorithms/algo/math/sorts"
)

// SelectSort - used as selection sort
func SelectSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		min := i
		for j := i + 1; j < len(ar); j++ {
			if ar[min] > ar[j] {
				min = j
			}
		}

		if min != i {
			sorts.Swap(ar, i, min)
		}
	}
}
