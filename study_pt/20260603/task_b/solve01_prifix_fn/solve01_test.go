package main

import (
	"testing"
)

func TestPrefixFunction(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{"single", "a", []int{0}},
		{"two same", "aa", []int{0, 1}},
		{"two diff", "ab", []int{0, 0}},
		{"three same", "aaa", []int{0, 1, 2}},
		{"bcabcab", "bcabcab", []int{0, 0, 0, 1, 2, 3, 4}},
		{"abcabc", "abcabc", []int{0, 0, 0, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixFunction(tt.s)
			if len(got) != len(tt.want) {
				t.Fatalf("len=%d, want len=%d", len(got), len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("at %d: got %d, want %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestPeriod(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"single char", "a", 1},
		{"two same", "aa", 1},
		{"zzz", "zzz", 1},
		{"ab", "ab", 2},
		{"abcabc", "abcabc", 3},
		{"bcabcab", "bcabcab", 3},
		{"ababab", "ababab", 2},
		{"ababa", "ababa", 2},
		{"abcdef", "abcdef", 6},
		{"aaaa", "aaaa", 1},
		{"aabaaab", "aabaaab", 4},
		{"abcdeabcde", "abcdeabcde", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pi := prefixFunction(tt.s)
			period := len(tt.s) - pi[len(tt.s)-1]
			if period != tt.want {
				t.Errorf("period(%q) = %d, want %d", tt.s, period, tt.want)
			}
		})
	}
}
