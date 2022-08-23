package main

// Stack - реализация стека, которая может хранить обобщенные типы. Т.е. значения любого типа.
// Но при этом мы не используем generics, а используем интерфейс.
type Stack struct {
	values []interface{}
}

func (s *Stack) Push(val interface{}) {
	s.values = append(s.values, val)
}
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.values) == 0 {
		return nil, false
	}

	top := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return top, true
}
