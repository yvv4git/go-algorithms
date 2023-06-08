package multition_karatsuba

import "math"

func getDecimalDigits(num int64) uint {
	var result uint

	if num == 0 {
		return 1
	}

	if num < 0 {
		num = -num
	}
	for num > 0 {
		result++
		num = num / 10
	}

	return result
}

func getHighAndLowDigits(num int64, digits uint) (int64, int64) {
	divisor := int64(math.Pow(10, float64(digits)))

	if num >= divisor {
		return num / divisor, num % divisor
	} else {
		return 0, num
	}
}

func karatsuba(x int64, y int64) int64 {
	var maxDigits uint
	positive := true

	if x == 0 || y == 0 {
		return 0
	}

	if (x > 0 && y < 0) || (x < 0 && y > 0) {
		positive = false
	}

	if x < 0 {
		x = -x
	}

	if y < 0 {
		y = -y
	}

	if x < 10 || y < 10 {
		return x * y
	}

	xDigits := getDecimalDigits(x)
	yDigits := getDecimalDigits(y)

	if xDigits >= yDigits {
		maxDigits = xDigits / 2
	} else {
		maxDigits = yDigits / 2
	}

	xHigh, xLow := getHighAndLowDigits(x, maxDigits)
	yHigh, yLow := getHighAndLowDigits(y, maxDigits)

	z0 := karatsuba(xLow, yLow)
	z1 := karatsuba(xLow+xHigh, yLow+yHigh)
	z2 := karatsuba(xHigh, yHigh)

	if positive {
		return (z2 * int64(math.Pow(10, float64(2*maxDigits)))) + (z1-z2-z0)*int64(math.Pow(10, float64(maxDigits))) + z0
	} else {
		return -((z2 * int64(math.Pow(10, float64(2*maxDigits)))) + (z1-z2-z0)*int64(math.Pow(10, float64(maxDigits))) + z0)
	}

}
