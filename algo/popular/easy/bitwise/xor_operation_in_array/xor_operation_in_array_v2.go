package xor_operation_in_array

func xorOperationV2(n int, start int) int {
	/*
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	result := start

	for i := 1; i < n; i++ {
		result = result ^ (start + 2*i)
	}

	return result
}
