package main

import "fmt"

func sortList(head *ListNode) *ListNode {
	/*
		METHOD: Merge sort

		TIME COMPLEXITY: O(n log n) - это время, необходимое для сортировки списка, где n - количество элементов в списке.

		SPACE COMPLEXITY: O(log n) - это дополнительная память, необходимая для хранения рекурсивных вызовов функции сортировки.
	*/
	// Если список пуст или состоит из одного элемента, то он уже отсортирован
	if head == nil || head.Next == nil {
		return head
	}

	// Находим середину списка
	var prev *ListNode = nil
	slow, fast := head, head

	// Цикл используется для нахождения середины списка.
	// Мы используем два указателя: slow и fast.
	// Slow перемещается на один узел за шаг, а fast - на два.
	// Когда fast достигнет конца списка, slow будет указывать на середину списка.
	for fast != nil && fast.Next != nil {
		// Перед перемещением slow, мы сохраняем его предыдущий узел в prev.
		prev = slow
		// Перемещаем slow на один узел вперед.
		slow = slow.Next
		// Перемещаем fast на два узла вперед.
		fast = fast.Next.Next
	}

	prev.Next = nil

	// Рекурсивно сортируем два подсписка
	left := sortList(head)
	right := sortList(slow)

	// Сливаем два отсортированных подсписка
	return merge(left, right)
}

// Функция merge объединяет два отсортированных списка в один отсортированный список.
func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	// Создаем фиктивный узел, который будет началом нового списка.
	dummy := &ListNode{}
	current := dummy

	// Проходимся по обоим спискам и выбираем меньший из текущих элементов.
	// Добавляем его в новый список и переходим к следующему элементу в выбранном списке.
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// Если один из списков закончился раньше, то добавляем оставшиеся элементы второго списка.
	if l1 != nil {
		current.Next = l1
	}

	if l2 != nil {
		current.Next = l2
	}

	// Возвращаем начало нового списка.
	return dummy.Next
}

func main() {
	// Создаем тестовый список
	head := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}

	// Сортируем список
	head = sortList(head)

	// Выводим отсортированный список
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}
