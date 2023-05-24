package number_complement

import "math/bits"

func findComplementV1(num int) int {
	return (1<<(bits.Len(uint(num))) - 1) ^ num
}
