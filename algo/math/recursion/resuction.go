package recursion

func rec(n int) {
	println(n)

	if n < 1 {
		return // base case - база рекурсии
	}

	rec(n - 1) // recursion case - рекурсивный случай
}
