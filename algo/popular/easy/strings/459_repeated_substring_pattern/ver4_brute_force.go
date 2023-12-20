package _59_repeated_substring_pattern

import "strings"

func repeatedSubstringPatternV4(s string) bool {
	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			substring := s[:i]
			var builder strings.Builder
			for j := 0; j < n/i; j++ {
				builder.WriteString(substring)
			}
			if builder.String() == s {
				return true
			}
		}
	}
	return false
}
