package main

//func printList(node *ListNode) {
//	for node != nil {
//		fmt.Print(node.Val, " ")
//		node = node.Next
//	}
//	fmt.Println()
//}

//func main() {
//	/*
//		Крайний случай, появляется dummy node перед удалением 1-ой ноды.
//		Нам концептуально надо дойти до предыдущей вершины, которую хотим удалить и удалить связь или поменять указатель.
//
//		1. Считаем дилну списка с dummy node.
//		2. После того, как нашли len, надо проитерироваться на ...
//		3. Возвращаем список от dummyNode.Next.
//	*/
//	head := &ListNode{Val: 1}
//	head.Next = &ListNode{Val: 2}
//	head.Next.Next = &ListNode{Val: 3}
//	head.Next.Next.Next = &ListNode{Val: 4}
//	head.Next.Next.Next.Next = &ListNode{Val: 5}
//
//	fmt.Print("Original List: ")
//	printList(head)
//
//	head = removeNthFromEnd(head, 2)
//
//	fmt.Print("List after removing 2nd node from end: ")
//	printList(head)
//}
