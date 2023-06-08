package main

func addBinaryV3(a string, b string) string {
	/*
		Method: Binary addition
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	var result string

	if len(b) > len(a) {
		a, b = b, a
	}

	if len(b) < len(a) {
		k := len(a) - len(b)
		for i := 0; i < k; i++ {
			b = "0" + b
		}
	}

	var carry = 0

	for i := len(a) - 1; i > -1; i-- {
		if a[i] == '0' && b[i] == '0' && carry == 0 {
			result = "0" + result

		} else if a[i] == '0' && b[i] == '0' && carry == 1 {
			result = "1" + result
			carry = 0

		} else if ((a[i] == '0' && b[i] == '1') || (a[i] == '1' && b[i] == '0')) && carry == 0 {
			result = "1" + result

		} else if ((a[i] == '0' && b[i] == '1') || (a[i] == '1' && b[i] == '0')) && carry == 1 {
			result = "0" + result

		} else if carry == 0 {
			result = "0" + result
			carry = 1

		} else {
			result = "1" + result
			carry = 1
		}
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}
