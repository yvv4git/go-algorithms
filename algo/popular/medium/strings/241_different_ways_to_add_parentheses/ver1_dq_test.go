package main

import (
	"reflect"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		expression string
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{expression: "1+1+1+1"},
			want: []int{4, 4, 4, 4, 4},
		},
		{
			name: "Test Case 2",
			args: args{expression: "2"},
			want: []int{2},
		},
		{
			name: "Test Case 3",
			args: args{expression: "2*3+4"},
			want: []int{14, 10},
		},
		//{
		//	name: "Test Case 4",
		//	args: args{expression: "2-1-1"},
		//	want: []int{0, 2},
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diffWaysToCompute(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
