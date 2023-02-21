package smallest_divisor

func SmallestDivisor(n int) int {
	result := n // Допустим, наименьший делитель есть само число.

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			if i <= result {
				result = i
			}
		}
	}

	return result
}
