package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	assert.Equal(t, 1, stack.Pop())
}

func TestStackEmptyValue(t *testing.T) {
	stack := NewStack()
	assert.Equal(t, -1, stack.Pop())
}

func TestStackOrder(t *testing.T) {
	payload := []int{1, 2, 3, 4, 5}

	stack := NewStack()
	for _, val := range payload {
		stack.Push(val)
	}

	var result []int
	for i := 0; i < len(payload); i++ {
		result = append(result, stack.Pop())
	}

	// Проверим, что значения достали в обратном порядке.
	assert.Equal(t, []int{5, 4, 3, 2, 1}, result)
}
