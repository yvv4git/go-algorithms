package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDebug1000(t *testing.T) {
	n, k, q := 1000, 1, 11000
	for _, mode := range []string{"late", "random", "corrupt"} {
		pi := make([]int, n)
		for i := range pi {
			pi[i] = i + 1
		}
		rng := rand.New(rand.NewSource(10))
		rng.Shuffle(n, func(i, j int) { pi[i], pi[j] = pi[j], pi[i] })
		j := &advJury{pi: pi, k: k, liesLeft: k, maxQ: q, mode: mode, rng: rng}
		got := recoverPermutationV3(n, k, q, j.ask)
		diff := 0
		for i := range pi {
			if got[i] != pi[i] {
				diff++
			}
		}
		fmt.Printf("mode=%s diff=%d queries=%d liesUsed=%d\n", mode, diff, j.queries, k-j.liesLeft)
	}
}
