package _66_palindrome_permutation

import "testing"

func Test_canPermutePalindromeV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{s: "aab"},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{s: "carerac"},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{s: "code"},
			want: false,
		},
		{
			name: "Test case 4",
			args: args{s: "aabb"},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{s: "aabbc"},
			want: true,
		},
		{
			name: "Test case 6",
			args: args{s: "aabbcc"},
			want: true,
		},
		{
			name: "Test case 7",
			args: args{s: "aabbccd"},
			want: true,
		},
		{
			name: "Test case 8",
			args: args{s: "aabbccde"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPermutePalindromeV2(tt.args.s); got != tt.want {
				t.Errorf("canPermutePalindromeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
