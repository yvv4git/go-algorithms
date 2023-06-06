package binary_gap

import (
	"sort"
	"strconv"
)

func binaryGapV1(n int) int {
	/*
		Method: NoBitwise
		Time complexity: O(n + (n Log(n))) = O(n)
		Space complexity: O(n)
	*/
	str := strconv.FormatInt(int64(n), 2)
	ranges := []int{}
	sum := 0

	// O(n)
	for _, s := range str {
		if s == '1' {
			ranges = append(ranges, sum)
			sum = 1
		} else {
			sum++
		}
	}
	sort.Ints(ranges) // O(n Log(n))

	return ranges[len(ranges)-1]
}
