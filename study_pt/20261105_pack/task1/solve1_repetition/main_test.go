package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_correctWithRepetition(t *testing.T) {
	tests := []struct {
		name string
		mode int
		s    string
		want string
	}{
		{name: "example encode", mode: 1, s: "01010", want: "000111000111000"},
		{name: "example decode without error", mode: 2, s: "000111000111000", want: "01010"},
		{name: "example decode one error", mode: 2, s: "000101000111000", want: "01010"},
		{name: "decode first bit error", mode: 2, s: "100111000111000", want: "01010"},
		{name: "decode last bit error", mode: 2, s: "000111000111001", want: "01010"},
		{name: "empty encode", mode: 1, s: "", want: ""},
		{name: "empty decode", mode: 2, s: "", want: ""},
		{name: "single bit encode", mode: 1, s: "1", want: "111"},
		{name: "single bit decode", mode: 2, s: "101", want: "1"},
		{name: "all zeros encode", mode: 1, s: "000", want: "000000000"},
		{name: "invalid mode", mode: 3, s: "01010", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := correctWithRepetition(tt.mode, tt.s); got != tt.want {
				t.Fatalf("correctWithRepetition(%d, %q) = %q, want %q", tt.mode, tt.s, got, tt.want)
			}
		})
	}
}

func Test_correctWithRepetition_allSingleErrors(t *testing.T) {
	const x = "1011001110"
	y := correctWithRepetition(1, x)

	tests := make([]struct {
		name string
		pos  int
	}, len(y))
	for i := range y {
		tests[i].name = fmt.Sprintf("flip position %d", i+1)
		tests[i].pos = i
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := []byte(y)
			if z[tt.pos] == '0' {
				z[tt.pos] = '1'
			} else {
				z[tt.pos] = '0'
			}
			if got := correctWithRepetition(2, string(z)); got != x {
				t.Fatalf("correctWithRepetition(2, %q) = %q, want %q", string(z), got, x)
			}
		})
	}
}

func Test_correctWithRepetition_maxSize(t *testing.T) {
	const n = 100000
	x := strings.Repeat("1", n)
	y := correctWithRepetition(1, x)
	if len(y) != 3*n {
		t.Fatalf("len(y) = %d, want %d", len(y), 3*n)
	}

	z := []byte(y)
	z[len(z)/2] = '0'
	if got := correctWithRepetition(2, string(z)); got != x {
		t.Fatalf("correctWithRepetition(2, z) returned a string of length %d, want %d", len(got), n)
	}
}
