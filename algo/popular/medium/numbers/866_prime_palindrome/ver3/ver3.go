package ver3

import "math"

func primePalindrome(n int) int {
	/*
		METHOD: Iteration
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность этой функции зависит от размера входного числа n. В худшем случае,
		когда n является простым числом-палиндромом, функция может пройти через все простые числа-палиндромы,
		пока не найдет число, которое больше или равно n. В этом случае временная сложность будет O(n),
		где n - это входное число.

		Space complexity
		Пространственная сложность этой функции также зависит от входного числа n. В худшем случае,
		когда n является простым числом-палиндромом, функция может хранить все простые числа-палиндромы,
		пока не найдет число, которое больше или равно n.
		В этом случае пространственная сложность будет O(n), где n - это входное число.
	*/
	if n <= 2 {
		return 2
	}

	for len := digits(n); ; len++ {
		for root := start(len); ; root++ {
			s := build(root, len)
			if s >= n && isPrime(s) && isPalindrome(s) {
				return s
			}
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
	return n == reverse(n)
}

func digits(n int) int {
	if n == 0 {
		return 1
	}

	res := 0
	for n > 0 {
		res++
		n /= 10
	}

	return res
}

func start(len int) int {
	if len%2 == 0 {
		return int(math.Pow10(len/2 - 1))
	}

	return int(math.Pow10((len - 1) / 2))
}

func build(root, len int) int {
	if len%2 == 0 {
		return root*int(math.Pow10(len/2)) + reverse(root)
	}

	return root*int(math.Pow10((len-1)/2)) + reverse(root)
}

func reverse(n int) int {
	res := 0
	for n > 0 {
		res = res*10 + n%10
		n /= 10
	}

	return res
}
