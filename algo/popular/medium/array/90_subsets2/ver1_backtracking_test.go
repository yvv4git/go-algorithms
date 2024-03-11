package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1: No duplicates",
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "Test Case 2: Duplicates",
			args: args{nums: []int{1, 2, 2}},
			want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
		{
			name: "Test Case 3: Single element",
			args: args{nums: []int{1}},
			want: [][]int{{}, {1}},
		},
		{
			name: "Test Case 4: Empty array",
			args: args{nums: []int{}},
			want: [][]int{{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetsWithDup(tt.args.nums)
			sort.Slice(got, func(i, j int) bool {
				return less(got[i], got[j])
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return less(tt.want[i], tt.want[j])
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsWithDupV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func less(a, b []int) bool {
	if len(a) != len(b) {
		return len(a) < len(b)
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return false
}
