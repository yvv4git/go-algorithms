package main

import "testing"

func Test_getMinimumDifferenceV3(t *testing.T) {
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
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			want: 1,
		},
		{
			name: "Test Case 2",
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
			if got := getMinimumDifferenceV3(tt.args.root); got != tt.want {
				t.Errorf("getMinimumDifferenceV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
