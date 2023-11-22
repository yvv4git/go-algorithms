package _3_roman_to_Integer

func romanToIntV1(s string) int {
	/*
		Method: Use map.
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	var romanMap = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var result = romanMap[s[len(s)-1]]

	for i := len(s) - 2; i >= 0; i-- {
		if romanMap[s[i]] < romanMap[s[i+1]] {
			result -= romanMap[s[i]]
		} else {
			result += romanMap[s[i]]
		}
	}

	return result
}
