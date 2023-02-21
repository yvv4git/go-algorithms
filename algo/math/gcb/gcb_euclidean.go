package gcb

// GCD - the greatest common divisor via Euclidean algorithm.
func gcbEuclidean(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
