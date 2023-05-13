package sqrt

var p = 0.0

func sqrtV1(x float64) float64 {
	z := 1.0

	for p-z != 0 {
		z = newton(z, x)
		p = newton(z, x)
	}

	return z
}

func newton(z, x float64) float64 {
	return z - (((z * z) - x) / (2 * z))
}
