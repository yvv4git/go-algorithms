package _78_first_bad_version

func firstBadVersionV1(n int) int {
	/*
		Method: Binary search
		Time complexity: O(log n)
		Space complexity: O(n)
	*/
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}

	return false
}
