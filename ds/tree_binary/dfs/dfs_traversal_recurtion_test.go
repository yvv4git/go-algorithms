package dfs

import (
	"testing"
)

func Test_dfs(t *testing.T) {
	type args struct {
		root *Node
	}

	tests := []struct {
		name string
		args args
		want []*Node
	}{
		{
			name: "CASE-1",
			args: args{
				root: &Node{
					Val: 1,
					Left: &Node{
						Val: 2,
						Left: &Node{
							Val: 4,
						},
						Right: &Node{
							Val: 5,
						},
					},
					Right: &Node{
						Val: 3,
						Left: &Node{
							Val: 6,
						},
						Right: &Node{
							Val: 7,
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dfsTraversalRecursion(tt.args.root)
			t.Logf("result: %v", result)

			for idx, node := range result {
				t.Logf("node[%d].Val = %v", idx, node.Val)
			}
		})
	}
}
