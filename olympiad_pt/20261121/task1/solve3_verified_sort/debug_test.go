package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDebugOne(t *testing.T) {
	n, k, q := 100, 20, 11000
	mode := "random"
	seed := int64(1)

	pi := make([]int, n)
	for i := range pi {
		pi[i] = i + 1
	}
	rng := rand.New(rand.NewSource(seed))
	rng.Shuffle(n, func(i, j int) { pi[i], pi[j] = pi[j], pi[i] })
	j := &advJury{pi: pi, k: k, liesLeft: k, maxQ: q, mode: mode, rng: rng}
	got := recoverPermutationV3(n, k, q, j.ask)

	// Find first mismatch
	mism := -1
	for i := range pi {
		if got[i] != pi[i] {
			mism = i
			break
		}
	}
	// Count inversions in got relative to pi
	// got[i] = value at position i+1
	// Check adjacent pairs in got's order of positions
	// Build position order: positions sorted by true value
	trueOrder := make([]int, n)
	for i := range trueOrder {
		trueOrder[i] = i + 1
	}
	// sort trueOrder by pi
	for i := 0; i < n; i++ {
		for j2 := i + 1; j2 < n; j2++ {
			if pi[trueOrder[i]-1] > pi[trueOrder[j2]-1] {
				trueOrder[i], trueOrder[j2] = trueOrder[j2], trueOrder[i]
			}
		}
	}
	// got order: p[pos-1]=value => position of value v is...
	// actually got[i] means pi(i+1)=got[i], so order of positions by value:
	gotOrder := make([]int, n)
	for pos := 1; pos <= n; pos++ {
		gotOrder[got[pos-1]-1] = pos
	}
	// gotOrder[v-1] = position of value v
	// Compare sorted positions
	adjInv := 0
	for i := 0; i+1 < n; i++ {
		// positions of values i+1 and i+2 in got
		p1, p2 := gotOrder[i], gotOrder[i+1]
		if p1 > p2 {
			adjInv++
		}
	}
	fmt.Printf("mismatch at %d, got=%v want=%v\n", mism+1, got[:min(20, n)], pi[:min(20, n)])
	fmt.Printf("queries=%d liesUsed=%d liesLeft_jury=%d adjValInv=%d\n", j.queries, k-j.liesLeft, j.liesLeft, adjInv)
	// Show where got differs
	diff := 0
	for i := range pi {
		if got[i] != pi[i] {
			diff++
		}
	}
	fmt.Printf("numDiff=%d\n", diff)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
