package rune_brackets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackSimple(t *testing.T) {
	stack := NewStack()

	stack.Push('A')
	assert.Equal(t, 'A', stack.Pop())
}

func TestStackFront(t *testing.T) {
	stack := NewStack()

	stack.Push('A')
	assert.Equal(t, 'A', stack.Front())
	assert.Equal(t, 1, stack.Size())
	assert.Equal(t, 'A', stack.Pop())
	assert.Equal(t, 0, stack.Size())
}

func TestStackMultipleValues(t *testing.T) {
	srcValues := []rune{'a', 'b', 'c', 'd', 'e'}
	stack := NewStack()

	for _, value := range srcValues {
		stack.Push(value)
	}

	var result []rune
	for {
		value := stack.Pop()
		if value == 0 {
			break
		}
		result = append(result, value)
	}

	assert.Equal(t, []rune{'e', 'd', 'c', 'b', 'a'}, result)
}
