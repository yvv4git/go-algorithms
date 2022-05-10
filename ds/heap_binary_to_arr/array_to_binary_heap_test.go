package main

import (
	"reflect"
	"testing"
)

func Test_arrayToBinaryHeap(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name     string
		args     args
		wantRoot *TreeNode
	}{
		{
			name: "CASE-1",
			args: args{
				arr: []int{1, 0, 2},
			},
			wantRoot: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			name: "CASE-2",
			args: args{
				arr: []int{3, 9, 20, 0, 0, 15, 7},
			},
			wantRoot: &TreeNode{
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRoot := arrayToBinaryHeap(tt.args.arr); !reflect.DeepEqual(gotRoot, tt.wantRoot) {
				t.Errorf("arrayToBinaryHeap() = %v, want %v", gotRoot, tt.wantRoot)
			}
		})
	}
}
