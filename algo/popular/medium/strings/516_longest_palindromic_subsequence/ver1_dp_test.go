package _16_longest_palindromic_subsequence

import "testing"

func Test_longestPalindromeSubseqV1(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{s: "bbbab"},
			want: 4, // "bbbb"
		},
		{
			name: "Test case 2",
			args: args{s: "cbbd"},
			want: 2, // "bb"
		},
		{
			name: "Test case 3",
			args: args{s: "a"},
			want: 1, // "a"
		},
		{
			name: "Test case 4",
			args: args{s: "abcdefgh"},
			want: 1, // "a"
		},
		{
			name: "Test case 5",
			args: args{s: "racecar"},
			want: 7, // "racecar"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseqV1(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubseqV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
