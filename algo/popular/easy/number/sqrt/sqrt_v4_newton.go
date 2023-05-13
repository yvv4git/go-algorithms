package sqrt

var p = 0.0

func mySqrtV4(x int) int {
	z := 1.0

	for p-z != 0 {
		z = newton(z, float64(x))
		p = newton(z, float64(x))
	}

	return int(z)
}

func newton(z, x float64) float64 {
	return z - (((z * z) - x) / (2 * z))
}
