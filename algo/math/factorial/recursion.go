package factorial

// FactorialByRecursion - used recursion for calc factorial
func FactorialByRecursion(n int) int {
	var result int
	if n == 0 || n == 1 {
		return 1
	}
	result = FactorialByRecursion(n-1) * n
	return result
}
