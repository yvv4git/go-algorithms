package main

import "testing"

func Test_findMinimumOperations(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"EqualStrings", args{"abc", "abb", "ab"}, 2},
		{"EqualStrings", args{"dac", "bac", "cac"}, -1},
		{"NotPossibleToForm", args{"abc", "def", "abdecf"}, -1},
		{"OneCharacterEachNotPossible", args{"a", "b", "ba"}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinimumOperations(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("findMinimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
