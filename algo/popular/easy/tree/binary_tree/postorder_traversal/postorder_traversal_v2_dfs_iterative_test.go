package postorder_traversal

import (
	"reflect"
	"testing"
)

func Test_postorderTraversalV2(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []int{1, 3, 2},
		},
		{
			name: "CASE-2",
			args: args{
				root: nil,
			},
			want: []int{},
		},
		{
			name: "CASE-3",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversalV2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversalV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
