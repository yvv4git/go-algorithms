package main

import (
	"strings"
	"testing"
)

func Test_countDistinctSubstringsV3(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"example 1", "abc", 6},
		{"example 2", "aaa", 3},
		{"single character", "a", 1},
		{"two same", "aa", 2},
		{"two different", "ab", 3},
		{"periodic", "abab", 7},
		{"all distinct", "abcd", 10},
		{"empty string", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDistinctSubstringsV3(tt.s); got != tt.want {
				t.Errorf("countDistinctSubstringsV3(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func Test_countDistinctSubstringsV3_maxSize(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"period two", strings.Repeat("ab", 512), 2047},
		{"alphabet pattern", strings.Repeat("abcdefghijklmnopqrstuvwxyz", 39) + "abcdefghij", 26299},
		{"pseudo random", pseudoRandomString(1024), 522968},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDistinctSubstringsV3(tt.s); got != tt.want {
				t.Errorf("countDistinctSubstringsV3(max size) = %d, want %d", got, tt.want)
			}
		})
	}
}

func pseudoRandomString(n int) string {
	b := make([]byte, n)
	x := uint32(123456789)
	for i := range b {
		x = x*1664525 + 1013904223
		b[i] = byte('a' + x%26)
	}
	return string(b)
}
