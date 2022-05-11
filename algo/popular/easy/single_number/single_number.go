package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0

	for _, n := range nums {
		fmt.Printf("res=%d  n=%d ===> %d\n", res, n, res^n)
		res ^= n
	}

	return res
}
