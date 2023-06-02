package binary_number_alternating_bits

func hasAlternatingBitsV3(n int) bool {
	currentRightBit := n & 1

	for n > 0 {
		if n&1 != currentRightBit {
			return false
		}

		currentRightBit = 1 ^ currentRightBit
		n >>= 1
	}

	return true
}
