package main

import (
	"reflect"
	"testing"
)

func Test_mergeV2(t *testing.T) {
	type args struct {
		intervals [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1: No overlapping intervals",
			args: args{
				intervals: [][]int{{1, 3}, {5, 7}, {9, 11}},
			},
			want: [][]int{{1, 3}, {5, 7}, {9, 11}},
		},
		{
			name: "Test Case 2: Overlapping intervals",
			args: args{
				intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "Test Case 3: Single interval",
			args: args{
				intervals: [][]int{{1, 3}},
			},
			want: [][]int{{1, 3}},
		},
		{
			name: "Test Case 4: Empty intervals",
			args: args{
				intervals: [][]int{},
			},
			want: [][]int{},
		},
		{
			name: "Test Case 5: Overlapping and non-overlapping intervals",
			args: args{
				intervals: [][]int{{1, 4}, {4, 5}},
			},
			want: [][]int{{1, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeV2(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
