package main

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{4, 1, 2, 1, 2, 6, 6, 7, 7},
			},
			want: 4,
		},
		{
			name: "CASE-4",
			args: args{
				nums: []int{4, 1, 2, 1, 2, 6, 6, 7, 7, 4},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXor(t *testing.T) {
	testCases := []struct {
		nums []int
	}{
		{
			nums: []int{1, 1},
		},
	}

	for _, tc := range testCases {
		res := 0
		for _, num := range tc.nums {
			fmt.Printf("num=%d[%08b] res=%d[%08b] res^num=%d[%08b] \n",
				num, num,
				res, res,
				res^num, res^num,
			)
			res ^= num
		}
	}
}
