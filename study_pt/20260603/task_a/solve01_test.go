package main

import (
	"reflect"
	"testing"
)

func TestPrefixFunction(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{"single char", "a", []int{0}},
		{"two same", "aa", []int{0, 1}},
		{"two diff", "ab", []int{0, 0}},
		{"all same", "aaaaa", []int{0, 1, 2, 3, 4}},
		{"repeat pattern", "abcabc", []int{0, 0, 0, 1, 2, 3}},
		{"example", "abracadabra", []int{0, 0, 0, 1, 0, 1, 0, 1, 2, 3, 4}},
		{"aabaaab", "aabaaab", []int{0, 1, 0, 1, 2, 2, 3}},
		{"no matches", "abcdef", []int{0, 0, 0, 0, 0, 0}},
		{"border", "abacab", []int{0, 0, 1, 0, 1, 2}},
		{"full match prefix", "aaaab", []int{0, 1, 2, 3, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixFunction(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixFunction(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestPrefixFunctionLargeInput(t *testing.T) {
	n := 1000000
	s := make([]byte, n)
	for i := range s {
		s[i] = 'a'
	}
	got := prefixFunction(string(s))
	if len(got) != n {
		t.Fatalf("expected length %d, got %d", n, len(got))
	}
	for i, v := range got {
		if v != i {
			t.Fatalf("at index %d: expected %d, got %d", i, i, v)
		}
	}
}
