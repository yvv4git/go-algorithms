package perfect_number

func checkPerfectNumberV2(num int) bool {
	/*
		Time complexity: O(logN)
		Space complexity: O(1)
	*/
	if num <= 1 {
		return false
	}

	sum := 0
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sum = sum + i + (num / i)
		}
	}

	return sum-num == num
}
