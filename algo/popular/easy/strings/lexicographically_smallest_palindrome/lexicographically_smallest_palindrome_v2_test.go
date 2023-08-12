package lexicographically_smallest_palindrome

import "testing"

func Test_makeSmallestPalindromeV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				s: "egcfe",
			},
			want: "efcfe",
			desc: `Explanation: The minimum number of operations to make "egcfe" a palindrome is 1, 
			and the lexicographically smallest palindrome string we can get by modifying one character is "efcfe", 
			by changing 'g'.`,
		},
		{
			name: "CASE-2",
			args: args{
				s: "abcd",
			},
			want: "abba",
			desc: `Explanation: The minimum number of operations to make "abcd" a palindrome is 2, 
			and the lexicographically smallest palindrome string we can get by modifying two characters is "abba".`,
		},
		{
			name: "CASE-3",
			args: args{
				s: "seven",
			},
			want: "neven",
			desc: `Explanation: The minimum number of operations to make "seven" a palindrome is 1, 
			and the lexicographically smallest palindrome string we can get by modifying one character is "neven".`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeSmallestPalindromeV2(tt.args.s); got != tt.want {
				t.Errorf("makeSmallestPalindromeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
