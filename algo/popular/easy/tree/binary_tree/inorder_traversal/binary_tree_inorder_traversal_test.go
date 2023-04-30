package main

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASA-1",
			args: args{
				root: func() *TreeNode {
					tn3 := &TreeNode{
						Val: 3,
					}
					tn2 := &TreeNode{
						Val:   2,
						Left:  tn3,
						Right: nil,
					}
					tn1 := &TreeNode{
						Val:   1,
						Left:  nil,
						Right: tn2,
					}
					return tn1
				}(),
			},
			want: []int{1, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
