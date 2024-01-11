package main

import "math"

func isValidBSTV3(root *TreeNode) bool {
	var stack Stack

	cn_root := CNode{Node: root, Min: math.MinInt, Max: math.MaxInt}
	stack.Push(cn_root)

	for !stack.IsEmpty() {
		next := stack.Pop()

		if !next.Check() {
			return false
		} else {
			if next.Node.Left != nil {
				left := CNode{Node: next.Node.Left, Min: next.Min, Max: next.Node.Val - 1}
				stack.Push(left)
			}

			if next.Node.Right != nil {
				right := CNode{Node: next.Node.Right, Min: next.Node.Val + 1, Max: next.Max}
				stack.Push(right)
			}
		}
	}

	return true
}

// CNODE
type CNode struct {
	Node *TreeNode
	Max  int
	Min  int
}

func (cn *CNode) Check() bool {
	return (cn.Node.Val >= cn.Min && cn.Node.Val <= cn.Max) && (cn.Node.Val != math.MinInt || cn.Node.Left == nil) && (cn.Node.Val != math.MaxInt || cn.Node.Right == nil)
}

// STACK
type Stack []CNode

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(en CNode) {
	*s = append(*s, en)
}

func (s *Stack) Pop() CNode {
	node := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return node
}
