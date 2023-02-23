package main

/*
2. Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example-1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
		Нужно просуммировать 2 linked list при чем значения складываются в обратном порядке.
	*/
	var out = new(ListNode)
	var current = out
	var mod int
	leftVal, rightVal, sum := 0, 0, 0
	left, right := l1, l2

	for ; left != nil || right != nil || mod != 0; left, right, current = Next(left), Next(right), current.Next {
		leftVal = getVal(left)
		rightVal = getVal(right)
		sum = leftVal + rightVal + mod
		if sum > 9 {
			sum -= 10
			mod = 1
		} else {
			mod &= 0
		}

		current.Val = sum
		if Next(left) != nil || Next(right) != nil || mod != 0 {
			current.Next = new(ListNode)
		}
	}

	return out
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getVal(node *ListNode) int {
	switch node {
	case nil:
		{
			return 0
		}
	default:
		return node.Val
	}
}

func Next(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}
	return nil
}
