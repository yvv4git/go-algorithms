package remove_palindromic_subsequences

import "testing"

func Test_removePalindromeSubV3(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				s: "ababa",
			},
			want: 1,
			desc: `Explanation: s is already a palindrome, so its entirety can be removed in a single step.`,
		},
		{
			name: "CASE-2",
			args: args{
				s: "abb",
			},
			want: 2,
			desc: `Explanation: "abb" -> "bb" -> "". Remove palindromic subsequence "a" then "bb".`,
		},
		{
			name: "CASE-3",
			args: args{
				s: "baabb",
			},
			want: 2,
			desc: `Explanation: "baabb" -> "b" -> "". Remove palindromic subsequence "baab" then "b".`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removePalindromeSubV3(tt.args.s); got != tt.want {
				t.Errorf("removePalindromeSubV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
