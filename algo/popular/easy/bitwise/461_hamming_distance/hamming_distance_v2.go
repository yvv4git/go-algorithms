package _61_hamming_distance

import "math/bits"

func hammingDistanceV2(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
