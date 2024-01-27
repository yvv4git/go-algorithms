package main

import "testing"

func Test_wordBreakV2(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				s:        "leetcode",
				wordDict: []string{"leet", "code"},
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				s:        "applepenapple",
				wordDict: []string{"apple", "pen"},
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				s:        "a",
				wordDict: []string{"a"},
			},
			want: true,
		},
		{
			name: "Test Case 5",
			args: args{
				s:        "ab",
				wordDict: []string{"a", "b"},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreakV2(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreakV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
