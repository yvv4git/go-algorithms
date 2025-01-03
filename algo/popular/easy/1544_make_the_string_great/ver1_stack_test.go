package main

import "testing"

func Test_makeGood(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				s: "leEeetcode",
			},
			want: "leetcode",
		},
		{
			name: "CASE-2",
			args: args{
				s: "abBAcC",
			},
			want: "",
		},
		{
			name: "CASE-3",
			args: args{
				s: "s",
			},
			want: "s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeGood(tt.args.s); got != tt.want {
				t.Errorf("makeGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
