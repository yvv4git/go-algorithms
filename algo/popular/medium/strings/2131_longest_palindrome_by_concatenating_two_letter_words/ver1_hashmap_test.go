package main

import "testing"

func Test_longestPalindromeV1(t *testing.T) {
	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				words: []string{"ab", "ba", "lc", "cl"},
			},
			want: 8,
		},
		{
			name: "Test Case 2",
			args: args{
				words: []string{"ab", "ty", "yt", "lc", "cl", "ab"},
			},
			want: 8,
		},
		{
			name: "Test Case 4",
			args: args{
				words: []string{"ab", "ba", "ab", "ba"},
			},
			want: 8,
		},
		{
			name: "Test Case 5",
			args: args{
				words: []string{"ab", "cd", "dc", "ba"},
			},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeV1(tt.args.words); got != tt.want {
				t.Errorf("longestPalindromeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
