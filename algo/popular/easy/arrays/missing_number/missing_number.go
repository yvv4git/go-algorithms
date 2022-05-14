package missingnumber

func missingNumber(nums []int) int {
	var sum int
	var ind int

	for i, v := range nums {
		sum += v
		ind += i + 1
	}

	return ind - sum
}
