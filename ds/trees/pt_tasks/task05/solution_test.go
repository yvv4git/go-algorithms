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

func TestTreeNode_IsFork(t *testing.T) {
	t.Run("nil node is not a fork", func(t *testing.T) {
		var node *TreeNode
		if node.IsFork() {
			t.Error("Expected nil node to not be a fork")
		}
	})

	t.Run("node with no children is not a fork", func(t *testing.T) {
		node := &TreeNode{Value: 5}
		if node.IsFork() {
			t.Error("Expected node with no children to not be a fork")
		}
	})

	t.Run("node with only left child is not a fork", func(t *testing.T) {
		node := &TreeNode{Value: 5, Left: &TreeNode{Value: 3}}
		if node.IsFork() {
			t.Error("Expected node with only left child to not be a fork")
		}
	})

	t.Run("node with only right child is not a fork", func(t *testing.T) {
		node := &TreeNode{Value: 5, Right: &TreeNode{Value: 7}}
		if node.IsFork() {
			t.Error("Expected node with only right child to not be a fork")
		}
	})

	t.Run("node with both children is a fork", func(t *testing.T) {
		node := &TreeNode{
			Value: 5,
			Left:  &TreeNode{Value: 3},
			Right: &TreeNode{Value: 7},
		}
		if !node.IsFork() {
			t.Error("Expected node with both children to be a fork")
		}
	})
}

func TestTreeNode_CollectForks(t *testing.T) {
	t.Run("empty tree", func(t *testing.T) {
		var root *TreeNode
		var result []int
		root.CollectForks(&result)
		if len(result) != 0 {
			t.Errorf("Expected empty result, got %v", result)
		}
	})

	t.Run("single node", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		var result []int
		root.CollectForks(&result)
		if len(result) != 0 {
			t.Errorf("Expected 0 elements (single node is not a fork), got %d", len(result))
		}
	})

	t.Run("left skewed tree", func(t *testing.T) {
		root := &TreeNode{Value: 5}
		root.Insert(4)
		root.Insert(3)
		root.Insert(2)
		root.Insert(1)

		var result []int
		root.CollectForks(&result)

		// In a left-skewed tree, no node has both children.
		expected := []int{}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements, got %d", len(expected), len(result))
		}
	})

	t.Run("right skewed tree", func(t *testing.T) {
		root := &TreeNode{Value: 1}
		root.Insert(2)
		root.Insert(3)
		root.Insert(4)
		root.Insert(5)

		var result []int
		root.CollectForks(&result)

		// In a right-skewed tree, no node has both children.
		expected := []int{}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements, got %d", len(expected), len(result))
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
		root.CollectForks(&result)

		// In a balanced tree, forks: 3, 5, 7.
		expected := []int{3, 5, 7}
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
		root.CollectForks(&result)

		// Forks of the tree from the example: 3, 5, 7.
		expected := []int{3, 5, 7}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements, got %d", len(expected), len(result))
		}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Position %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("tree with only one fork", func(t *testing.T) {
		// Tree:
		//       5
		//      / \
		//     3   7
		root := &TreeNode{Value: 5}
		root.Insert(3)
		root.Insert(7)

		var result []int
		root.CollectForks(&result)

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

	t.Run("tree with duplicates", func(t *testing.T) {
		var root *TreeNode
		values := []int{5, 3, 7, 3, 5, 7, 2, 8}
		for _, v := range values {
			root = root.Insert(v)
		}

		var result []int
		root.CollectForks(&result)

		// Duplicates are not added.
		// Tree structure:
		//     5
		//    / \
		//   3   7
		//  /     \
		// 2       8
		// Only node 5 has both children.
		expected := []int{5}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements (duplicates skipped), got %d", len(expected), len(result))
		}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Position %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})

	t.Run("complex tree", func(t *testing.T) {
		// Tree:
		//         10
		//        /  \
		//       5    15
		//      / \   / \
		//     3   7 12  20
		//    / \       /  \
		//   1   4    18   25
		var root *TreeNode
		values := []int{10, 5, 15, 3, 7, 12, 20, 1, 4, 18, 25}
		for _, v := range values {
			root = root.Insert(v)
		}

		var result []int
		root.CollectForks(&result)

		// Forks: 3, 5, 10, 15, 20.
		expected := []int{3, 5, 10, 15, 20}
		if len(result) != len(expected) {
			t.Errorf("Expected %d elements, got %d", len(expected), len(result))
		}
		for i, v := range result {
			if v != expected[i] {
				t.Errorf("Position %d: expected %d, got %d", i, expected[i], v)
			}
		}
	})
}
