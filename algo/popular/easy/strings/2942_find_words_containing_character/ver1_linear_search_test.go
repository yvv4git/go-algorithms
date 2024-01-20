package main

import (
	"reflect"
	"testing"
)

func Test_findWordsContaining(t *testing.T) {
	type args struct {
		words []string
		x     byte
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test with words containing 'o'",
			args: args{
				words: []string{"hello", "world", "golang", "programming", "code"},
				x:     'o',
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "Test with words containing 'a'",
			args: args{
				words: []string{"abc", "bcd", "aaaa", "cbc"},
				x:     'a',
			},
			want: []int{0, 2},
		},
		{
			name: "Test with words containing 'z'",
			args: args{
				words: []string{"abc", "bcd", "aaaa", "cbc"},
				x:     'z',
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWordsContaining(tt.args.words, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWordsContaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
