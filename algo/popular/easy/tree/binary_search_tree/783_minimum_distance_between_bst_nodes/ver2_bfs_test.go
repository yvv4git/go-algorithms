package main

import "testing"

func Test_minDiffInBSTV2(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 48,
						Left: &TreeNode{
							Val: 12,
						},
						Right: &TreeNode{
							Val: 49,
						},
					},
				},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDiffInBSTV2(tt.args.root); got != tt.want {
				t.Errorf("minDiffInBSTV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
