package _04_sum_of_left_leaves

import "testing"

func Test_sumOfLeftLeavesV1(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: 24,
		},
		{
			name: "CASE-2",
			args: args{
				root: &TreeNode{
					Val: 1,
				},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeavesV1(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeavesV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
