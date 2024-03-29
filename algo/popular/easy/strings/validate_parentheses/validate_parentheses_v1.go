package main

import "fmt"

func isValidV1(s string) bool {
	/*
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	var stack []rune
	var checkParameter = map[rune]rune{')': '(', '}': '{', ']': '['}

	fmt.Println(s)
	for _, val := range s {
		if val == '(' || val == '{' || val == '[' {
			stack = append(stack, val)
		} else if len(stack) > 0 && checkParameter[val] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
