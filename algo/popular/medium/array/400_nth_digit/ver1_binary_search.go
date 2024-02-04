package _00_nth_digit

// findNthDigitV1 находит n-ю цифру в последовательности 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...
func findNthDigitV1(n int) int {
	/*
		METHOD: Binary search
		TIME COMPLEXITY: O(log n), так как мы делим пространство поиска пополам в цикле
		SPACE COMPLEXITY: O(1), так как мы используем небольшое количество переменных, которые не зависят от размера входных данных
	*/

	if n <= 9 {
		return n
	}
	res := 0

	for lo, hi := 9, n+1; lo < hi; {
		mid := lo + ((hi - lo) >> 1)
		digits, digitsBefore := helper(mid)

		if n > digitsBefore+len(digits) {
			lo = mid + 1
		} else if n <= digitsBefore {
			hi = mid
		} else {
			digitsBefore++

			index := 0

			for digitsBefore < n {
				digitsBefore++
				index++
			}

			res = digits[index]
			break
		}
	}

	return res
}

func helper(n int) ([]int, int) {
	digits := []int{}
	digitsBefore := 0

	for x := n; x > 0; x /= 10 {
		digits = append(digits, x%10)
	}

	reverse(digits)
	shift, sub := 9, 1

	for i := 1; i < len(digits); i++ {
		digitsBefore += (i * shift)
		shift *= 10
		sub *= 10
	}

	digitsBefore += (n - sub) * len(digits)

	return digits, digitsBefore
}

func reverse(values []int) {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
}
