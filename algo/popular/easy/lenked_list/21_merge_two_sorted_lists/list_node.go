package main

// ListNode - used as node
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode - used as simplefactory for create instance of ListNode
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

// AddNodeWithValue - used for add next node
func (n *ListNode) AddNodeWithValue(val int) *ListNode {
	n.Next = NewListNode(val)
	return n.Next
}
