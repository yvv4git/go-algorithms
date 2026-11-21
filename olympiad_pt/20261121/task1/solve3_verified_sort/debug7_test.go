package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// Copy of algorithm with logging - instrument via wrapper
func TestDebugPasses200(t *testing.T) {
	n, k, q := 200, 10, 11000
	pi := make([]int, n)
	for i := range pi {
		pi[i] = i + 1
	}
	rng := rand.New(rand.NewSource(20))
	rng.Shuffle(n, func(i, j int) { pi[i], pi[j] = pi[j], pi[i] })
	j := &advJury{pi: pi, k: k, liesLeft: k, maxQ: q, mode: "random", rng: rng}

	// Just run and see
	got := recoverPermutationV3(n, k, q, j.ask)
	diff := 0
	for i := range pi {
		if got[i] != pi[i] {
			diff++
		}
	}
	// count adjacent inversions in got's induced order
	// got[i] = value at position i+1
	// order of positions by value:
	posOfVal := make([]int, n+1)
	for pos := 1; pos <= n; pos++ {
		posOfVal[got[pos-1]] = pos
	}
	adjInv := 0
	for v := 1; v < n; v++ {
		if posOfVal[v] > posOfVal[v+1] {
			adjInv++
		}
	}
	fmt.Printf("diff=%d adjValInv=%d queries=%d liesUsed=%d\n", diff, adjInv, j.queries, k-j.liesLeft)
}
