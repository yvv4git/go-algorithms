package main

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case with k = 2",
			args: args{
				s: "abcdefg",
				k: 2,
			},
			want: "bacdfeg",
		},
		{
			name: "Test case with k = 3",
			args: args{
				s: "abcdefgh",
				k: 3,
			},
			want: "cbadefhg",
		},
		{
			name: "Test case with k = 1",
			args: args{
				s: "hello",
				k: 1,
			},
			want: "hello",
		},
		{
			name: "Test case with k = length of string",
			args: args{
				s: "test",
				k: 4,
			},
			want: "tset",
		},
		{
			name: "Test case with k > length of string",
			args: args{
				s: "test",
				k: 5,
			},
			want: "tset",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
