package main

import "testing"

func Test_checkInt32(t *testing.T) {
	type args struct {
		x int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				x: 123454321,
			},
			want: true,
		},
		{
			name: "CASE-3",
			args: args{
				x: 1234554321,
			},
			want: true,
		},
		{
			name: "CASE-4",
			args: args{
				x: 1234564321,
			},
			want: false,
		},
		{
			name: "CASE-5",
			args: args{
				x: 12345,
			},
			want: false,
		},
		{
			name: "CASE-6",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "CASE-7",
			args: args{
				x: -12321,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeInt32(tt.args.x); got != tt.want {
				t.Errorf("checkInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isPalindromeInt32(x int32) bool {
	var reverse, remainder, tmp int32
	reverse = 0

	if x < 0 {
		return false
	}

	tmp = x

	for {
		remainder = x % 10
		reverse = reverse*10 + remainder
		x /= 10

		if x == 0 {
			break
		}
	}

	return tmp == reverse
}
