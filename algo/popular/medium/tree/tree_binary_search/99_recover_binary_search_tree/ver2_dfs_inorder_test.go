package main

import "testing"

func Test_recoverTreeV2(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Test Case 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			want: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTreeV2(tt.args.root)
		})
	}
}
