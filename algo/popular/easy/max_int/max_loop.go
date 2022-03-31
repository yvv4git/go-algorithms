package main

// MaxByLoop - used for find max element in array
func MaxByLoop(arr []int) int {
	maxEnd := arr[0]
	for _, val := range arr {
		if maxEnd < val {
			maxEnd = val
		}
	}

	return maxEnd
}
