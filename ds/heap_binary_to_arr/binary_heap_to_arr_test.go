package main

import (
	"reflect"
	"testing"
)

func Test_convertBinaryTreeToArray(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				root: func() *TreeNode {
					return &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 5,
							},
							Right: &TreeNode{
								Val: 9,
							},
						},
						Right: &TreeNode{
							Val: 6,
							Left: &TreeNode{
								Val: 8,
							},
						},
					}
				}(),
			},
			want: []int{1, 3, 6, 5, 9, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// result := convertBinaryTreeToArray(tt.args.root)
			// t.Logf("Binary tree to array: %v", result)
			if got := convertBinaryTreeToArray(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBinaryTreeToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
