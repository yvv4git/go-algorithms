package missingnumber

func missingNumberV2(nums []int) int {
	/*
		METHOD: Gauss sum
		Time complexity: O(n)
		Worst time complexity: O(n)
		Space complexity: O(1)
	*/
	// Initialize with total sum from 0 to n
	missingNumber := (len(nums) * (len(nums) + 1)) / 2

	for _, number := range nums {
		missingNumber -= number
	}

	return missingNumber
}
