package main

// Stack - реализация stack с использованием generics.
type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(val T) {
	s.values = append(s.values, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.values) == 0 {
		var zero T
		return zero, false
	}

	top := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	return top, true
}
