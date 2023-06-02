package binary_number_alternating_bits

func hasAlternatingBitsV2(n int) bool {
	/*
		Method: Bitwise
		Time complexity: ???
		Space complexity: O(1)
	*/
	prevBit := n & 1 // Если младший(правый) бит будет 1, то & даст 1, иначе 0.

	// Начинаем со второго бита и сдвигаем биты вправо на каждой итерации.
	for curNum := n >> 1; curNum > 0; curNum >>= 1 {
		curBit := curNum & 1 // Сравниваем текущий бит с 1. На выходе получаем либо 0, либо 1.
		if curBit == prevBit {
			return false
		}
		prevBit = curBit
	}

	return true
}
