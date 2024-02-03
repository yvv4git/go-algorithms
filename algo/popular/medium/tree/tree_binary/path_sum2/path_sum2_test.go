package path_sum2

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "CASE-1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{
								Val: 2,
							},
						},
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val: 13,
						},
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 5,
							},
							Right: &TreeNode{
								Val: 1,
							},
						},
					},
				},
				targetSum: 22,
			},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			name: "CASE-2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				targetSum: 5,
			},
			want: [][]int{},
		},
		{
			name: "CASE-3",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
				targetSum: 0,
			},
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
