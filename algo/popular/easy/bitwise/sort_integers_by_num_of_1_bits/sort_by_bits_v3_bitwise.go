package sort_integers_by_num_of_1_bits

import (
	"math/bits"
	"sort"
)

// 562573 | 1948 ns/op | 472 B/op | 3 allocs/op
func sortByBitsV3(arr []int) []int {
	/*
		METHOD: Bitwise
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	sort.SliceStable(arr, func(a, b int) bool {
		bitsA := bits.OnesCount(uint(arr[a])) // Можно использовать готовую функцию для подсчета битов-1
		bitsB := bits.OnesCount(uint(arr[b]))
		if bitsA == bitsB {
			return arr[a] < arr[b]
		}

		return bitsA < bitsB
	})

	return arr
}
