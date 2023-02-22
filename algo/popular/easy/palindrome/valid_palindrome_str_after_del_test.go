package main

import "testing"

/*
Problem:
Given a string s, return true if the s can be palindrome after deleting at most one character from it.

Constraints:
1 <= s.length <= 10^5
s consists of lowercase English letters.


Example 1:
Input: s = "aba"
Output: true

Example 3:
Input: s = "abc"
Output: false
*/

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func validPalindrome(s string) bool {
	valid, l, r := isPalindrome(s, 0, len(s)-1)
	if valid {
		return true
	}

	if valid, _, _ := isPalindrome(s, l+1, r); valid {
		return true
	}

	if valid, _, _ := isPalindrome(s, l, r-1); valid {
		return true
	}

	return false
}

func isPalindrome(s string, l, r int) (bool, int, int) {
	for l < r {
		if s[l] != s[r] {
			return false, l, r
		}
		l++
		r--
	}
	return true, 0, 0
}
