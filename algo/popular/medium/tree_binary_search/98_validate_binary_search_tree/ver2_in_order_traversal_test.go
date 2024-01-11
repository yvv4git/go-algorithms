package main

import "testing"

func Test_isValidBSTV2(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "invalid BST",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 6,
							Left: &TreeNode{
								Val: 4,
							},
							Right: &TreeNode{
								Val: 7,
							},
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
					Right: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBSTV2(tt.args.root); got != tt.want {
				t.Errorf("isValidBSTV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
