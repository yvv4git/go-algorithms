package number_of_even_and_odd

func evenOddBitV1(n int) []int {
	/*
		METHOD: Hash
		Time complexity: O(num), where num is count of bits
		Space complexity: O(1)
	*/
	res := make([]int, 2, 2)

	for i := 0; n > 0; i++ {
		v := n % 2
		n /= 2

		if v != 1 {
			continue
		}

		if i%2 == 0 {
			res[0]++
		} else {
			res[1]++
		}
	}

	return res
}
