package main

import (
	"reflect"
	"testing"
)

func Test_createBinaryTreeFromDescriptions(t *testing.T) {
	type args struct {
		descriptions [][]int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Single node tree",
			args: args{
				descriptions: [][]int{{1, 2, 1}},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			name: "Tree with multiple levels",
			args: args{
				descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}},
			},
			want: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 17,
					},
				},
				Right: &TreeNode{
					Val: 80,
					Left: &TreeNode{
						Val: 19,
					},
				},
			},
		},
		{
			name: "Tree with only right children",
			args: args{
				descriptions: [][]int{{10, 20, 0}, {20, 30, 0}, {30, 40, 0}},
			},
			want: &TreeNode{
				Val: 10,
				Right: &TreeNode{
					Val: 20,
					Right: &TreeNode{
						Val: 30,
						Right: &TreeNode{
							Val: 40,
						},
					},
				},
			},
		},
		{
			name: "Tree with only left children",
			args: args{
				descriptions: [][]int{{10, 20, 1}, {20, 30, 1}, {30, 40, 1}},
			},
			want: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 30,
						Left: &TreeNode{
							Val: 40,
						},
					},
				},
			},
		},
		{
			name: "Empty tree",
			args: args{
				descriptions: [][]int{},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createBinaryTreeFromDescriptions(tt.args.descriptions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBinaryTreeFromDescriptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
