package ver1_dp

import "testing"

func Test_shortestPalindrome(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1: Palindrome string",
			args: args{s: "aacecaaa"},
			want: "aaacecaaa",
		},
		{
			name: "Test case 2: Non-palindrome string",
			args: args{s: "abcd"},
			want: "dcbabcd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("shortestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
