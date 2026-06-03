package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	if reverse("abc") != "cba" {
		t.Error("reverse of abc should be cba")
	}
	if reverse("a") != "a" {
		t.Error("reverse of a should be a")
	}
	if reverse("") != "" {
		t.Error("reverse of empty should be empty")
	}
}

func TestFindMatchesAtMostOneMismatch(t *testing.T) {
	tests := []struct {
		name string
		p    string
		t    string
		want []string
	}{
		{"example", "aaaa", "Caaabdaaaa", []string{"1", "2", "6", "7"}},
		{"exact match", "abc", "abcabc", []string{"1", "4"}},
		{"one mismatch middle", "abcd", "abxd", []string{"1"}},
		{"one mismatch first", "abcd", "xbcd", []string{"1"}},
		{"one mismatch last", "abcd", "abcx", []string{"1"}},
		{"two mismatches", "abcd", "abxx", nil},
		{"pattern longer", "abcd", "abc", nil},
		{"single char", "a", "abcba", []string{"1", "2", "3", "4", "5"}},
		{"no match", "abc", "def", nil},
		{"overlap", "aaa", "aaabaaa", []string{"1", "2", "3", "4", "5"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, n := len(tt.p), len(tt.t)
			if m > n {
				if len(tt.want) != 0 {
					t.Errorf("expected no matches when pattern longer, got %v", tt.want)
				}
				return
			}
			zFwd := zFunction(tt.p + "#" + tt.t)
			zBwd := zFunction(reverse(tt.p) + "#" + reverse(tt.t))

			var got []string
			for i := 0; i <= n-m; i++ {
				forward := zFwd[m+1+i]
				if forward >= m {
					got = append(got, fmt.Sprint(i+1))
					continue
				}
				revPos := n - i - m
				backward := zBwd[m+1+revPos]
				if forward+backward >= m-1 {
					got = append(got, fmt.Sprint(i+1))
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
