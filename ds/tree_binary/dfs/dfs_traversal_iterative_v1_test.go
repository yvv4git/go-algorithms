package dfs

import (
	"testing"
)

func Test_dfsIterative(t *testing.T) {
	type args struct {
		root *Node
	}

	tests := []struct {
		name string
		args args
		want []int
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
			result := dfsTraversalIterativeV1(tt.args.root)
			t.Logf("result=%#v", result)
		})
	}
}
