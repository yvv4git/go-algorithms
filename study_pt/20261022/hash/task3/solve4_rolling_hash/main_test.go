package main

import (
	"strings"
	"testing"
)

func Test_minBaseLengthV4(t *testing.T) {
	tests := []struct {
		name string
		t    string
		want int
	}{
		{"example 1", "zzz", 1},
		{"example 2", "bcabcab", 3},
		{"single character", "a", 1},
		{"all characters same", "aaaa", 1},
		{"no repetition", "abc", 3},
		{"repeated block", "ababab", 2},
		{"partial block", "abcab", 3},
		{"border without divisor", "aabaa", 3},
		{"longer partial block", "abcabcab", 3},
		{"empty string", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minBaseLengthV4(tt.t); got != tt.want {
				t.Errorf("minBaseLengthV4(%q) = %d, want %d", tt.t, got, tt.want)
			}
		})
	}
}

func Test_minBaseLengthV4_maxSize(t *testing.T) {
	tests := []struct {
		name string
		t    string
		want int
	}{
		{"period two", strings.Repeat("ab", 25000), 2},
		{"period one", strings.Repeat("a", 50000), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minBaseLengthV4(tt.t); got != tt.want {
				t.Errorf("minBaseLengthV4(max size) = %d, want %d", got, tt.want)
			}
		})
	}
}
