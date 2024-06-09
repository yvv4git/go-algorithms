package main

import (
	"reflect"
	"testing"
)

func Test_buildTreeV2(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Full binary tree",
			args: args{
				inorder:   []int{4, 2, 5, 1, 6, 3, 7},
				postorder: []int{4, 5, 2, 6, 7, 3, 1},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			name: "Empty tree",
			args: args{
				inorder:   []int{},
				postorder: []int{},
			},
			want: nil,
		},
		{
			name: "Single node tree",
			args: args{
				inorder:   []int{1},
				postorder: []int{1},
			},
			want: &TreeNode{
				Val: 1,
			},
		},
		{
			name: "Left-skewed tree",
			args: args{
				inorder:   []int{3, 2, 1},
				postorder: []int{3, 2, 1},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTreeV2(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTreeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
