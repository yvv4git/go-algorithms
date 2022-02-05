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
