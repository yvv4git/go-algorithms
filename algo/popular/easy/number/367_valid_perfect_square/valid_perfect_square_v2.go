package _67_valid_perfect_square

func isPerfectSquareV2(num int) bool {

	i := 1
	for i*i < num {
		i++
	}

	return i*i == num
}
