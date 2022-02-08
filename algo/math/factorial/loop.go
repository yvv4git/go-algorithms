package factorial

// FactorialByLoop - calculate factorial without recursion
func FactorialByLoop(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
	}

	return result
}
