package non_overlapping_interfals

import "sort"

func sortIntervals(intervals [][]int) {
	/*
		METHOD: Default sort.
		TIME COMPLEXITY: O(n*log(n))
		SPACE COMPLEXITY: O(1)
	*/
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
}
