package main

import "strconv"

var mem = 0

func addBinaryV2(a string, b string) string {
	/*
		METHOD: Recursion
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	if a == "" && b == "" {
		if mem != 0 {
			mem = 0
			return "1"
		}

		mem = 0
		return ""
	}

	if a == "" {
		a = "0"
	}

	if b == "" {
		b = "0"
	}

	numA, _ := strconv.Atoi(a[len(a)-1:])
	numB, _ := strconv.Atoi(b[len(b)-1:])

	sum := numA + numB + mem
	mod := sum % 2

	res := strconv.Itoa(mod)

	if sum > 1 && (numA == 1 || numB == 1) {
		mem = 1
	} else {
		mem = 0
	}

	return addBinaryV2(a[:len(a)-1], b[:len(b)-1]) + res
}
