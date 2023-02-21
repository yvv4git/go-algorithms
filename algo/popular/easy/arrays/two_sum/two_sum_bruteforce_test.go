package main

import (
	"reflect"
	"testing"
)

func Test_twoSumBruteforce(t *testing.T) {
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
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "CASE-2",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
		{
			name: "CASE-3",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
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
			if got := twoSumBruteforce(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Log(got)
				t.Errorf("twoSumBruteforce() = %v, want %v", got, tt.want)
			}
		})
	}
}
