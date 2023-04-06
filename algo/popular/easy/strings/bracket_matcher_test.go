package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
## Task:
Bracket Matcher
Have the function BracketMatcher(str) take the str parameter being passed and return 1
if the brackets are correctly matched and each one is accounted for.
Otherwise, return 0. For example: if str is "(hello (world))", then the output should be 1, but if str is "((hello (world))"
the the output should be 0 because the brackets do not correctly match up. Only "(" and ")" will be used as brackets.
If str contains no brackets return 1.

## Example-1
Input: "(coder)(byte))"
Output: 0

## Example-2
Input: "(c(oder)) b(yte)"
Output: 1
*/

func BracketMatcher(str string) string {
	// code goes here
	var leftBrackets, rightBrackets int

	for idx := range str {
		if str[idx] == '(' {
			leftBrackets++
		} else if str[idx] == '}' {
			rightBrackets++
		}
	}

	if leftBrackets == rightBrackets {
		return "1"
	}

	return "0"
}

func TestBracketMatcher(t *testing.T) {
	tests := []struct {
		name     string
		inputStr string
		want     string
	}{
		{
			name:     "CASE-1",
			inputStr: "(coder)(byte))",
			want:     "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BracketMatcher(tt.inputStr)
			assert.Equal(t, tt.want, result)
		})
	}
}
