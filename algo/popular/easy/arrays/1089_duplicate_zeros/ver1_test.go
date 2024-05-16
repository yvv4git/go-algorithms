package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case with no zeros",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test case with one zero at the end",
			args: args{
				arr: []int{1, 2, 0},
			},
			want: []int{1, 2, 0},
		},
		{
			name: "Test case with one zero in the middle",
			args: args{
				arr: []int{1, 0, 2},
			},
			want: []int{1, 0, 0},
		},
		{
			name: "Test case with multiple zeros",
			args: args{
				arr: []int{1, 0, 2, 3, 0, 4, 5, 0},
			},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name: "Test case with zeros at the beginning and end",
			args: args{
				arr: []int{0, 1, 2, 3, 0},
			},
			want: []int{0, 0, 1, 2, 3},
		},
		{
			name: "Test case !!!",
			args: args{
				arr: []int{1, 0, 2, 3, 0, 4, 5, 0},
			},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.args.arr)
			assert.Equal(t, tt.want, tt.args.arr)
		})
	}
}
