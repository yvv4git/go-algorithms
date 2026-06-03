package main

import (
	"testing"
)

func TestZFunction(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{"single", "a", []int{0}},
		{"zzz", "zzz", []int{0, 2, 1}},
		{"bcabcab", "bcabcab", []int{0, 0, 0, 4, 0, 0, 1}},
		{"abcabc", "abcabc", []int{0, 0, 0, 3, 0, 0}},
		{"aaaa", "aaaa", []int{0, 3, 2, 1}},
		{"ababa", "ababa", []int{0, 0, 3, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zFunction(tt.s)
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

func TestPeriodZ(t *testing.T) {
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
			got := periodZ(tt.s)
			if got != tt.want {
				t.Errorf("periodZ(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
