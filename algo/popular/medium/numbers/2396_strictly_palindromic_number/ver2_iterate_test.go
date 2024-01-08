package main

import "testing"

func Test_isStrictlyPalindromicV2(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{n: 4},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{n: 9},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{n: 15},
			want: false,
		},
		{
			name: "Test case 4",
			args: args{n: 2},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{n: 3},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStrictlyPalindromicV2(tt.args.n); got != tt.want {
				t.Errorf("isStrictlyPalindromicV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
