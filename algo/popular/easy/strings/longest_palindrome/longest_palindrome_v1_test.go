package longest_palindrome

import "testing"

func Test_longestPalindromeV1(t *testing.T) {
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
				s: "abccccdd",
			},
			want: 7,
		},
		{
			name: "CASE-1",
			args: args{
				s: "a",
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeV1(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
