package main

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case 1: Sorted array",
			args: args{arr: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test case 2: Reverse sorted array",
			args: args{arr: []int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test case 3: Array with duplicate elements",
			args: args{arr: []int{5, 2, 5, 2, 3}},
			want: []int{2, 2, 3, 5, 5},
		},
		{
			name: "Test case 4: Array with negative numbers",
			args: args{arr: []int{-5, -2, 5, 2, 3}},
			want: []int{-5, -2, 2, 3, 5},
		},
		{
			name: "Test case 5: Single element array",
			args: args{arr: []int{5}},
			want: []int{5},
		},
		{
			name: "Test case 6: Empty array",
			args: args{arr: []int{}},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
