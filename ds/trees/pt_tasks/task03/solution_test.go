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

func TestTreeNode_InOrderTraversal(t *testing.T) {
	t.Run("empty tree", func(t *testing.T) {
		var root *TreeNode
		var result []int
		root.InOrderTraversal(&result)
		if len(result) != 0 {
			t.Errorf("Expected empty result, got %v", result)
		}
	})

	t.Run("single node", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		var result []int
		root.InOrderTraversal(&result)
		if len(result) != 1 {
			t.Errorf("Expected 1 element, got %d", len(result))
		}
		if result[0] != 5 {
			t.Errorf("Expected [5], got %v", result)
		}
	})

	t.Run("left skewed tree", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		root.Insert(4)
		root.Insert(3)
		root.Insert(2)
		root.Insert(1)

		var result []int
		root.InOrderTraversal(&result)

		expected := []int{1, 2, 3, 4, 5}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements, got %d", len(expected), len(result))
		}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Position %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("right skewed tree", func(t *testing.T) {
		root := &TreeNode{Value: 1}
		root.Insert(2)
		root.Insert(3)
		root.Insert(4)
		root.Insert(5)

		var result []int
		root.InOrderTraversal(&result)

		expected := []int{1, 2, 3, 4, 5}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements, got %d", len(expected), len(result))
		}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Position %d: expected %d, got %d", i, expected[i], v)
			}
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

		var result []int
		root.InOrderTraversal(&result)

		expected := []int{2, 3, 4, 5, 6, 7, 8}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements, got %d", len(expected), len(result))
		}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Position %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("example from task", func(t *testing.T) {
		var root *TreeNode
		values := []int{7, 3, 2, 1, 9, 5, 4, 6, 8}
		for _, v := range values {
			root = root.Insert(v)
		}

		var result []int
		root.InOrderTraversal(&result)

		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements, got %d", len(expected), len(result))
		}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Position %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("with duplicates", func(t *testing.T) {
		var root *TreeNode
		values := []int{5, 3, 7, 3, 5, 7, 2}
		for _, v := range values {
			root = root.Insert(v)
		}

		var result []int
		root.InOrderTraversal(&result)

		// Duplicates should not be added
		expected := []int{2, 3, 5, 7}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements (duplicates skipped), got %d", len(expected), len(result))
		}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Position %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})
}
