package main

import (
	"reflect"
	"testing"
)

func Test_relativeSortArray(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				arr1: []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
				arr2: []int{2, 1, 4, 3, 9},
			},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			name: "Test Case 2",
			args: args{
				arr1: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
				arr2: []int{2, 1, 4, 3, 9},
			},
			want: []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			name: "Test Case 3",
			args: args{
				arr1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name: "Test Case 4",
			args: args{
				arr1: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				arr2: []int{},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relativeSortArray(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relativeSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
