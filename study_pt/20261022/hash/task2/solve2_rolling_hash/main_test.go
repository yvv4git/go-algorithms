package main

import (
	"strings"
	"testing"
)

func Test_minCyclicShiftV2(t *testing.T) {
	tests := []struct {
		name string
		x    string
		y    string
		want int
	}{
		{"example 1", "zabcd", "abcdz", 4},
		{"same string", "abc", "abc", 0},
		{"shift by one", "abcde", "eabcd", 1},
		{"shift by two", "abcde", "deabc", 2},
		{"not a shift", "abc", "acb", -1},
		{"single character same", "a", "a", 0},
		{"single character different", "a", "b", -1},
		{"periodic shift", "abab", "baba", 1},
		{"all characters same", "aaaa", "aaaa", 0},
		{"different lengths", "abc", "ab", -1},
		{"empty strings", "", "", 0},
		{"empty and non-empty", "", "a", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCyclicShiftV2(tt.x, tt.y); got != tt.want {
				t.Errorf("minCyclicShiftV2(%q, %q) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func Test_minCyclicShiftV2_maxSize(t *testing.T) {
	x := strings.Repeat("ab", 5000)
	y := x[len(x)-1:] + x[:len(x)-1]

	if got := minCyclicShiftV2(x, y); got != 1 {
		t.Errorf("minCyclicShiftV2(max size) = %d, want 1", got)
	}
}
