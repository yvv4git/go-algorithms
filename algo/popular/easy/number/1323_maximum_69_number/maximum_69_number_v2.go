package _323_maximum_69_number

func maximum69NumberV2(num int) int {
	/*
		METHOD: Math
		TIME COMPLEXITY: O(1) - because we have not more than 4 operations.
		Space complexity: O(1) - because there are not additional memory for slices.
	*/
	div := 1000 // power of 10
	n := num

	for div > 0 {
		if n/div == 6 {
			num += 3 * div // replace 6 with 9
			return num
		} else {
			n = n % div    // move to next digit
			div = div / 10 // lower power of 10
		}
	}

	return num
}
