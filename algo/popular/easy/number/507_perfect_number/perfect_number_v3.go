package _07_perfect_number

func checkPerfectNumberV3(num int) bool {
	/*
		TIME COMPLEXITY: O(1)
		Space complexity: O(1)

		Объяснение: в задаче стоит ограничение, что максимальное число 10^8 = 100 000 000.
		Совершенные числа начинаются с 6 и их не так уж много, следовательно, мы можем запомнить заранее эти числа, которые
		существуют в рамках нашего ограничения, а затем вытаскивать их моментально из хранилища.
	*/
	perfectNumbers := []int{6, 28, 496, 8128, 33550336, 8589869056}

	for _, perfectNumber := range perfectNumbers {
		if num == perfectNumber {
			return true
		}
	}

	return false
}
