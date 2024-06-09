package main

import (
	"reflect"
	"testing"
)

func Test_singleNumberV2(t *testing.T) {
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
				nums: []int{1, 2, 1, 3, 2, 5},
			},
			want: []int{3, 5},
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{1, 0},
			},
			want: []int{1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberV2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumberV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
