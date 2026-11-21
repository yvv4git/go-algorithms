package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func TestGenPI1000(t *testing.T) {
	n := 1000
	pi := make([]int, n)
	for i := range pi {
		pi[i] = i + 1
	}
	rng := rand.New(rand.NewSource(8))
	rng.Shuffle(n, func(i, j int) { pi[i], pi[j] = pi[j], pi[i] })
	parts := make([]string, n)
	for i, v := range pi {
		parts[i] = fmt.Sprint(v)
	}
	os.Stdout.WriteString("PI=" + strings.Join(parts, ",") + "\n")
}
