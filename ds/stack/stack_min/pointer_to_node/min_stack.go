package main

// MinStack - структура, представляющая стек с поддержкой операции получения минимального элемента за O(1).
type MinStack struct {
	head *StackNode // Указатель на вершину стека.
}

// Constructor - создает и возвращает новый пустой стек.
func Constructor() MinStack {
	return MinStack{
		head: nil, // Изначально стек пуст, поэтому head равен nil.
	}
}

// Push - добавляет элемент в стек.
func (s *MinStack) Push(val int) {
	if s.head == nil {
		// Если стек пуст, создаем первый узел.
		// Минимальное значение в этом случае равно самому значению, так как это единственный элемент.
		s.head = NewStackNode(val, val, nil)
	} else {
		// Если стек не пуст, определяем минимальное значение на момент добавления нового элемента.
		min := val // Предполагаем, что минимальное значение — это текущее значение.
		if s.head.min < min {
			// Если минимальное значение в текущей вершине стека меньше, чем текущее значение,
			// то минимальное значение остается таким же, как в текущей вершине.
			min = s.head.min
		}
		// Создаем новый узел, который станет новой вершиной стека.
		// Новый узел ссылается на текущую вершину стека (this.head).
		s.head = NewStackNode(
			val,    // Текущее значение.
			min,    // Минимальное значение в стеке на момент добавления этого узла.
			s.head, // Указатель на предыдущую вершину стека.
		)
	}
}

// Pop - удаляет верхний элемент стека.
func (s *MinStack) Pop() {
	// Просто перемещаем указатель head на следующий узел.
	// Узел, на который раньше указывал head, автоматически удаляется сборщиком мусора.
	s.head = s.head.next
}

// Top - возвращает значение верхнего элемента стека, не удаляя его.
func (s *MinStack) Top() int {
	// Возвращаем значение текущей вершины стека.
	return s.head.val
}

// GetMin - возвращает минимальное значение в стеке за O(1).
func (s *MinStack) GetMin() int {
	// Минимальное значение всегда хранится в вершине стека.
	return s.head.min
}

// StackNode - структура, представляющая узел стека.
type StackNode struct {
	val  int        // Значение текущего элемента.
	min  int        // Минимальное значение в стеке на момент добавления этого узла.
	next *StackNode // Указатель на следующий узел (предыдущий элемент в стеке).
}

// NewStackNode - фабричная функция для создания нового узла стека.
func NewStackNode(val, min int, next *StackNode) *StackNode {
	return &StackNode{
		val:  val,  // Значение текущего элемента.
		min:  min,  // Минимальное значение в стеке на момент добавления этого узла.
		next: next, // Указатель на следующий узел (предыдущий элемент в стеке).
	}
}
