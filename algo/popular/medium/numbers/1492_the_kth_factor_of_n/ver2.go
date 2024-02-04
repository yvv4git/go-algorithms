package _492_the_kth_factor_of_n

func kthFactorV2(n int, k int) int {
	/*
		METHOD: Math
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	var counter int

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			counter++
		}

		if counter == k {
			return i
		}
	}

	return -1
}
