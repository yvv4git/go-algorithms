package _492_the_kth_factor_of_n

func kthFactorV1(n int, k int) int {
	/*
		METHOD: Math
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)
	*/
	for i := 1; i <= n/2+1; i++ {
		if n%i == 0 {
			k--
		}
		if k == 0 {
			return i
		}
	}

	if k == 1 {
		return n
	}

	return -1
}
