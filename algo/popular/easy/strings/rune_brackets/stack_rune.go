package rune_brackets

import (
	"container/list"
)

// Stack - common contract.
type Stack interface {
	Push(rune)
	Pop() rune
	Size() int
}

type RuneStack struct {
	stack *list.List
}

func NewStack() *RuneStack {
	return &RuneStack{
		stack: list.New(),
	}
}

func (s *RuneStack) Push(value rune) {
	s.stack.PushFront(value)
}

func (s *RuneStack) Pop() rune {
	if s.stack.Len() > 0 {
		element := s.stack.Front()
		value, ok := element.Value.(rune)
		if !ok {
			return 0
		}
		s.stack.Remove(element)

		return value
	}

	return 0
}

func (s *RuneStack) Size() int {
	return s.stack.Len()
}

// Front - позволяет посмотреть элемент на верхушке стека.
func (s *RuneStack) Front() rune {
	if s.stack.Len() > 0 {
		if value, ok := s.stack.Front().Value.(rune); ok {
			return value
		}

		return 0
	}

	return 0
}
