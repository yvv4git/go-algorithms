package main

import "container/list"

import (
	"errors"
)

const (
	error_empty           = "Stack is empty"
	error_type_not_string = "Stack is empty"
)

// listStack - реализация стека на основе списка.
type listStack struct {
	stack *list.List
}

// Push - т.е. добавить элемент в стек.
func (s *listStack) Push(value string) {
	s.stack.PushFront(value)
}

// Pop - достать элемент.
func (s *listStack) Pop() (result string, err error) {
	if s.stack.Len() > 0 {
		element := s.stack.Front()
		value, ok := element.Value.(string)
		if !ok {
			return "", errors.New(error_type_not_string)
		}
		s.stack.Remove(element)
		return value, nil
	}

	return "", errors.New(error_empty)
}

// Front - можно посмотреть значение первого элемента.
func (s *listStack) Front() (result string, err error) {
	if s.stack.Len() > 0 {
		if value, ok := s.stack.Front().Value.(string); ok {
			return value, nil
		}
		return "", errors.New(error_type_not_string)
	}

	return "", errors.New(error_empty)
}

// Size - возвращает размер стека.
func (s *listStack) Size() int {
	return s.stack.Len()
}

// Empty - определяет, пустой ли стек.
func (s *listStack) Empty() bool {
	return s.stack.Len() == 0
}
