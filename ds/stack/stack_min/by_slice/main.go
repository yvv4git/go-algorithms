package main

import (
	"fmt"
)

// MinStack - структура, представляющая стек с поддержкой операции получения минимального элемента за O(1).
type MinStack struct {
	min   []int // Массив для хранения минимальных значений на каждом уровне стека.
	stack []int // Основной массив для хранения элементов стека.
}

// Constructor - создает и возвращает новый пустой стек.
// Time Complexity: O(1)
// Space Complexity: O(1)
func Constructor() MinStack {
	return MinStack{}
}

// Push - добавляет элемент в стек.
// Time Complexity: O(1) (в худшем случае, из-за добавления элемента в массив)
// Space Complexity: O(1) (дополнительное место для хранения одного элемента)
func (s *MinStack) Push(val int) {
	if len(s.stack) == 0 {
		// Если стек пуст, добавляем первый элемент и устанавливаем его как минимальное значение.
		s.stack = []int{val}
		s.min = []int{val}
	} else {
		// Добавляем элемент в основной стек.
		s.stack = append(s.stack, val)

		// Обновляем массив минимальных значений.
		if val < s.min[len(s.min)-1] {
			// Если текущий элемент меньше текущего минимального значения, добавляем его в массив min.
			s.min = append(s.min, val)
		} else {
			// Иначе, добавляем текущее минимальное значение.
			s.min = append(s.min, s.min[len(s.min)-1])
		}
	}
}

// Pop - удаляет верхний элемент стека.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (s *MinStack) Pop() {
	// Удаляем последний элемент из основного стека.
	s.stack = s.stack[:len(s.stack)-1]
	// Удаляем последний элемент из массива минимальных значений.
	s.min = s.min[:len(s.min)-1]
}

// Top - возвращает значение верхнего элемента стека, не удаляя его.
// Time Complexity: O(1)
// Space Complexity: O(1)
func (s *MinStack) Top() int {
	// Возвращаем последний элемент основного стека.
	return s.stack[len(s.stack)-1]
}

// GetMin - возвращает минимальное значение в стеке за O(1).
// Time Complexity: O(1)
// Space Complexity: O(1)
func (s *MinStack) GetMin() int {
	// Возвращаем последний элемент массива минимальных значений.
	return s.min[len(s.min)-1]
}

// Метод main для демонстрации работы MinStack.
func main() {
	// Создаем новый стек.
	stack := Constructor()

	// Добавляем элементы в стек.
	stack.Push(3)
	stack.Push(5)
	stack.Push(2)
	stack.Push(1)

	// Получаем минимальное значение в стеке.
	fmt.Println("Минимальное значение в стеке:", stack.GetMin()) // Вывод: 1

	// Удаляем верхний элемент стека.
	stack.Pop()

	// Снова получаем минимальное значение.
	fmt.Println("Минимальное значение после Pop:", stack.GetMin()) // Вывод: 2

	// Получаем верхний элемент стека.
	fmt.Println("Верхний элемент стека:", stack.Top()) // Вывод: 2

	// Удаляем еще один элемент.
	stack.Pop()

	// Получаем минимальное значение.
	fmt.Println("Минимальное значение после второго Pop:", stack.GetMin()) // Вывод: 3
}
