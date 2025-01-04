package main

import "testing"

func Test_maxPower(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				s: "leetcode",
			},
			want: 2,
		},
		{
			name: "CASE-2",
			args: args{
				s: "abbcccddddeeeeedcba",
			},
			want: 5,
		},
		{
			name: "CASE-3",
			args: args{
				s: "triplepillooooow",
			},
			want: 5,
		},
		{
			name: "CASE-4",
			args: args{
				s: "hooraaaaaaaaaaay",
			},
			want: 11,
		},
		{
			name: "CASE-5",
			args: args{
				s: "tourist",
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.args.s); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
