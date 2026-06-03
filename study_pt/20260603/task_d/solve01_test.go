package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsVowel(t *testing.T) {
	vowels := "aeiouy"
	consonants := "bcdfghjklmnpqrstvwxz"
	for _, c := range vowels {
		if !isVowel(byte(c)) {
			t.Errorf("expected %c to be vowel", c)
		}
	}
	for _, c := range consonants {
		if isVowel(byte(c)) {
			t.Errorf("expected %c to be consonant", c)
		}
	}
}

func TestClass(t *testing.T) {
	if class('a') != '0' || class('y') != '0' {
		t.Error("vowel should map to 0")
	}
	if class('b') != '1' || class('z') != '1' {
		t.Error("consonant should map to 1")
	}
}

func TestTransform(t *testing.T) {
	if transform("aeiouy") != "000000" {
		t.Error("all vowels should map to 0s")
	}
	if transform("bcdfg") != "11111" {
		t.Error("all consonants should map to 1s")
	}
	if transform("ab") != "01" {
		t.Errorf("got %q, want 01", transform("ab"))
	}
}

func TestFindMatches(t *testing.T) {
	tests := []struct {
		name string
		t    string
		s    string
		want []string
	}{
		{"example 1", "abracadabra", "abr", []string{"0", "7"}},
		{"example 2", "notsoverylongtext", "word", []string{"0", "9", "13"}},
		{"vowel match", "abecid", "abcd", nil},
		{"no match", "abc", "xyz", nil},
		{"single char match", "abcabc", "a", []string{"0", "3"}},
		{"all vowels", "aeiouyaeiouy", "aaaa", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8"}},
		{"all consonants", "bcdfgbcdfg", "bbbb", []string{"0", "1", "2", "3", "4", "5", "6"}},
		{"same length mismatch", "abcdef", "ghijkl", nil},
		{"pattern longer", "abc", "abcd", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := transform(tt.t)
			ps := transform(tt.s)
			pi := prefixFunction(ps)

			var got []string
			j := 0
			for i := 0; i < len(pt); i++ {
				for j > 0 && pt[i] != ps[j] {
					j = pi[j-1]
				}
				if pt[i] == ps[j] {
					j++
				}
				if j == len(ps) {
					got = append(got, fmt.Sprint(i-len(ps)+1))
					j = pi[j-1]
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
