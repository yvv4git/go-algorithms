package _05_convert_number_to_hex

func toHexV1(num int) string {
	/*
		Method: Bitwise
		Time complexity: O(1)
		Space complexity: O(1)
	*/
	const (
		hexCharacters = "0123456789abcdef"
		mask          = 0xf
	)
	var result = make([]byte, 8)
	var nonzeroIndex = 7

	for i := 7; i >= 0; i-- {
		val := num & mask
		if val > 0 {
			nonzeroIndex = i
		}
		result[i] = hexCharacters[val]
		num >>= 4
	}
	return string(result[nonzeroIndex:])
}
