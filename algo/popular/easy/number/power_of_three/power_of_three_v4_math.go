package power_of_three

import "math"

func isPowerOfThreeV4(n int) bool {
	if n < 1 {
		return false
	}

	// dividing log(n) by log(3) gives us log_3(n)
	x := math.Round(math.Log(float64(n)) / math.Log(3))

	return n == int(math.Pow(3, x))
}
