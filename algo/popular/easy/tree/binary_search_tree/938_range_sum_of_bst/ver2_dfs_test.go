package main

import "testing"

func Test_rangeSumBSTV2(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
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
					Val: 10,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
					Right: &TreeNode{
						Val: 15,
						Right: &TreeNode{
							Val: 18,
						},
					},
				},
				low:  7,
				high: 15,
			},
			want: 32,
		},
		{
			name: "Test Case 2",
			args: args{
				root: &TreeNode{
					Val: 10,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 1,
							},
						},
						Right: &TreeNode{
							Val: 7,
							Left: &TreeNode{
								Val: 6,
							},
						},
					},
					Right: &TreeNode{
						Val: 15,
						Left: &TreeNode{
							Val: 13,
						},
						Right: &TreeNode{
							Val: 18,
						},
					},
				},
				low:  6,
				high: 10,
			},
			want: 23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBSTV2(tt.args.root, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("rangeSumBSTV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
