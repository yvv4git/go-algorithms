package _56_132_pattern

import "math"

func find132patternV1(nums []int) bool {
	/*
		Method: Stack
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность этого алгоритма - O(n), где n - количество элементов в массиве nums,
		так как мы проходим по массиву два раза: один раз слева направо, а второй - справа налево.

		Space complexity
		Пространственная сложность - O(n), так как в худшем случае мы можем хранить все элементы nums в стеке.
	*/
	stack := make([]int, 0, len(nums))
	ak := math.MinInt64

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < ak {
			return true
		}

		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			ak = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, nums[i])
	}

	return false
}
