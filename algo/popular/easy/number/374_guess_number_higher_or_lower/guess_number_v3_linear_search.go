package _74_guess_number_higher_or_lower

func linearSearch(n int) int {
	for i := 1; i <= n; i++ {
		if guess(i) == 0 {
			return i
		}
	}
	return -1
}
