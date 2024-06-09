package main

import (
	"reflect"
	"testing"
)

func Test_subsetsV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{0},
			},
			want: [][]int{
				{},
				{0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsV2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
