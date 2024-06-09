package _630_arithmetic_subarrays

import "sort"

func checkArithmeticSubArraysV4(nums []int, l []int, r []int) []bool {

	// we need to clone
	// otherwise sorting slice affects the main array
	isArithmetic := func(arr []int) bool {
		la := len(arr)
		if la < 2 {
			return true
		}

		// clone slice and then sort the clone
		arrC := make([]int, la)
		copy(arrC, arr)
		sort.Ints(arrC)

		// check each subsequent difference
		isA := true
		d_0 := arrC[1] - arrC[0]
		for i := 2; i < la; i++ {
			d_i := arrC[i] - arrC[i-1]
			if d_0 != d_i {
				isA = false
				break
			}
		}
		return isA
	}

	ans := []bool{}
	for i, li := range l {
		ri := r[i]
		// fmt.Println(li, ri, nums[li:ri+1], isArithmetic(nums[li: ri+1]))
		ans = append(ans, isArithmetic(nums[li:ri+1]))
	}

	return ans
}
