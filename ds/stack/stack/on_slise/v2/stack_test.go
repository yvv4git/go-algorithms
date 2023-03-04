package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStackOnSlice(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	stack := StackSlice{}
	for i := range data {
		stack.Push(data[i])
	}
	require.Equal(t, 5, stack.Size(), "error on check size of stack after push")

	for i, j := 0, stack.Size(); i < 5; i, j = i+1, j-1 {
		val := stack.Pop()
		require.Equal(t, j, val, "error on check pop value")
	}

	require.Nil(t, stack.Pop())
}
