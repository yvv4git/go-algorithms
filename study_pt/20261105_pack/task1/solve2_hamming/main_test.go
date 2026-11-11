package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_correctWithHamming(t *testing.T) {
	tests := []struct {
		name string
		mode int
		s    string
		want string
	}{
		{name: "example encode", mode: 1, s: "1010", want: "1011010"},
		{name: "example decode without error", mode: 2, s: "1011010", want: "1010"},
		{name: "example decode data error", mode: 2, s: "1011110", want: "1010"},
		{name: "example decode parity error", mode: 2, s: "1010010", want: "1010"},
		{name: "empty encode", mode: 1, s: "", want: ""},
		{name: "empty decode", mode: 2, s: "", want: ""},
		{name: "single bit encode", mode: 1, s: "0", want: "000"},
		{name: "single bit decode", mode: 2, s: "010", want: "0"},
		{name: "all zeros encode", mode: 1, s: "0000", want: "0000000"},
		{name: "invalid mode", mode: 3, s: "1010", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := correctWithHamming(tt.mode, tt.s); got != tt.want {
				t.Fatalf("correctWithHamming(%d, %q) = %q, want %q", tt.mode, tt.s, got, tt.want)
			}
		})
	}
}

func Test_correctWithHamming_allSingleErrors(t *testing.T) {
	const x = "0101011010"
	y := correctWithHamming(1, x)

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
			if got := correctWithHamming(2, string(z)); got != x {
				t.Fatalf("correctWithHamming(2, %q) = %q, want %q", string(z), got, x)
			}
		})
	}
}

func Test_correctWithHamming_maxSize(t *testing.T) {
	const n = 100000
	x := strings.Repeat("10", n/2)
	y := correctWithHamming(1, x)

	r := hammingParityBits(n)
	if len(y) != n+r {
		t.Fatalf("len(y) = %d, want %d", len(y), n+r)
	}

	z := []byte(y)
	if z[len(z)/2] == '0' {
		z[len(z)/2] = '1'
	} else {
		z[len(z)/2] = '0'
	}
	if got := correctWithHamming(2, string(z)); got != x {
		t.Fatalf("correctWithHamming(2, z) returned a string of length %d, want %d", len(got), n)
	}
}
