package _217_find_palindrome_with_fixed_length

func kthPalindromeV2(queries []int, intLength int) []int64 {
	/*
		METHOD: DP + math
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)

		Time complexity
		Временная сложность функции kthPalindrome составляет O(n), где n - количество запросов.
		Это происходит потому, что функция проходит по каждому запросу и выполняет некоторые операции.

		Space complexity
		Пространственная сложность функции kthPalindrome составляет O(n), где n - количество запросов.
		Это происходит потому, что функция создает новый массив ans, размер которого равен количеству запросов.
	*/
	half := (intLength-1)/2 + 1

	start := 1
	for i := 1; i < half; i++ {
		start *= 10
	}

	numPalin := start * 9
	start--
	doubleLast := (intLength & 1) == 0

	ans := make([]int64, len(queries))
	for i := range queries {
		if queries[i] > numPalin {
			ans[i] = -1
		} else {
			ans[i] = mirror(start+queries[i], doubleLast)
		}
	}

	return ans
}

func mirror(v int, doubleLast bool) int64 {
	/*
		TIME COMPLEXITY: O(log(v)
		Space complexity: O(1)

		Time complexity
		Временная сложность функции mirror составляет O(log(v)), где v - значение палиндрома.
		Это происходит потому, что функция проходит по каждой цифре числа v и выполняет некоторые операции.
		Количество цифр в числе v равно log10(v) + 1.
		Это связано с тем, что в десятичной системе счисления количество цифр в числе равно log10(v) + 1, где log10 - это логарифм по основанию 10.
		Таким образом, временная сложность функции mirror будет O(log(v)), где v - это входное значение.

		Space complexity
		Пространственная сложность функции mirror составляет O(1), так как функция не использует дополнительное пространство,
		которое зависит от значения палиндрома v.
	*/
	ans := 0
	p := 1 // power of 10
	for v > 0 {
		d := v % 10 // digit
		v /= 10

		if p == 1 { // first
			if doubleLast {
				ans = d * 11
				p = 100
			} else {
				ans = d
				p = 10
			}
		} else {
			ans = 10*(d*p+ans) + d
			p *= 100
		}
	}

	return int64(ans)
}
