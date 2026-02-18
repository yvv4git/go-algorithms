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

func TestTreeNode_IsLeaf(t *testing.T) {
	t.Run("nil node is not a leaf", func(t *testing.T) {
		var node *TreeNode
		if node.IsLeaf() {
			t.Error("Expected nil node to not be a leaf")
		}
	})

	t.Run("node with no children is a leaf", func(t *testing.T) {
		node := &TreeNode{Value: 5}
		if !node.IsLeaf() {
			t.Error("Expected node with no children to be a leaf")
		}
	})

	t.Run("node with left child is not a leaf", func(t *testing.T) {
		node := &TreeNode{Value: 5, Left: &TreeNode{Value: 3}}
		if node.IsLeaf() {
			t.Error("Expected node with left child to not be a leaf")
		}
	})

	t.Run("node with right child is not a leaf", func(t *testing.T) {
		node := &TreeNode{Value: 5, Right: &TreeNode{Value: 7}}
		if node.IsLeaf() {
			t.Error("Expected node with right child to not be a leaf")
		}
	})

	t.Run("node with both children is not a leaf", func(t *testing.T) {
		node := &TreeNode{
			Value: 5,
			Left:  &TreeNode{Value: 3},
			Right: &TreeNode{Value: 7},
		}
		if node.IsLeaf() {
			t.Error("Expected node with both children to not be a leaf")
		}
	})
}

func TestTreeNode_CollectLeaves(t *testing.T) {
	t.Run("empty tree", func(t *testing.T) {
		var root *TreeNode
		var result []int
		root.CollectLeaves(&result)
		if len(result) != 0 {
			t.Errorf("Expected empty result, got %v", result)
		}
	})

	t.Run("single node", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		var result []int
		root.CollectLeaves(&result)
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
		root.CollectLeaves(&result)

		// В лево-перекошенном дереве только самый левый узел будет листом.
		expected := []int{1}
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
		root.CollectLeaves(&result)

		// В право-перекошенном дереве только самый правый узел будет листом.
		expected := []int{5}
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
		root.CollectLeaves(&result)

		// В сбалансированном дереве листья: 2, 4, 6, 8.
		expected := []int{2, 4, 6, 8}
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
		root.CollectLeaves(&result)

		// Листья дерева из примера: 1, 4, 6, 8.
		expected := []int{1, 4, 6, 8}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements, got %d", len(expected), len(result))
		}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Position %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("tree with only leaves at same level", func(t *testing.T) {
		// Дерево:
		//       5
		//      / \
		//     3   7
		//    / \ / \
		//   1  4 6  8
		root := &TreeNode{Value: 5}
		root.Insert(3)
		root.Insert(7)
		root.Insert(1)
		root.Insert(4)
		root.Insert(6)
		root.Insert(8)

		var result []int
		root.CollectLeaves(&result)

		expected := []int{1, 4, 6, 8}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements, got %d", len(expected), len(result))
		}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Position %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("tree with duplicates", func(t *testing.T) {
		var root *TreeNode
		values := []int{5, 3, 7, 3, 5, 7, 2, 8}
		for _, v := range values {
			root = root.Insert(v)
		}

		var result []int
		root.CollectLeaves(&result)

		// Дубликаты не добавляются, листья: 2, 8.
		expected := []int{2, 8}
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
