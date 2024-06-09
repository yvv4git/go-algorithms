package _66_prime_palindrome

import "math"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPalindrome(n int) bool {
	reversed := 0
	original := n

	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}

	return original == reversed
}
