package happy_number

func isHappyV1(n int) bool {
	/*
		Time complexity: O(1)
		Space complexity: O(1)
	*/
	for {
		summa := 0
		for n > 0 {
			d := n % 10
			//fmt.Printf("n=%d d=%d \n", n, d)
			summa += d * d
			n /= 10
		}

		if summa == 1 {
			return true
		}

		if summa == 4 {
			return false
		}
		n = summa
	}
}
