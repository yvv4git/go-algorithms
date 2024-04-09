package main

import "fmt"

func oddEvenList(head *ListNode) *ListNode {
	/*
		METHOD: Two pointers

		TIME COMPLEXITY: O(n), где n - количество узлов в списке, так как мы проходим по каждому узлу только один раз.

		SPACE COMPLEXITY: O(1), так как мы используем фиксированное количество дополнительной памяти для хранения указателей на начало списков с четными и нечетными числами.
	*/
	if head == nil {
		return nil
	}

	// Инициализируем два указателя на начало списков с четными и нечетными числами
	evenHead := head.Next
	odd, even := head, evenHead

	// Проходим по списку, пока не достигнем конца
	for even != nil && even.Next != nil {
		// Перемещаем узлы в соответствующие списки
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	// Связываем конец списка с четными числами с началом списка с нечетными числами
	odd.Next = evenHead

	return head
}

func main() {
	// Создаем тестовый список
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	// Разделяем список на четные и нечетные числа
	head = oddEvenList(head)

	// Выводим результат
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("NULL")
}
