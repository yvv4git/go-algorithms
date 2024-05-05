package main

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
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
				nums: []int{4, 2, 5, 7},
			},
			want: []int{4, 5, 2, 7},
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{2, 3},
			},
			want: []int{2, 3},
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{2, 1, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParityII(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParityII() = %v, want %v", got, tt.want)
			}
		})
	}
}
