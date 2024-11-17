package main

import "testing"

func Test_shortestDistance(t *testing.T) {
	type args struct {
		wordsDict []string
		word1     string
		word2     string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
				word1:     "coding",
				word2:     "practice",
			},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{
				wordsDict: []string{"apple", "banana", "cherry", "apple", "banana"},
				word1:     "apple",
				word2:     "banana",
			},
			want: 1,
		},
		{
			name: "Test Case 3",
			args: args{
				wordsDict: []string{"hello", "world", "hello", "world", "hello"},
				word1:     "hello",
				word2:     "world",
			},
			want: 1,
		},
		{
			name: "Test Case 4",
			args: args{
				wordsDict: []string{"a", "b", "c", "d", "e"},
				word1:     "a",
				word2:     "e",
			},
			want: 4,
		},
		{
			name: "Test Case 5",
			args: args{
				wordsDict: []string{"one", "two", "three", "four", "five"},
				word1:     "two",
				word2:     "four",
			},
			want: 2,
		},
		{
			name: "Test Case 6",
			args: args{
				wordsDict: []string{"same", "word", "same", "word", "same"},
				word1:     "same",
				word2:     "word",
			},
			want: 1,
		},
		{
			name: "Test Case 7",
			args: args{
				wordsDict: []string{"different", "words", "different", "words", "different"},
				word1:     "different",
				word2:     "words",
			},
			want: 1,
		},
		{
			name: "Test Case 8",
			args: args{
				wordsDict: []string{"first", "second", "third", "fourth", "fifth"},
				word1:     "first",
				word2:     "fifth",
			},
			want: 4,
		},
		{
			name: "Test Case 9",
			args: args{
				wordsDict: []string{"a", "a", "a", "b", "b", "b"},
				word1:     "a",
				word2:     "b",
			},
			want: 1,
		},
		{
			name: "Test Case 10",
			args: args{
				wordsDict: []string{"x", "y", "z", "x", "y", "z"},
				word1:     "x",
				word2:     "z",
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistance(tt.args.wordsDict, tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("shortestDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
