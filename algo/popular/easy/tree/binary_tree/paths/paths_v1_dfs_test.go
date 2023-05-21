package paths

import (
	"reflect"
	"testing"
)

func Test_binaryTreePathsV1(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "CASE-1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []string{"1->2->5", "1->3"},
		},
		{
			name: "CASE-2",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: []string{"1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreePathsV1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreePathsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
