package main

import (
	"reflect"
	"testing"
)

func Test_groupAnagramsV2(t *testing.T) {
	type args struct {
		strs []string
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test Case 1: No Anagrams",
			args: args{
				strs: []string{"cat", "dog", "bird", "fish"},
			},
			want: [][]string{
				{"cat"},
				{"dog"},
				{"bird"},
				{"fish"},
			},
		},
		{
			name: "Test Case 2: All Anagrams",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name: "Test Case 3: Empty Input",
			args: args{
				strs: []string{},
			},
			want: [][]string{},
		},
		{
			name: "Test Case 4: Single Element",
			args: args{
				strs: []string{"single"},
			},
			want: [][]string{
				{"single"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagramsV2(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagramsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
