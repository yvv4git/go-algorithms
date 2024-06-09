package main

import "fmt"

// insertionSortList - функция для сортировки вставками связного списка
func insertionSortList(head *ListNode) *ListNode {
	/*
		METHOD: Insertion sort & Dummy node
		Используемый подход здесь - это сортировка вставками, которая позволяет нам создать отсортированный список,
		проходя по нему и вставляя каждый элемент в правильное место.

		TIME COMPLEXITY: O(n^2), где n - количество узлов в списке, потому что в худшем случае мы можем пройти по всему списку для каждого узла.

		SPACE COMPLEXITY: O(1), так как мы используем фиктивный узел и не выделяем дополнительно памяти, кроме переменных для указателей.
	*/
	// Создаем фиктивный узел, который будет началом отсортированного списка
	dummy := &ListNode{Val: 0}
	curr := head

	// Проходим по всему списку
	for curr != nil {
		// Сохраняем ссылку на следующий узел, так как после вставки текущего узла он будет удален из списка
		next := curr.Next
		// Ищем место для вставки текущего узла в отсортированную часть списка
		prev := dummy
		for prev.Next != nil && prev.Next.Val < curr.Val {
			prev = prev.Next
		}
		// Вставляем текущий узел в отсортированную часть списка
		curr.Next = prev.Next
		prev.Next = curr
		// Переходим к следующему узлу
		curr = next
	}

	// Возвращаем начало отсортированного списка
	return dummy.Next
}

func main() {
	// Пример использования
	head := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
	sortedList := insertionSortList(head)
	for sortedList != nil {
		fmt.Println(sortedList.Val)
		sortedList = sortedList.Next
	}
}
