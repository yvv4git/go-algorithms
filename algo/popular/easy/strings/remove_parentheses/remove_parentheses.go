package remove_parentheses

import "strings"

/*
Надо подчистить скобки в строке.
Чтобы остались только открывающие и закрывающие.

For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings
*/
func removeOuterParentheses(s string) string {
	state := 0
	var sb strings.Builder

	for _, v := range s {
		if v == '(' {
			state++
			if state == 1 {
				continue
			}
		}

		if v == ')' {
			state--
			if state == 0 {
				continue
			}
		}
		sb.WriteRune(v)
	}

	return sb.String()
}
