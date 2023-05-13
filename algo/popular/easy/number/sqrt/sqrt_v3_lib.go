package sqrt

import "math"

func mySqrtV3(x int) int {
	if x >= 0 {
		s := math.Sqrt(float64(x))

		return int(s)
	}

	return 0
}
