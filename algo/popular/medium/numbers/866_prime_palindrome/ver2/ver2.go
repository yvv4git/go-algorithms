package ver2

import "strconv"

func isPrime(n int) bool {
	if n < 2 || n%2 == 0 {
		return n == 2
	}

	for x := 3; x*x <= n; x += 2 {
		if n%x == 0 {
			return false
		}
	}

	return true
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func primePalindrome(n int) int {
	if 8 <= n && n <= 11 {
		return 11
	}

	for x := 1; x < 100000; x++ {
		s := strconv.Itoa(x)
		r := reverseString(s)
		y, _ := strconv.Atoi(s + r[1:])
		if y >= n && isPrime(y) {
			return y
		}
	}

	return -1
}
