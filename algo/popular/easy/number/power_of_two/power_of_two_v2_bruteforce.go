package power_of_two

func isPowerOfTwoV2(n int) bool {
	/*
		METHOD: Bruteforce
		Time complexity: O(log n) - последовательность на каждом шаге будет уменьшаться в 2 раза.
		Space complexity: O(1)
	*/
	if n <= 0 {
		return false
	}

	// Если остаток от деления равен 0, то продолжаем делить на 2
	for n%2 == 0 {
		n /= 2
	}

	// Если число является степенью двойки, то в конце концов, 2/2 = 1.
	// И на выходе у нас будет число 1. В противном случае вернем false.
	return n == 1
}
