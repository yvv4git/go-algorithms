package non_overlapping_interfals

func eraseOverlapIntervalsV2(intervals [][]int) int {
	/*
		METHOD:
		Time complexity: O(n*log(n))
		Space complexity: O(1)
	*/
	cnt, result := len(intervals), 0
	if cnt <= 1 {
		return result
	}

	// Time: O(cnt*log(n))
	sortIntervals(intervals)

	top := 0

	// Интервалы пересекаются, если их y интервала больше чем x другого интервала.
	for i := 1; i < cnt; i++ {
		if intervals[top][1] <= intervals[i][0] {
			top = i
		} else {
			if intervals[top][1] > intervals[i][1] {
				top = i
			}
			result++
		}
	}

	return result
}
