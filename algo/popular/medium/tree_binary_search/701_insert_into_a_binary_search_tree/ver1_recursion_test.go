package main

import (
	"reflect"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Insert into empty tree",
			args: args{
				root: nil,
				val:  10,
			},
			want: &TreeNode{Val: 10},
		},
		{
			name: "Insert into non-empty tree",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 15,
					},
				},
				val: 12,
			},
			want: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 12,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
