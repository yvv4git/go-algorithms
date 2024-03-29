package non_overlapping_interfals

func eraseOverlapIntervalsV1(intervals [][]int) int {
	/*
		METHOD:
		TIME COMPLEXITY: O(nlogn + n) = O(nlogn)
		SPACE COMPLEXITY: O(1)
	*/
	if len(intervals) <= 1 {
		return 0
	}

	// Time: O(nlogn)
	sortIntervals(intervals)

	// Time: O(n)
	nbNonOverlappingIntervals := numberOfNonOverlappingIntervals(intervals)

	return len(intervals) - nbNonOverlappingIntervals
}

func numberOfNonOverlappingIntervals(intervals [][]int) int {
	/*
		METHOD: Loop
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	currentInterval := intervals[0]
	var nbNonOverlappingIntervals int = 1

	// Time: O(n)
	for _, interval := range intervals[1:] {
		if interval[0] >= currentInterval[1] {
			currentInterval = interval
			nbNonOverlappingIntervals++
		}
	}

	return nbNonOverlappingIntervals
}
