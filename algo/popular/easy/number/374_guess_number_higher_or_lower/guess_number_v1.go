package _74_guess_number_higher_or_lower

var pick = 6

func guessNumberV1(n int) int {
	/*
		Method: Binary search.
		Time Complexity - O(log(n))
		Space Complexity - O(1)
	*/
	for i := 1; i <= n; i++ {
		if guess(i) == 0 {
			return i
		}
	}
	return -1
}

// Функция guess должна быть реализована в зависимости от вашей задачи.
// Это просто заглушка для примера.
func guess(num int) int {
	if num > pick {
		return -1
	} else if num < pick {
		return 1
	} else {
		return 0
	}
}
