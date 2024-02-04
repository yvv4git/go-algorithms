package main

import (
	"fmt"
	"strings"
)

func addBinaryV4(a string, b string) string {
	/*
		METHOD: Bitwise
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)
	*/
	// have to catch this case because of the trim at the end of the function
	if a == "0" && b == "0" {
		return "0"
	}
	var n, m, k, c uint64
	i, j := len(a)-1, len(b)-1
	// used to convert between character and uint64
	zero := uint64('0')
	// mask used to read/write the "carry" (64th) bit of the uint64 numbers we're dealing with
	var mask uint64 = 1 << 63
	var result string
	for i >= 0 || j >= 0 {
		// if i is still in bound, add its digit to the appropriate position of n
		if i >= 0 {
			// convert the '0' or '1' character to either a 1 or 0 uint64 and shift it left to the appropriate
			// position then set that bit in n
			n |= (uint64(a[i]) - zero) << k
			// decrement i since we start at the end of the string
			i--
		}
		// same for j
		if j >= 0 {
			m |= (uint64(b[j]) - zero) << k
			j--
		}

		// increment k
		k++
		// if we haven't accumulated 63 bits and j and i aren't both emtpy
		if k < 63 && (j >= 0 || i >= 0) {
			continue
		}

		// We accumulated 63 bits or ran out of bits from j and i. Now we add the numbers and the carry bit
		// from the previous add operation last cycle

		// reset k to 0
		k = 0

		// calculate sum
		sum := n + m + c

		// extract the carry from the 64th bit
		c = (sum & mask) >> 63
		// clear the bit
		sum &^= mask

		// clear n and m
		n, m = 0, 0
		// format as binary number of width 63 with leading zeros
		result = fmt.Sprintf("%063b", sum) + result
	}
	// trim any zeros on the left side
	return strings.TrimLeft(result, "0")
}
