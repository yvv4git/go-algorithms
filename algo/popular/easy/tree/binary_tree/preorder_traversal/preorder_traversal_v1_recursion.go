package preorder_traversal

type Res struct {
	Res []int
}

func preorderTraversalV1(root *TreeNode) []int {
	/*
		Method: Recursion.
		Time complexity : O(n)
		Space complexity : O(1)
	*/
	r := Res{}
	r.traversal(root)
	return r.Res
}

func (r *Res) traversal(node *TreeNode) {
	if node != nil {
		r.Res = append(r.Res, node.Val)
		r.traversal(node.Left)
		r.traversal(node.Right)
	}
}
