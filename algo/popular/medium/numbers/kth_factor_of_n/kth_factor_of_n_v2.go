package kth_factor_of_n

func kthFactorV2(n int, k int) int {
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
