package increasing_order_search_tree

import "testing"

func TestIncreasingOrder(t *testing.T) {
	node1 := TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		},
	}

	// Теперь выведем все ноды как linked list.
	result := increasingBST(&node1)
	currentNode := result
	for {
		if currentNode == nil {
			break
		}

		t.Logf("Node val: %d", (*currentNode).Val)
		currentNode = currentNode.Right
	}
}
