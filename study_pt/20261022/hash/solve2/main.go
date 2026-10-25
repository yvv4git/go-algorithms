package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	hashMod1 int64 = 1_000_000_007
	hashMod2 int64 = 1_000_000_009
	hashBase int64 = 911_382_323
)

type rollingHash struct {
	mod  int64
	pref []int64
	pow  []int64
}

func newRollingHash(s string, mod, base int64) rollingHash {
	n := len(s)
	pref := make([]int64, n+1)
	pow := make([]int64, n+1)
	pow[0] = 1
	for i := range n {
		pref[i+1] = (pref[i]*base + int64(s[i])) % mod
		pow[i+1] = pow[i] * base % mod
	}
	return rollingHash{mod: mod, pref: pref, pow: pow}
}

func (h rollingHash) substr(left, length int) int64 {
	value := h.pref[left+length] - h.pref[left]*h.pow[length]%h.mod
	if value < 0 {
		value += h.mod
	}
	return value
}

func solve(s string, queries [][3]int) []string {
	first := newRollingHash(s, hashMod1, hashBase)
	second := newRollingHash(s, hashMod2, hashBase)
	answers := make([]string, len(queries))
	for i, query := range queries {
		length, a, b := query[0], query[1], query[2]
		if first.substr(a, length) == first.substr(b, length) &&
			second.substr(a, length) == second.substr(b, length) {
			answers[i] = "yes"
		} else {
			answers[i] = "no"
		}
	}
	return answers
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	if _, err := fmt.Fscan(in, &s); err != nil {
		return
	}

	var q int
	if _, err := fmt.Fscan(in, &q); err != nil {
		return
	}

	queries := make([][3]int, q)
	for i := range queries {
		if _, err := fmt.Fscan(in, &queries[i][0], &queries[i][1], &queries[i][2]); err != nil {
			return
		}
	}

	for _, answer := range solve(s, queries) {
		fmt.Fprintln(out, answer)
	}
}
