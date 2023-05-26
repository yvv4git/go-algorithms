package convert_number_to_hex

func toHexV3(num int) string {
	arr := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

	if num == 0 {
		return "0"
	}

	if num < 0 {
		num = -num
		num ^= 0xffffffff
		num += 1
	}

	result := ""

	for num > 0 {
		lowerByte := num % 16
		result = string(arr[lowerByte]) + result
		num /= 16
	}

	return result
}
