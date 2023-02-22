package main

import (
	"strconv"
	"testing"
)

func Test_isPalindromeInt64(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				n: 121,
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				n: 123454321,
			},
			want: true,
		},
		{
			name: "CASE-3",
			args: args{
				n: 1234567,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeInt64(tt.args.n); got != tt.want {
				t.Errorf("isPalindromeInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

// IsPalindromeInt64 несколько оптимизирован за счет.
func isPalindromeInt64(n int64) bool {
	if n < 0 {
		n = -n
	}

	s := strconv.FormatInt(n, 10)
	bound := (len(s) / 2) + 1
	for i := 0; i < bound; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
