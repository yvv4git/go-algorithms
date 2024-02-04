package missingnumber

func missingNumberV3(nums []int) int {
	/*
		METHOD: Pattern xor
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	a, b := 0, 0

	for i := 0; i <= len(nums); i++ {
		a = a ^ i
	}

	for _, v := range nums {
		b = b ^ v
	}

	return a ^ b
}
