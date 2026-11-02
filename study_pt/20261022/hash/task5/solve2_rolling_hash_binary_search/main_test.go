package main

import (
	"strings"
	"testing"
)

func Test_maxSubstringV2(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"example 1", "abacaba", "caba"},
		{"example 2", "babb", "bb"},
		{"single character", "a", "a"},
		{"all equal", "aaaa", "aaaa"},
		{"increasing", "abc", "c"},
		{"decreasing", "cba", "cba"},
		{"banana", "banana", "nana"},
		{"alternating", "abab", "bab"},
		{"max at end", "az", "z"},
		{"max at start", "za", "za"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubstringV2(tt.s); got != tt.want {
				t.Fatalf("maxSubstringV2(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func Test_maxSubstringV2_maxSize(t *testing.T) {
	bigA := strings.Repeat("a", 999999)

	tests := []struct {
		name string
		s    string
		want string
	}{
		{"max char at end", bigA + "z", "z"},
		{"max char at start", "z" + bigA, "z" + bigA},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubstringV2(tt.s); got != tt.want {
				t.Fatalf("maxSubstringV2 of length %d = %q, want %q", len(tt.s), got, tt.want)
			}
		})
	}
}
