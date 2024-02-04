package _630_arithmetic_subarrays

func checkArithmeticSubArraysV3(nums []int, l []int, r []int) []bool {
	/*
		METHOD: ???
		TIME COMPLEXITY: O(mn log(n))
		Space complexity: O(n)
	*/
	result := make([]bool, len(l))
	for i := range l {
		result[i] = isArithmeticSeq(nums[l[i] : r[i]+1])
	}

	return result
}

func isArithmeticSeq(nums []int) bool {
	n := len(nums)

	max, min := nums[0], nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	// all equal?
	if max == min {
		return true
	}

	// diff between 2 number
	d := (max - min) / (n - 1)

	// check if (max-min) % (n-1) == 0
	if (max - min) != d*(n-1) {
		return false
	}

	// We expect this slice to be arith, thus: min min+d min+2d ... max
	//
	// In another word: min min+d ... must exist inside nums.
	//
	// So, let's build a set and check each number exists or not.
	m := make(map[int]struct{}) // m is hashset
	for _, v := range nums {
		m[v] = struct{}{}
	}

	// check each number, we only need to check n - 1 times
	for n--; n > 0; n-- {
		min += d // each time -> increase by d
		if _, ok := m[min]; !ok {
			return false
		}
	}

	return true
}
