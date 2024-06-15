package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Solution - определение структуры Solution для решения задачи
type Solution struct {
	head *ListNode
}

// Constructor - конструктор для инициализации объекта Solution с заданным связным списком
func Constructor(head *ListNode) Solution {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	return Solution{head: head}
}

// GetRandom - метод для получения случайного узла из связного списка
func (this *Solution) GetRandom() int {
	/*
		METHOD: Reservoir Sampling
		Для решения задачи используется алгоритм "резервуарная выборка" (Reservoir Sampling).
		Этот алгоритм позволяет выбирать k элементов из потока данных с равной вероятностью,
		даже если общее количество элементов заранее неизвестно.
		В данном случае k = 1, так как нам нужно выбрать только один случайный узел.

		TIME COMPLEXITY: O(n), где n — количество узлов в списке.

		SPACE COMPLEXITY: O(1), так как мы используем только несколько переменных для хранения текущего узла, счетчика и результата.
	*/
	result := this.head.Val // Инициализация результата значением первого узла
	current := this.head.Next
	count := 1

	// Проход по всем узлам списка
	for current != nil {
		count++
		// Генерация случайного числа в диапазоне от 0 до count-1
		if rand.Intn(count) == 0 {
			result = current.Val // Обновление результата, если сгенерированное число равно 0
		}
		current = current.Next
	}

	return result
}

// Пример использования
func main() {
	// Создание связного списка: 1 -> 2 -> 3
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	solution := Constructor(head)

	// Получение случайного узла
	fmt.Println(solution.GetRandom())
}
