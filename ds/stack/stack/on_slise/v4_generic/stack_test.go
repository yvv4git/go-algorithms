package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStackGenericsInt(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	stack := Stack[int]{}

	for i := range data {
		stack.Push(data[i])
	}
	require.Equal(t, 5, stack.Size(), "error on check push operation")

	for i, j := 0, len(data); i < 5; i, j = i+1, j-1 {
		val, ok := stack.Pop()
		require.True(t, ok, "ok must bee true")
		require.Equal(t, j, val, "error on check pop operation")
	}

	val, ok := stack.Pop()
	require.False(t, ok, "ok mast bee false")
	require.Equal(t, 0, val)
}
