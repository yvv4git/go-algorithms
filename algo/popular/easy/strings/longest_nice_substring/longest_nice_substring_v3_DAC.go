package longest_nice_substring

import "fmt"

func longestNiceSubstringV3(s string) string {
	/*
		METHOD: Divide And Conquer
		Time complexity: ???
		Space complexity: ???
	*/
	return string(helper([]byte(s)))
}

func helper(b []byte) []byte {
	fmt.Printf("b:%#v %v\n", b, string(b))
	if len(b) < 2 {
		return []byte{}
	}

	d := make(map[byte]bool)

	for _, v := range b {
		d[v] = true
	}

	for k, v := range b {
		if v > 96 && d[v-32] {
			continue
		}

		if v < 91 && d[v+32] {
			continue
		}

		sub1, sub2 := helper(b[:k]), helper(b[k+1:])

		if len(sub1) >= len(sub2) {
			return sub1
		}

		return sub2
	}

	return b
}
