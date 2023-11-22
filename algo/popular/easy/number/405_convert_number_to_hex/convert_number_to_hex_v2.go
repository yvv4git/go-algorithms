package _05_convert_number_to_hex

func toHexV2(num int) string {
	if num == 0 {
		return "0"
	}

	var i int
	bytes := []byte("0123456789abcdef        ") // 8 spaces

	for i = 23; num != 0; i-- {
		bytes[i] = bytes[num&15]
		num = int(uint32(num) >> 4) // Деление на 16
	}

	return string(bytes[i+1:])
}
