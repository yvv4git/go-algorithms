package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	Nums []int
}

func Constructor(nums []int) Solution {
	var n Solution
	n.Nums = nums
	return n
}

func (this *Solution) Reset() []int {
	return this.Nums
}

func (this *Solution) Shuffle() []int {
	slice := make([]int, len(this.Nums))
	copy(slice, this.Nums)
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	solution := Constructor(nums)

	fmt.Println("Original array:", solution.Reset())
	fmt.Println("Shuffled array:", solution.Shuffle())
	fmt.Println("Reset array:", solution.Reset())
	fmt.Println("Shuffled array again:", solution.Shuffle())

}
