package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	orig []int
	copy []int
}

func Constructor(nums []int) Solution {
	copyArr := make([]int, len(nums))
	copy(copyArr, nums)
	return Solution{
		orig: nums,
		copy: copyArr,
	}
}

func (this *Solution) Reset() []int {
	return this.orig
}

func (this *Solution) Shuffle() []int {
	rand.Shuffle(len(this.orig), func(i, j int) {
		this.copy[i], this.copy[j] = this.copy[j], this.copy[i]
	})

	return this.copy
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main() {
	nums := []int{1, 2, 3, 4, 5}
	solution := Constructor(nums)

	fmt.Println("Original array:", solution.Reset())
	fmt.Println("Shuffled array:", solution.Shuffle())
	fmt.Println("Reset array:", solution.Reset())
	fmt.Println("Shuffled array again:", solution.Shuffle())
}
