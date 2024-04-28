package main

import "testing"

func Test_isSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1: Subtree is a leaf node",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 0,
							},
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				subRoot: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: false, // Поддерево subRoot не является поддеревом дерева root, так как узел 0 в subRoot не является дочерним узлом в root.
		},
		{
			name: "Test Case 2: Subtree is a subtree of root",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				subRoot: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: true, // Поддерево subRoot является поддеревом дерева root.
		},
		{
			name: "Test Case 3: Subtree is not a subtree of root",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 1,
						},
						Right: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 0,
							},
						},
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				subRoot: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3, // Значение должно быть 0, чтобы поддерево subRoot было поддеревом дерева root.
						},
					},
				},
			},
			want: false, // Поддерево subRoot не является поддеревом дерева root, так как узел 3 в subRoot не является дочерним узлом в root.
		},
		// Добавьте свои тестовые кейсы здесь.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
