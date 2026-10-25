package main

import (
	"reflect"
	"testing"
)

func TestEqualSubstrings(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		a      int
		b      int
		length int
		want   bool
	}{
		{"same position", "abc", 0, 0, 3, true},
		{"equal different positions", "acabaca", 4, 0, 3, true},
		{"different", "acabaca", 3, 2, 4, false},
		{"single characters different", "abc", 0, 2, 1, false},
		{"single character same", "abc", 1, 1, 1, true},
		{"overlapping equal", "aaaa", 1, 0, 3, true},
		{"overlapping different", "abab", 0, 1, 3, false},
		{"whole string", "abcabc", 0, 3, 3, true},
		{"end of string", "abcabc", 3, 0, 3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := equalSubstrings(tt.s, tt.a, tt.b, tt.length)
			if got != tt.want {
				t.Errorf("equalSubstrings(%q, %d, %d, %d) = %v, want %v", tt.s, tt.a, tt.b, tt.length, got, tt.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		queries [][3]int
		want    []string
	}{
		{
			name:    "example 1",
			s:       "acabaca",
			queries: [][3]int{{4, 3, 2}, {3, 4, 0}, {2, 0, 1}},
			want:    []string{"no", "yes", "no"},
		},
		{
			name: "example 2",
			s:    "caeabaeadedcbdcdccec",
			queries: [][3]int{
				{13, 4, 3},
				{2, 12, 15},
				{10, 1, 3},
				{3, 8, 15},
				{13, 5, 6},
				{7, 2, 6},
				{9, 8, 8},
				{19, 0, 0},
				{19, 0, 0},
				{6, 7, 13},
			},
			want: []string{"no", "no", "no", "no", "no", "no", "yes", "yes", "yes", "no"},
		},
		{
			name:    "no queries",
			s:       "abc",
			queries: [][3]int{},
			want:    []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.s, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve(%q, %v) = %v, want %v", tt.s, tt.queries, got, tt.want)
			}
		})
	}
}
