package placements

// Placements
// Размещение - это когда у нас есть n элементов, но из них мы будем
// отбирать m элементов и переставлять всеми возможными способами.
func Placements(n int, m int) int {
	return factorialRecursion(n) / factorialRecursion(n-m)
}

// factorialRecursion is used for calculate factorial with recursion.
func factorialRecursion(n int) int {
	var result int
	if n == 0 || n == 1 {
		return 1
	}
	result = factorialRecursion(n-1) * n
	return result
}
