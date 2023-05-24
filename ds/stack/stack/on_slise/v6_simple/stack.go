package v6_simple

type stack []int

// Push - used for push element to stack.
func (s *stack) Push(v int) {
	*s = append(*s, v)
}

// Pop - used for pop element from stack.
func (s *stack) Pop() int {
	if s.Empty() {
		return 0
	}

	result := (*s)[len(*s)-1] // Получаем значение последнего элемента
	*s = (*s)[:len(*s)-1]     // Удаляем последний элемент

	return result
}

// Empty - used for check stack on emptiness.
func (s *stack) Empty() bool {
	return len(*s) == 0
}

// Peak - show last element.
func (s *stack) Peak() int {
	if s.Empty() {
		return 0
	}

	v := (*s)[len(*s)-1]

	return v
}

// Size - used for getting size of stack.
func (s *stack) Size() int {
	return len(*s)
}
