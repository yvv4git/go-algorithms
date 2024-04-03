package main

import "testing"

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1: Same strings",
			args: args{
				word1: "kitten",
				word2: "kitten",
			},
			want: 0, // No edits needed as strings are the same
		},
		{
			name: "Test case 2: One insertion",
			args: args{
				word1: "kitten",
				word2: "sitting",
			},
			want: 3, // Need to replace 's' with 'k', replace 'e' with 'i', and insert 'g'
		},
		{
			name: "Test case 3: One deletion",
			args: args{
				word1: "sitting",
				word2: "kitten",
			},
			want: 3, // Need to replace 's' with 'k', replace 'e' with 'i', and delete 'g'
		},
		{
			name: "Test case 4: Completely different strings",
			args: args{
				word1: "cat",
				word2: "dog",
			},
			want: 3, // Need to replace 'c' with 'd', replace 'a' with 'o', and replace 't' with 'g'
		},
		{
			name: "Test case 5: Empty strings",
			args: args{
				word1: "",
				word2: "",
			},
			want: 0, // No edits needed as strings are empty
		},
		{
			name: "Test case 6: One empty string",
			args: args{
				word1: "",
				word2: "hello",
			},
			want: 5, // Need to insert 'h', 'e', 'l', 'l', 'o'
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
