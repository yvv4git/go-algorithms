package main

import "testing"

func Test_amountOfTimeV3(t *testing.T) {
	type args struct {
		root  *TreeNode
		start int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one node tree",
			args: args{
				root:  &TreeNode{Val: 1},
				start: 1,
			},
			want: 0,
		},
		{
			name: "case-2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 9,
							},
							Right: &TreeNode{
								Val: 2,
							},
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 10,
						},
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
				start: 3,
			},
			want: 4,
		},
		{
			name: "case-3",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 4,
							},
							Right: &TreeNode{
								Val: 5,
							},
						},
					},
					Right: &TreeNode{
						Val: 6,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 8,
						},
					},
				},
				start: 1,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := amountOfTimeV3(tt.args.root, tt.args.start); got != tt.want {
				t.Errorf("amountOfTimeV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
