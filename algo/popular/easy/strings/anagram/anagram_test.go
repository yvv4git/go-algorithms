package main

import (
	"reflect"
	"testing"
)

func TestAnagrams(t *testing.T) {
	type args struct {
		word  string
		words []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "CASE-1",
			args: args{
				word:  "abba",
				words: []string{"aabb", "abcd", "bbaa", "dada"},
			},
			want: []string{"aabb", "bbaa"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Anagrams(tt.args.word, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Anagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
