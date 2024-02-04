package complement_of_base_10_int

import (
	"math"
	"math/bits"
)

func bitwiseComplementV1(n int) int {
	/*
		TIME COMPLEXITY: ???
		SPACE COMPLEXITY: ???
	*/
	if n == 0 {
		return 1
	}

	l := bits.Len(uint(n))
	tmp := math.Pow(2, float64(l)) - 1

	return n ^ int(tmp)
}
