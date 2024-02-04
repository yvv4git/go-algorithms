package main

func isPalindromeV4(head *ListNode) bool {
	/*
		METHOD: Recursion
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity:
		Временная сложность для всех этих решений будет O(n), где n - количество узлов в списке.
		Это связано с тем, что мы проходим по всем узлам списка только один раз.

		Space complexity.
		Пространственная сложность для решения с использованием стека или очереди будет O(n),
		так как в худшем случае мы можем хранить в стеке или очереди все n узлов списка.

	*/
	frontPointer := head
	var recursivelyCheck func(*ListNode) bool
	recursivelyCheck = func(curNode *ListNode) bool {
		if curNode != nil {
			if !recursivelyCheck(curNode.Next) {
				return false
			}
			if curNode.Val != frontPointer.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}
	return recursivelyCheck(head)
}
