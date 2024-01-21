package main

func maximumOddBinaryNumberV2(s string) string {
	/*
		Method: Iteration
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	count := 0
	for _, r := range s {
		if r == '1' {
			count++
		}
	}

	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 || count > 1 {
			res = append(res, '1')
			count--
		} else {
			res = append(res, '0')
		}
	}

	return string(res)
}
