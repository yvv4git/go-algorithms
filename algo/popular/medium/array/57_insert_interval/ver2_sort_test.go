package main

import (
	"reflect"
	"testing"
)

func Test_insertV2(t *testing.T) {
	type args struct {
		intervals [][]int
		new       []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test case 1",
			args: args{
				intervals: [][]int{{1, 3}, {6, 9}},
				new:       []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "Test case 2",
			args: args{
				intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				new:       []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name: "Test case 3",
			args: args{
				intervals: [][]int{{1, 5}},
				new:       []int{6, 8},
			},
			want: [][]int{{1, 5}, {6, 8}},
		},
		{
			name: "Test case 4",
			args: args{
				intervals: [][]int{{1, 5}},
				new:       []int{0, 0},
			},
			want: [][]int{{0, 0}, {1, 5}},
		},
		{
			name: "Test case 5",
			args: args{
				intervals: [][]int{{1, 5}},
				new:       []int{0, 3},
			},
			want: [][]int{{0, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertV2(tt.args.intervals, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
