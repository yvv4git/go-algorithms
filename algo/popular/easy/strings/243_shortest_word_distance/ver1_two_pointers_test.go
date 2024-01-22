package main

import "testing"

func Test_shortestDistance(t *testing.T) {
	type args struct {
		words []string
		word1 string
		word2 string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				words: []string{"practice", "makes", "perfect", "coding", "makes"},
				word1: "coding",
				word2: "practice",
			},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{
				words: []string{"practice", "makes", "perfect", "coding", "makes"},
				word1: "makes",
				word2: "coding",
			},
			want: 1,
		},
		{
			name: "Test Case 3",
			args: args{
				words: []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"},
				word1: "fox",
				word2: "dog",
			},
			want: 5,
		},
		{
			name: "Test Case 4",
			args: args{
				words: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
				word1: "a",
				word2: "z",
			},
			want: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistance(tt.args.words, tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("shortestDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
