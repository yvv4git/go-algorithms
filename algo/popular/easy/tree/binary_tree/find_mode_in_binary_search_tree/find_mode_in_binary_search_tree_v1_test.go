package find_mode_in_binary_search_tree

import (
	"reflect"
	"testing"
)

func Test_findMode(t *testing.T) {
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
						Val: 2,
					},
				},
			},
			want: []int{2},
		},
		{
			name: "CASE-2",
			args: args{
				root: &TreeNode{
					Val: 0,
				},
			},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findModeV1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findModeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
