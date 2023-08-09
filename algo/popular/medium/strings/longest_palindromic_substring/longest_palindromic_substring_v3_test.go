package longest_palindromic_substring

import "testing"

func Test_longestPalindromeV3(t *testing.T) {
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
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "CASE-2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeV3(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
