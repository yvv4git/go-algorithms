package array_to_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{-10, -3, 0, 9, 5},
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 9,
					},
				},
			},
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{1, 3},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortedArrayToBST(tt.args.nums)
			assert.EqualValues(t, tt.want, result)
		})
	}
}
