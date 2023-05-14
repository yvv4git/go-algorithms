package happy_number

func isHappyV3(n int) bool {
	/*
		Time complexity : O(1)
		Space complexity : O(1)
	*/
	sum := 0

	if n < 10 {
		// If 7, the number will also be happy:
		// 7^2 = 49
		// 4^2 + 9^2 = 97
		// 9^2 + 7^2 = 130
		// 2^2 + 3^2 + 0^2 = 13
		// 1^2 + 3^2 = 10
		// 1^2 + 0^2 = 1 <- Happy!
		return n == 1 || n == 7
	}

	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	if sum == 1 {
		return true
	}

	return isHappyV3(sum)
}
