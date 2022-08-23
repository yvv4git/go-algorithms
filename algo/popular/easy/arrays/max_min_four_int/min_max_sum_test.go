package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_miniMaxSum(t *testing.T) {
	type args struct {
		arr []int32
	}

	testCases := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "CASE-1",
			args: args{
				arr: []int32{},
			},
			want: []int32{},
		},
		{
			name: "CASE-2",
			args: args{
				arr: []int32{5, 8, 0, 9, 4, 7},
			},
			want: []int32{16, 29},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := miniMaxSum(tt.args.arr)
			assert.Equal(t, tt.want, result)
		})
	}
}
