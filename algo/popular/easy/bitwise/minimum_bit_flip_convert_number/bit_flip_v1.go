package minimum_bit_flip_convert_number

import "math/bits"

func minBitFlipsV1(start, goal int) int {
	// User XOR(^) for finding different bits
	// then count result bits
	return bits.OnesCount32(uint32(start ^ goal))
}
