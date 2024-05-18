package main

import "testing"

func Test_findTiltV2(t *testing.T) {
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
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: 1,
		},
		{
			name: "Test Case 2",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 9,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: 15,
		},
		{
			name: "Test Case 3",
			args: args{
				root: nil,
			},
			want: 0,
		},
		{
			name: "Test Case 4",
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
			if got := findTiltV2(tt.args.root); got != tt.want {
				t.Errorf("findTiltV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
