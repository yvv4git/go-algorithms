package basis

import "strconv"

func isPalindromeInt32(x int32) bool {
	var reverse, remainder, tmp int32
	reverse = 0

	if x < 0 {
		return false
	}

	tmp = x

	for {
		remainder = x % 10
		reverse = reverse*10 + remainder
		x /= 10

		if x == 0 {
			break
		}
	}

	return tmp == reverse
}

// isPalindromeNumber несколько оптимизирован за счет.
func isPalindromeNumber(n int64) bool {
	if n < 0 {
		n = -n
	}

	s := strconv.FormatInt(n, 10)
	bound := (len(s) / 2) + 1
	for i := 0; i < bound; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
