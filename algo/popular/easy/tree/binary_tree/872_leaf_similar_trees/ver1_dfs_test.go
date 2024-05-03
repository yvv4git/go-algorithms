package main

import "testing"

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1: Same trees",
			args: args{
				root1: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}},
				root2: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}},
			},
			want: true,
		},
		{
			name: "Test Case 2: Different trees with different leaf sequence",
			args: args{
				root1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
				root2: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}},
			},
			want: false,
		},
		{
			name: "Test Case 3: One empty tree",
			args: args{
				root1: nil,
				root2: &TreeNode{Val: 1},
			},
			want: false,
		},
		{
			name: "Test Case 4: Both empty trees",
			args: args{
				root1: nil,
				root2: nil,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
