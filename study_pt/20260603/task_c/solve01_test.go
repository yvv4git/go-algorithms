package main

import (
	"reflect"
	"testing"
)

func TestZFunction(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{"single char", "a", []int{0}},
		{"two same", "aa", []int{0, 1}},
		{"two diff", "ab", []int{0, 0}},
		{"all same", "aaaaa", []int{0, 4, 3, 2, 1}},
		{"example", "abracadabra", []int{0, 0, 0, 1, 0, 1, 0, 4, 0, 0, 1}},
		{"abcabc", "abcabc", []int{0, 0, 0, 3, 0, 0}},
		{"aabaaab", "aabaaab", []int{0, 1, 0, 2, 3, 1, 0}},
		{"no matches", "abcdef", []int{0, 0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zFunction(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zFunction(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestZFunctionLargeInput(t *testing.T) {
	n := 1000000
	s := make([]byte, n)
	for i := range s {
		s[i] = 'a'
	}
	got := zFunction(string(s))
	if len(got) != n {
		t.Fatalf("expected length %d, got %d", n, len(got))
	}
	for i, v := range got {
		expected := n - i
		if i == 0 {
			expected = 0
		}
		if v != expected {
			t.Fatalf("at index %d: expected %d, got %d", i, expected, v)
		}
	}
}
