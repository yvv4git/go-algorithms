package rune_brackets

import "github.com/yvv4git/go-algorithms/ds/stack/stack/on_list/v2_container_list"

/*
Given string s.
s contains only '[', ']', '(', ')', '{', '}'

Examples:
"[](){}" -> true
"[({})]" -> true
"[(])" -> false
"}{" -> false
"[](){[({})]}[][({})](){}[[(])](){}[]([({})]){}[](){}"
*/

// Validate - used for validate string of brackets.
func Validate(s string) bool {
	checkMap := map[rune]rune{
		'[': ']',
		//']': '[',
		'(': ')',
		//')': '(',
		'{': '}',
		//'}': '{',
	}

	stack := v2_container_list.NewStack()

	var lastSymbol rune
	for _, symbol := range s {
		lastSymbol = stack.Front()
		if checkMap[lastSymbol] == symbol {
			stack.Pop()

			continue
		}

		stack.Push(symbol)
	}

	if stack.Size() == 0 {
		return true
	}

	return false
}
