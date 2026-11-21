package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDebug100_20(t *testing.T) {
	n, k, q := 100, 20, 11000
	pi := make([]int, n)
	for i := range pi {
		pi[i] = i + 1
	}
	rng := rand.New(rand.NewSource(1))
	rng.Shuffle(n, func(i, j int) { pi[i], pi[j] = pi[j], pi[i] })
	j := &advJury{pi: pi, k: k, liesLeft: k, maxQ: q, mode: "late", rng: rng}
	got := recoverPermutationV3(n, k, q, j.ask)
	diff := 0
	for i := range pi {
		if got[i] != pi[i] {
			diff++
		}
	}
	fmt.Printf("diff=%d queries=%d liesUsed=%d liesLeft=%d\n", diff, j.queries, k-j.liesLeft, j.liesLeft)
}
