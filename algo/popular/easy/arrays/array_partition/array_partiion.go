package array_partition

func arrayPairSum(nums []int) int {
	/*
		Time complexity : O(n) - no nested loops
		Space complexity : O(min+max), where min is minimum value in array, max is maximum value in array.

		Steps:
		- Sort array
		- Iterate over even index in array and add the value to the sum

		Algorithm:
		- Iterate over even positions of array
		- Add number on even index to the sum
		- Until we summed n integers in total or we reached the end of an array
	*/
	sortedArr := sortCounting(nums)

	sum := 0
	count := 0
	n := len(nums)

	for i := 0; count < n/2 && i < n; i += 2 {
		sum += sortedArr[i]
		count++
	}

	return sum
}

func sortCounting(nums []int) []int {
	min, max := getMinMax(nums)

	offset := 0
	if min > 0 {
		min = 0
	} else {
		offset = -min
	}

	memory := make([]int, max-min+1)
	for _, number := range nums {
		memory[number+offset] += 1 // add offset in case the number is negative
	}

	n := len(nums)
	sorted := make([]int, 0, n)
	for number, occurrences := range memory {
		for i := occurrences; i > 0; i-- {
			sorted = append(sorted, number-offset) // subtract offset to form array of original values
		}
	}

	return sorted
}

func getMinMax(nums []int) (int, int) {
	min := nums[0]
	max := nums[0]
	for _, number := range nums {
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
	}
	return min, max
}
