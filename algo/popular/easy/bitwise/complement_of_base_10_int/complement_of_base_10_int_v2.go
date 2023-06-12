package complement_of_base_10_int

import "math/bits"

func bitwiseComplementV2(n int) int {
	if n == 0 {
		return 1
	}

	shift := bits.Len(uint(n))

	return n ^ (1<<shift - 1)
}
