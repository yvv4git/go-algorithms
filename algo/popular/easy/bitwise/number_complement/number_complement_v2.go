package number_complement

func findComplementV2(num int) int {
	x := ^0
	for x&num > 0 {
		x <<= 1
	}

	return ^x ^ num
}
