package _630_arithmetic_subarrays

import "sort"

func checkArithmeticSubArraysV1(nums []int, l []int, r []int) []bool {
	/*
		METHOD: ???
		TIME COMPLEXITY: O(mn log(n))
		SPACE COMPLEXITY: O(n)
	*/
	resul := make([]bool, len(l))

	for i := 0; i < len(l); i++ {
		resul[i] = true
		size := r[i] - l[i] + 1
		if size < 2 {
			resul[i] = false
			continue
		}

		arr := make([]int, size)
		copy(arr, nums[l[i]:r[i]+1])
		sort.Ints(arr)
		diff := arr[1] - arr[0]

		for j := 2; j < len(arr); j++ {
			if arr[j]-arr[j-1] != diff {
				resul[i] = false
				break
			}
		}
	}

	return resul
}
