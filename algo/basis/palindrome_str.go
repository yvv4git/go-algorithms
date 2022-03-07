package basis

// Сложность алгоритма O(n), ибо приходится
// линейно идти по массиву/строке.
func isPalindromeStr(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}

	return true
}

func isPalindromeStrOptimise(str string) bool {
	bound := (len(str) / 2) + 1
	for i := 0; i < bound; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
