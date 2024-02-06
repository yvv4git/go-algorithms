package main

import (
	"reflect"
	"testing"
)

func Test_findAnagramsV2(t *testing.T) {
	type args struct {
		s string
		p string
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				s: "cbaebabacd",
				p: "abc",
			},
			want: []int{0, 6},
		},
		{
			name: "Test Case 2",
			args: args{
				s: "abab",
				p: "ab",
			},
			want: []int{0, 1, 2},
		},
		{
			name: "Test Case 3",
			args: args{
				s: "abacbabc",
				p: "abc",
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "Test Case 4",
			args: args{
				s: "aaaaaaa",
				p: "a",
			},
			want: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			name: "Test Case 5",
			args: args{
				s: "abcdefg",
				p: "hijklmn",
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagramsV2(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagramsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
