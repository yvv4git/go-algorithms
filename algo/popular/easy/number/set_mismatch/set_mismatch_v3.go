package set_mismatch

func findErrorNums(nums []int) []int {
	/*
		Method: XOR
		Time complexity: O(n)
		Space complexity: O(n)

		Explanation: XOR of integer can cancel out each other. we can find the missing integer by using xor
		e.g: (1^2^3^5) ^ (1^2^3^4^5) = 4 (4 is the missing integer in the set)
	*/
	arr := make([]int, len(nums)+1)

	result := make([]int, 0, 2)
	var xor int
	var xorComp int
	for i, num := range nums {
		if i == 0 {
			xorComp = i + 1
		} else {
			xorComp = xorComp ^ (i + 1)
		}

		if arr[num] == 0 {
			arr[num] = num
			if i == 0 {
				xor = num
			} else {
				xor = xor ^ num
			}
		} else {
			result = append(result, num)
		}
	}

	result = append(result, xorComp^xor)

	return result
}
