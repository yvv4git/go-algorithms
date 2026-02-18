package main

import "testing"

func TestTreeNode_InsertDepth(t *testing.T) {
	t.Run("insert into empty tree returns false", func(t *testing.T) {
		var root *TreeNode
		depth, inserted := root.InsertDepth(5)
		if inserted {
			t.Error("Expected inserted = false for nil tree")
		}
		if depth != 0 {
			t.Errorf("Expected depth 0, got %d", depth)
		}
	})

	t.Run("insert smaller value", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		depth, inserted := root.InsertDepth(3)
		if !inserted {
			t.Error("Expected inserted = true")
		}
		if depth != 2 {
			t.Errorf("Expected depth 2, got %d", depth)
		}
		if root.Left == nil {
			t.Error("Expected left child to be non-nil")
		}
		if root.Left.Value != 3 {
			t.Errorf("Expected left.Value = 3, got %d", root.Left.Value)
		}
	})

	t.Run("insert larger value", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		depth, inserted := root.InsertDepth(7)
		if !inserted {
			t.Error("Expected inserted = true")
		}
		if depth != 2 {
			t.Errorf("Expected depth 2, got %d", depth)
		}
		if root.Right == nil {
			t.Error("Expected right child to be non-nil")
		}
		if root.Right.Value != 7 {
			t.Errorf("Expected right.Value = 7, got %d", root.Right.Value)
		}
	})

	t.Run("insert duplicate value", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		depth, inserted := root.InsertDepth(5)
		if inserted {
			t.Error("Expected inserted = false for duplicate")
		}
		if depth != 0 {
			t.Errorf("Expected depth 0 for duplicate, got %d", depth)
		}
		if root.Left != nil || root.Right != nil {
			t.Error("Expected no children for duplicate insert")
		}
	})

	t.Run("insert multiple levels", func(t *testing.T) {
		root := &TreeNode{Value: 5}

		depth, _ := root.InsertDepth(3)
		if depth != 2 {
			t.Errorf("Expected depth 2 for 3, got %d", depth)
		}

		depth, _ = root.InsertDepth(2)
		if depth != 3 {
			t.Errorf("Expected depth 3 for 2, got %d", depth)
		}

		depth, _ = root.InsertDepth(1)
		if depth != 4 {
			t.Errorf("Expected depth 4 for 1, got %d", depth)
		}
	})
}

func TestInsertDepthSequence(t *testing.T) {
	t.Run("example from task", func(t *testing.T) {
		values := []int{7, 3, 2, 1, 9, 5, 4, 6, 8}
		expected := []int{1, 2, 3, 4, 2, 3, 4, 4, 3}

		var root *TreeNode
		var depths []int

		for _, v := range values {
			if root == nil {
				root = &TreeNode{Value: v}
				depths = append(depths, 1)
			} else {
				if depth, inserted := root.InsertDepth(v); inserted {
					depths = append(depths, depth)
				}
			}
		}

		if len(depths) != len(expected) {
			t.Errorf("Expected %d depths, got %d", len(expected), len(depths))
		}

		for i, got := range depths {
			if got != expected[i] {
				t.Errorf("Position %d: expected depth %d, got %d", i, expected[i], got)
			}
		}
	})

	t.Run("left skewed tree", func(t *testing.T) {
		values := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}

		var root *TreeNode
		var depths []int

		for _, v := range values {
			if root == nil {
				root = &TreeNode{Value: v}
				depths = append(depths, 1)
			} else {
				if depth, inserted := root.InsertDepth(v); inserted {
					depths = append(depths, depth)
				}
			}
		}

		for i, got := range depths {
			if got != expected[i] {
				t.Errorf("Position %d: expected depth %d, got %d", i, expected[i], got)
			}
		}
	})

	t.Run("right skewed tree", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}

		var root *TreeNode
		var depths []int

		for _, v := range values {
			if root == nil {
				root = &TreeNode{Value: v}
				depths = append(depths, 1)
			} else {
				if depth, inserted := root.InsertDepth(v); inserted {
					depths = append(depths, depth)
				}
			}
		}

		for i, got := range depths {
			if got != expected[i] {
				t.Errorf("Position %d: expected depth %d, got %d", i, expected[i], got)
			}
		}
	})

	t.Run("with duplicates", func(t *testing.T) {
		values := []int{5, 3, 7, 3, 5, 7, 2}
		expected := []int{1, 2, 2, 3}

		var root *TreeNode
		var depths []int

		for _, v := range values {
			if root == nil {
				root = &TreeNode{Value: v}
				depths = append(depths, 1)
			} else {
				if depth, inserted := root.InsertDepth(v); inserted {
					depths = append(depths, depth)
				}
			}
		}

		if len(depths) != len(expected) {
			t.Errorf("Expected %d depths (duplicates skipped), got %d", len(expected), len(depths))
		}

		for i, got := range depths {
			if got != expected[i] {
				t.Errorf("Position %d: expected depth %d, got %d", i, expected[i], got)
			}
		}
	})

	t.Run("single element", func(t *testing.T) {
		values := []int{42}
		expected := []int{1}

		var root *TreeNode
		var depths []int

		for _, v := range values {
			if root == nil {
				root = &TreeNode{Value: v}
				depths = append(depths, 1)
			} else {
				if depth, inserted := root.InsertDepth(v); inserted {
					depths = append(depths, depth)
				}
			}
		}

		if len(depths) != len(expected) {
			t.Errorf("Expected %d depths, got %d", len(expected), len(depths))
		}

		if depths[0] != expected[0] {
			t.Errorf("Expected depth %d, got %d", expected[0], depths[0])
		}
	})
}
