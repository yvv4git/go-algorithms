package main

import (
	"reflect"
	"testing"
)

func Test_searchBSTV2(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Тест 1",
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
						Val: 6,
					},
				},
				val: 4,
			},
			want: &TreeNode{
				Val: 4,
			},
		},
		{
			name: "Тест 2",
			args: args{
				root: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 2,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: 10,
						Right: &TreeNode{
							Val: 20,
						},
					},
				},
				val: 20,
			},
			want: &TreeNode{
				Val: 20,
			},
		},
		{
			name: "Тест 3",
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
					},
				},
				val: 7,
			},
			want: &TreeNode{
				Val: 7,
			},
		},
		{
			name: "Тест 4",
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
					},
				},
				val: 1,
			},
			want: &TreeNode{
				Val: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBSTV2(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBSTV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
