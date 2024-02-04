package sqrt

func mySqrtV1(x int) int {
	/*
		METHOD: Bruteforce.
		Time complexity : O(n)
		Space complexity : O(1)
	*/
	var out = 0

	for out*out < x {
		out++
	}

	if out*out > x {
		out--
	}

	return out
}
