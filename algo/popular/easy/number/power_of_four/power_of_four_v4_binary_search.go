package power_of_four

import "math"

func isPowerOfFourV4(n int) bool {
	/*
		METHOD: Binary search
		TIME COMPLEXITY: O(log(n))
		SPACE COMPLEXITY: O(1)
	*/
	left := 0
	right := 16
	for right-left > 1 {
		mid := math.Floor(float64(left+right) / float64(2))
		q := int(math.Pow(4, mid))
		if q == n {
			return true
		}
		if q > n {
			right = int(mid)
		} else {
			left = int(mid)
		}
	}
	if int(math.Pow(4, float64(left))) == n {
		return true
	}
	return false
}
