package binary

func binarySearch(list []int, target int) int {
	var low, max = 0, len(list) - 1
	if max < low {
		return -1
	}

	for low <= max {
		var mid = low + (max-low)/2
		if list[mid] == target {
			return mid
		}

		if target > list[mid] {
			low = mid + 1
		} else {
			max = mid - 1
		}
	}

	return -1
}
