package main

import (
	"reflect"
	"testing"
)

func Test_buildTreeRecursion(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Test Case 1",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				preorder: []int{1, 2, 3},
				inorder:  []int{2, 1, 3},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			name: "Test Case 3",
			args: args{
				preorder: []int{1},
				inorder:  []int{1},
			},
			want: &TreeNode{
				Val: 1,
			},
		},
		{
			name: "Test Case 4",
			args: args{
				preorder: []int{},
				inorder:  []int{},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTreeRecursion(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTreeRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
