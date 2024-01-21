package main

import "testing"

func Test_maximumOddBinaryNumberV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "First example",
			args: args{s: "010"},
			want: "001",
		},
		{
			name: "Second example",
			args: args{s: "0101"},
			want: "1001",
		},
		{
			name: "All ones",
			args: args{s: "1111"},
			want: "1111",
		},
		{
			name: "Single zero",
			args: args{s: "0"},
			want: "1",
		},
		{
			name: "Single one",
			args: args{s: "1"},
			want: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumOddBinaryNumberV2(tt.args.s); got != tt.want {
				t.Errorf("maximumOddBinaryNumberV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
