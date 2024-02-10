package main

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{0, 2, 1, 5, 3, 4},
			},
			want: []int{0, 1, 2, 4, 5, 3},
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{5, 0, 1, 2, 3, 4},
			},
			want: []int{4, 5, 0, 1, 2, 3},
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{4, 3, 2, 1, 0},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "Test Case 4",
			args: args{
				nums: []int{0, 1, 2, 3, 4},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "Test Case 5",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
