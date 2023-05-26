package bits_count

import "fmt"

func bitsCountV1(n int8) int {
	var result int

	for i := 0; i < 8; i++ {
		fmt.Printf("n: %d(%08b) ", n, n)
		if n&1 == 1 {
			result++
		}
		fmt.Printf("result: %d(%08b) \n", result, result)
		n >>= 1
	}

	return result
}
