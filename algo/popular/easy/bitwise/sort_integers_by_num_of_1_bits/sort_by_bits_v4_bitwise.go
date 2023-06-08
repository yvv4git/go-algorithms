package sort_integers_by_num_of_1_bits

import "sort"

// 423775 | 2577 ns/op | 472 B/op | 3 allocs/op
func sortByBitsV4(arr []int) []int {
	sort.SliceStable(arr, func(a, b int) bool {
		bitsA := countBits(arr[a]) // Самостоятельно посчитаем количество разрядов с 1
		bitsB := countBits(arr[b])
		if bitsA == bitsB {
			return arr[a] < arr[b]
		}

		return bitsA < bitsB
	})

	return arr
}

func countBits(x int) int {
	count := 0
	for x != 0 {
		count += x & 1
		x >>= 1
	}
	return count
}
