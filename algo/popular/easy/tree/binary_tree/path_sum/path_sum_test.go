package path_sum

import "testing"

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}

	tests := []struct {
		name string
		args args
		want bool
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
							Right: &TreeNode{
								Val: 1,
							},
						},
					},
				},
				targetSum: 22,
			},
			want: true,
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
			want: false,
		},
		{
			name: "CASE-3",
			args: args{
				root:      nil,
				targetSum: 0,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := hasPathSum(tt.args.root, tt.args.targetSum); result != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", result, tt.want)
			}
		})
	}
}
