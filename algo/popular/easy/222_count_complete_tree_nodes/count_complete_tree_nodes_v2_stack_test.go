package _22_count_complete_tree_nodes

import "testing"

func Test_countNodesV2(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
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
					},
				},
			},
			want: 6,
		},
		{
			name: "CASE-2",
			args: args{},
			want: 0,
		},
		{
			name: "CASE-3",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodesV2(tt.args.root); got != tt.want {
				t.Errorf("countNodesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
