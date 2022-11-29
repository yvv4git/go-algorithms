package dijkstra_test

import (
	"git.cryptology.com/quotes-provider/quotes-provider/internal/utils/dijkstra"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinPath(t *testing.T) {
	minPath := &dijkstra.MinPath{}

	minPathLen := minPath.Len()
	require.Equalf(t, minPathLen, 0, "at first the path is empty")

	// Check push method.
	minPath.Push(dijkstra.NewPath(3.0, []string{"BTC"}))
	require.Equalf(t, minPath.Len(), 1, "added first node")

	minPath.Push(dijkstra.NewPath(5.2, []string{"USD"}))
	require.Equalf(t, minPath.Len(), 2, "added second node")

	minPath.Push(dijkstra.NewPath(4.89, []string{"EUR"}))
	require.Equalf(t, minPath.Len(), 3, "added third node")

	// Check pop method.
	require.Equalf(
		t,
		dijkstra.NewPath(4.89, []string{"EUR"}),
		minPath.Pop().(dijkstra.Path),
		"we want a specific value",
	)
	require.Equalf(t, minPath.Len(), 2, "there are 2 nodes left")
}
