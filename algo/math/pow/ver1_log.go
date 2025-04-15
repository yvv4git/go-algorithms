package main

import (
	"fmt"
)

// Power возводит число base в степень exponent за O(log n)
func Power(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}

	negative := exponent < 0
	if negative {
		base = 1 / base
		exponent = -exponent
	}

	result := 1.0
	for exponent > 0 {
		if exponent%2 == 1 {
			result *= base
		}
		base *= base
		exponent /= 2
	}

	return result
}

func main() {
	fmt.Println(Power(2, 10))  // 1024
	fmt.Println(Power(2, -2))  // 0.25
	fmt.Println(Power(5, 0))   // 1
	fmt.Println(Power(2.5, 3)) // 15.625
}
