package main

import (
	"math"
	"strconv"
	"unicode"
)

func myAtoiV5(s string) int {
	/*
		METHOD: Iteration + lib
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1) - в общем случае, в худшем случае, если строка состоит только из цифр, то n записей придется копировать.
	*/
	result := ""
outerLoop:
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			result += string(s[i])
		} else {
			if result != "" {
				break
			}
			switch string(s[i]) {
			case "+":
				result += string(s[i])
				break
			case "-":
				result += string(s[i])
				break
			case " ":
				break
			default:
				break outerLoop
			}
		}
	}
	i, _ := strconv.Atoi(result)
	if i > math.MaxInt32 {
		return math.MaxInt32
	} else if i < math.MinInt32 {
		return math.MinInt32
	}
	return i
}
