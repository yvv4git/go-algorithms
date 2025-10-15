package dijkstra

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHeap_Push(t *testing.T) {
	h := NewHeap()

	originalPath := Path{
		weight: 3.1,
		nodes: []string{
			"BTC",
		},
	}

	h.Push(originalPath)
	resultPath := h.Pop()

	require.Equalf(t, originalPath, resultPath, "the values must be the same")
}
