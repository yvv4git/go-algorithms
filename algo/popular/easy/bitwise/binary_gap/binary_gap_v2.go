package binary_gap

func BinaryGapV2(N int) int {
	/*
		METHOD: Bitwise
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)

		Explanation: We check each bit of number from right to left.
		After finding a "1" bit, start to count and continue to find next "1" and so on
	*/
	var res int = 0
	var right int64 = 1
	var n int64 = int64(N)
	for right <= n {
		for right <= n && (right|n) != n {
			right <<= 1
		}
		if right > n {
			break
		}
		var dis int = 1
		var left int64 = right << 1
		for left <= n && (left|n) != n {
			left <<= 1
			dis++
		}
		if left > n {
			break
		}
		if dis > res {
			res = dis
		}
		right = left
	}
	return res
}
