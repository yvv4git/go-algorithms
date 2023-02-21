package main

import (
	"reflect"
	"testing"
)

func Test_mergeV1(t *testing.T) {
	type args struct {
		num1 []int
		num2 []int
	}

	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "CASE-1",
			args: args{
				num1: []int{1, 2, 5},
				num2: []int{1, 2, 3, 4, 6},
			},
			wantResult: []int{1, 1, 2, 2, 3, 4, 5, 6},
		},
		{
			name: "CASE-2",
			args: args{
				num1: []int{1, 1},
				num2: []int{1, 2, 3, 4, 6},
			},
			wantResult: []int{1, 1, 1, 2, 3, 4, 6},
		},
		{
			name: "CASE-3",
			args: args{
				num1: []int{},
				num2: []int{1, 2, 3, 4, 6},
			},
			wantResult: []int{1, 2, 3, 4, 6},
		},
		{
			name: "CASE-4",
			args: args{
				num1: []int{},
				num2: []int{},
			},
			wantResult: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := mergeV1(tt.args.num1, tt.args.num2); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("mergeV1() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	result := make([]int, 5)
	num1 := []int{1, 2, 3, 4, 5}

	copy(result[4:5], num1[0:1])
	t.Logf("%v", result)
}
