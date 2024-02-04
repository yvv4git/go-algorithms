package ver1_bruteforce

import "strconv"

func primePalindromeV1(N int) int {
	/*
		METHOD: Bruteforce
		TIME COMPLEXITY: O(n sqrt(n)), это связано с двумя циклами: внешним, который проверяет каждое число на простоту и палиндромность, и внутренним, который проверяет каждую цифру в числе на палиндромность.
		SPACE COMPLEXITY: O(1)
	*/
	if N <= 2 {
		return 2
	}

	// Если N четное и больше 2, то простое число не может быть палиндромом
	if N%2 == 0 {
		N++
	}
	for ; ; N += 2 {
		if isPalindrome(N) && isPrime(N) {
			return N
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
