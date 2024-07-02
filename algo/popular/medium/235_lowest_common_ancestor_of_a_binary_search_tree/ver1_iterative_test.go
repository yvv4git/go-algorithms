package main

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "LCA of 2 and 8 is 6",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
				p: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
				q: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
			},
			want: &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}},
		},
		{
			name: "LCA of 2 and 4 is 2",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
				p: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
				q: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}},
			},
			want: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		},
		{
			name: "LCA of 0 and 5 is 2",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
				p: &TreeNode{Val: 0},
				q: &TreeNode{Val: 5},
			},
			want: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		},
		{
			name: "LCA of 7 and 9 is 8",
			args: args{
				root: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 0,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 3,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 9,
						},
					},
				},
				p: &TreeNode{Val: 7},
				q: &TreeNode{Val: 9},
			},
			want: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
