package main

import (
	"reflect"
	"testing"
)

func Test_levelOrderV2(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		//{
		//	name: "Empty tree",
		//	args: args{root: nil},
		//	want: [][]int{},
		//},
		{
			name: "Single node tree",
			args: args{root: &TreeNode{Val: 1}},
			want: [][]int{{1}},
		},
		{
			name: "Two levels tree",
			args: args{root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val: 3,
				},
			}},
			want: [][]int{{1}, {2, 3}},
		},
		{
			name: "Three levels tree",
			args: args{root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  &TreeNode{Val: 10},
					Right: &TreeNode{Val: 11},
				},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			}},
			want: [][]int{{3}, {9, 20}, {10, 11, 15, 7}},
		},
		{
			name: "Unbalanced tree",
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 5,
							},
						},
					},
				},
			}},
			want: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderV2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
