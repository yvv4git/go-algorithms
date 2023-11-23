package _74_guess_number_higher_or_lower

func interpolationSearch(key int, array []int) int {
	min, max := array[0], array[len(array)-1]

	low, high := 0, len(array)-1

	for low <= high && key >= min && key <= max {
		// make a guess of the location
		guess := low + (high-low)*(key-min)/(max-min)

		// maybe we found it?
		if array[guess] == key {
			// we found it, return the index
			return guess
		}

		// if we guessed too high
		if array[guess] > key {
			// adjust 'high' to guess - 1
			high = guess - 1
			max = array[high]
		} else {
			// adjust 'low' to guess + 1
			low = guess + 1
			min = array[low]
		}
	}

	return -1
}
