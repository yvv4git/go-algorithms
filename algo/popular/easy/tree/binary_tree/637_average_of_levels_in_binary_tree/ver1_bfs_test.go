package main

import (
	"reflect"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Test Case 1",
			args: args{
				root: nil,
			},
			want: []float64{},
		},
		{
			name: "Test Case 2",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: []float64{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
