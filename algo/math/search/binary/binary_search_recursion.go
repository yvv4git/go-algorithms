package binary

func binarySearchRecursion(list []int, target int) int {
	low, max := 0, len(list)-1

	return searchInPart(list, low, max, target)
}

func searchInPart(list []int, low, max, target int) int {
	if (max - low) < 0 {
		return -1
	}

	mid := low + (max-low)/2
	if list[mid] == target {
		return mid
	} else if list[mid] < target {
		return searchInPart(list, mid+1, max, target)
	} else {
		return searchInPart(list, low, mid-1, target)
	}
}
