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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{}
	dummy.Next = nil
	tail := &dummy

	for l1 != nil || l2 != nil {
		if l1 == nil {
			(*tail).Next = l2
			break
		} else if l2 == nil {
			(*tail).Next = l1
			break
		} else {
			if l1.Val <= l2.Val {
				(*tail).Next = l1
				l1 = (*l1).Next
			} else {
				(*tail).Next = l2
				l2 = (*l2).Next
			}
			tail = tail.Next
		}
	}

	return dummy.Next
}
