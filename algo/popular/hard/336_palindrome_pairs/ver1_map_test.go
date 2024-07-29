package main

import (
	"reflect"
	"testing"
)

func Test_palindromePairs(t *testing.T) {
	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				words: []string{"bat", "tab", "cat"},
			},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "Example 2",
			args: args{
				words: []string{"abcd", "dcba", "lls", "s", "sssll"},
			},
			want: [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}},
		},
		{
			name: "Example 3",
			args: args{
				words: []string{"a", ""},
			},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "Example 4",
			args: args{
				words: []string{"a", "abc", "aba", ""},
			},
			want: [][]int{{0, 3}, {3, 0}, {2, 3}, {3, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePairs(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("palindromePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
