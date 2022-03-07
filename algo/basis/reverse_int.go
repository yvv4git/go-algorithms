package basis

import "math"

func reverseInt(x int) int {
	var result int
	result = 0

	negative := false
	if x < 0 {
		negative = true
		x = -x
	}

	for x > 0 {
		result = result*10 + x%10
		x = x / 10
	}

	if result >= math.MaxInt32 {
		return 0
	}

	if negative {
		return -result
	}

	return result
}
