package _2_reverse_linked_list_II

func reverseBetweenV3(head *ListNode, left int, right int) *ListNode {
	/*
		METHOD: Recursion.
		TIME COMPLEXITY: O(n).
		SPACE COMPLEXITY: O(n).

		Time complexity - O(n):
		В начале функции мы проходим по всему списку до узла left, что занимает O(n) времени.
		Затем мы проходим по всему списку от узла left до узла right, добавляя узлы в стек, что также занимает O(n) времени.
		После этого мы проходим по стеку, извлекая узлы и добавляя их в новый список, что также занимает O(n) времени.

		Space complexity - O(n):
		В этом алгоритме мы используем дополнительное пространство для хранения узлов в стеке.
		Стек может содержать до n элементов в худшем случае, когда все узлы находятся в диапазоне [left, right].
		Таким образом, в худшем случае, когда left и right близки к концу списка,
		пространственная сложность алгоритма будет пропорциональна количеству узлов в списке.
	*/

	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	stack := []*ListNode{}
	current := prev.Next

	for i := 0; i < right-left+1; i++ {
		stack = append(stack, current)
		current = current.Next
	}

	for len(stack) > 0 {
		last := len(stack) - 1
		prev.Next = stack[last]
		stack = stack[:last]
		prev = prev.Next
	}

	prev.Next = current

	return dummy.Next
}
