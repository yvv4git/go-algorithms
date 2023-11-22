package _3_roman_to_Integer

var dict = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func romanToIntV3(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			result += dict[s[i]]
			return result
		}

		if dict[s[i]] >= dict[s[i+1]] {
			result += dict[s[i]]
		} else {
			result += dict[s[i+1]] - dict[s[i]]
			i++
		}
	}

	return result
}
