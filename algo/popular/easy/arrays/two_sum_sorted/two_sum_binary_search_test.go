package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_twoSumBinarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				nums:   []int{1, 2, 7, 10},
				target: 9,
			},
			want: []int{1, 2},
		},
		{
			name: "CASE-2",
			args: args{
				nums:   []int{1, 2, 3, 4, 8},
				target: 6,
			},
			want: []int{1, 3},
		},
		{
			name: "CASE-3",
			args: args{
				nums:   []int{3, 3, 4},
				target: 6,
			},
			want: []int{0, 1},
		},
		{
			name: "CASE-4",
			args: args{
				nums:   []int{3, 2, 4},
				target: 7,
			},
			want: []int{0, 2},
		},
		{
			name: "CASE-5",
			args: args{
				nums:   []int{3, 2, 4},
				target: 999,
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumBinarySearch(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Log(got)
				t.Errorf("twoSumBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVisualizeMid(t *testing.T) {
	/*
		Хочу посмотреть, как вычисляется средний элемент.
	*/

	nums := []int{0, 1, 2, 3, 4, 5}
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2
		fmt.Printf("[%d - %d - %d] \n", l, m, r)
		l++
		r--
	}
}
