package perfect_number

func checkPerfectNumberV1(num int) bool {
	/*
		Time complexity: ???
		Space complexity: O(1)
	*/
	if num == 1 {
		return false
	}

	sum := 1
	left := 2
	for {
		right := num / left
		if left >= right {
			break
		}
		if num%left == 0 {
			sum = sum + left + right
		}
		left++
	}

	if sum == num {
		return true
	}

	return false
}
