package main

import (
	"reflect"
	"testing"
)

func Test_deleteNodeV2(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Delete leaf node",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
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
				key: 4,
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
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
		{
			name: "Delete node with one child",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
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
					},
				},
				key: 7,
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		{
			name: "Delete node with two children",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
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
				key: 5,
			},
			want: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "Delete root node",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
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
				key: 5,
			},
			want: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "Delete non-existent node",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
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
				key: 9,
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
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
		{
			name: "Delete from empty tree",
			args: args{
				root: nil,
				key:  1,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNodeV2(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNodeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
