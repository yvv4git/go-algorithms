package strings

import "testing"

/*
# 301. Remove Invalid Parentheses
# HARD
Given a string s that contains parentheses and letters, remove the minimum number of invalid parentheses to make the input string valid.
Return a list of unique strings that are valid with the minimum number of removals. You may return the answer in any order.
*/

func removeInvalidParentheses(s string) []string {
	N := len(s)
	ans := make([]string, 0)
	if N == 0 {
		return nil
	}

	reverse := func(s string) string {
		r := []rune(s)
		n := len(r)
		for i := 0; i < n/2; i++ {
			r[i], r[n-i-1] = r[n-i-1], r[i]
		}

		return string(r)
	}

	// find first parenthesis close before open
	findFirstInvalid := func(s string, from int, in, out byte) int {
		n := len(s)
		stack := 0
		for i := from; i < n; i++ {
			if s[i] == in {
				stack++
			} else if s[i] == out {
				stack--
			}
			if stack < 0 {
				return i
			}
		}
		return -1
	}

	var remove func(s string, scanFrom, delFrom int, in, out byte)
	remove = func(s string, scanFrom, delFrom int, in, out byte) {
		invalidAt := findFirstInvalid(s, scanFrom, in, out)

		if invalidAt >= 0 {
			// remove every removeable parenthesis expect continuous ones
			for i := delFrom; i <= invalidAt; i++ {
				if s[i] == out && (i == delFrom || s[i-1] != out) {
					remove(s[:i]+s[i+1:], invalidAt, i, in, out)
				}
			}
		} else {
			if in == '(' {
				// reverse and check again
				remove(reverse(s), 0, 0, out, in)
			} else {
				ans = append(ans, reverse(s))
			}
		}
	}

	remove(s, 0, 0, '(', ')')

	return ans
}

func TestRemoveInvalidParentheses(t *testing.T) {
	tests := []struct {
		name  string
		inStr string
		want  []string
	}{
		{
			name:  "CASE-1",
			inStr: "()())()",
			want:  []string{"(a())()", "(a)()()"},
		},
		{
			name:  "CASE-2",
			inStr: "(a)())()",
			want:  []string{"(a())()", "(a)()()"},
		},
		{
			name:  "CASE-3",
			inStr: ")(",
			want:  []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeInvalidParentheses(tt.inStr)
			t.Logf("%v", result)
		})
	}
}
