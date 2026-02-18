package main

import "testing"

func TestTreeNode_Insert(t *testing.T) {
	t.Run("insert into empty tree", func(t *testing.T) {
		var root *TreeNode
		root = root.Insert(5)
		if root == nil {
			t.Error("Expected root to be non-nil")
		}
		if root.Value != 5 {
			t.Errorf("Expected root.Value = 5, got %d", root.Value)
		}
	})

	t.Run("insert smaller value", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		root.Insert(3)
		if root.Left == nil {
			t.Error("Expected left child to be non-nil")
		}
		if root.Left.Value != 3 {
			t.Errorf("Expected left.Value = 3, got %d", root.Left.Value)
		}
	})

	t.Run("insert larger value", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		root.Insert(7)
		if root.Right == nil {
			t.Error("Expected right child to be non-nil")
		}
		if root.Right.Value != 7 {
			t.Errorf("Expected right.Value = 7, got %d", root.Right.Value)
		}
	})

	t.Run("insert duplicate value", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		root.Insert(5)
		if root.Left != nil || root.Right != nil {
			t.Error("Expected no children for duplicate insert")
		}
	})
}

func TestTreeNode_Height(t *testing.T) {
	t.Run("empty tree height", func(t *testing.T) {
		var root *TreeNode
		if got := root.Height(); got != 0 {
			t.Errorf("Expected height 0, got %d", got)
		}
	})

	t.Run("single node height", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		if got := root.Height(); got != 1 {
			t.Errorf("Expected height 1, got %d", got)
		}
	})

	t.Run("left skewed tree", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		root.Insert(4)
		root.Insert(3)
		root.Insert(2)
		root.Insert(1)
		if got := root.Height(); got != 5 {
			t.Errorf("Expected height 5, got %d", got)
		}
	})

	t.Run("right skewed tree", func(t *testing.T) {
		root := &TreeNode{Value: 1}
		root.Insert(2)
		root.Insert(3)
		root.Insert(4)
		root.Insert(5)
		if got := root.Height(); got != 5 {
			t.Errorf("Expected height 5, got %d", got)
		}
	})

	t.Run("balanced tree", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		root.Insert(3)
		root.Insert(7)
		root.Insert(2)
		root.Insert(4)
		root.Insert(6)
		root.Insert(8)
		if got := root.Height(); got != 3 {
			t.Errorf("Expected height 3, got %d", got)
		}
	})

	t.Run("example from task", func(t *testing.T) {
		var root *TreeNode
		values := []int{7, 3, 2, 1, 9, 5, 4, 6, 8}
		for _, v := range values {
			root = root.Insert(v)
		}
		if got := root.Height(); got != 4 {
			t.Errorf("Expected height 4, got %d", got)
		}
	})
}
