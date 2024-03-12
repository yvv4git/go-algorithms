package main

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrderV2(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1: Symmetric Tree",
			args: args{
				root: &TreeNode{
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
			want: [][]int{
				{1},
				{3, 2},
				{4, 5, 6, 7},
			},
		},
		{
			name: "Test Case 2: Asymmetric Tree",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: [][]int{
				{1},
				{3, 2},
				{4},
			},
		},
		{
			name: "Test Case 3: Single Node",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "Test Case 4: Empty Tree",
			args: args{
				root: nil,
			},
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zigzagLevelOrderV2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrderV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
