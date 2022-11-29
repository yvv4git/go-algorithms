package dijkstra

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinPath(t *testing.T) {
	minPath := &MinPath{}

	minPathLen := minPath.Len()
	require.Equalf(t, minPathLen, 0, "at first the path is empty")

	// Check push method.
	minPath.Push(NewPath(3.0, []string{"BTC"}))
	require.Equalf(t, minPath.Len(), 1, "added first node")

	minPath.Push(NewPath(5.2, []string{"USD"}))
	require.Equalf(t, minPath.Len(), 2, "added second node")

	minPath.Push(NewPath(4.89, []string{"EUR"}))
	require.Equalf(t, minPath.Len(), 3, "added third node")

	// Check pop method.
	require.Equalf(
		t,
		NewPath(4.89, []string{"EUR"}),
		minPath.Pop().(Path),
		"we want a specific value",
	)
	require.Equalf(t, minPath.Len(), 2, "there are 2 nodes left")
}
