package array_to_binary_tree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}

// TreeNode - Tree node implementation
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
