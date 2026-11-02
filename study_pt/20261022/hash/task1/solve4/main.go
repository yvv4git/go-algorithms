package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type samState struct {
	next   [26]int
	link   int
	length int
}

func newSAMState(length int) samState {
	state := samState{link: -1, length: length}
	for i := range state.next {
		state.next[i] = -1
	}
	return state
}

func buildSAM(s string) ([]samState, []int) {
	states := make([]samState, 0, 2*len(s)+1)
	states = append(states, newSAMState(0))
	prefixState := make([]int, len(s)+1)
	last := 0

	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		current := len(states)
		states = append(states, newSAMState(states[last].length+1))

		p := last
		for p != -1 && states[p].next[c] == -1 {
			states[p].next[c] = current
			p = states[p].link
		}

		if p == -1 {
			states[current].link = 0
		} else {
			q := states[p].next[c]
			if states[p].length+1 == states[q].length {
				states[current].link = q
			} else {
				clone := len(states)
				cloneState := newSAMState(states[p].length + 1)
				cloneState.next = states[q].next
				cloneState.link = states[q].link
				states = append(states, cloneState)

				for p != -1 && states[p].next[c] == q {
					states[p].next[c] = clone
					p = states[p].link
				}
				states[q].link = clone
				states[current].link = clone
			}
		}

		last = current
		prefixState[i+1] = current
	}
	return states, prefixState
}

func buildLCA(states []samState) ([][]int32, []int) {
	m := len(states)
	levels := 1
	for (1 << levels) < m {
		levels++
	}

	up := make([][]int32, levels)
	up[0] = make([]int32, m)
	up[0][0] = 0
	for v := 1; v < m; v++ {
		up[0][v] = int32(states[v].link)
	}
	for level := 1; level < levels; level++ {
		up[level] = make([]int32, m)
		for v := 0; v < m; v++ {
			up[level][v] = up[level-1][up[level-1][v]]
		}
	}

	depth := make([]int, m)
	order := make([]int, m)
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		return states[order[i]].length < states[order[j]].length
	})
	for _, v := range order {
		if v == 0 {
			continue
		}
		depth[v] = depth[states[v].link] + 1
	}

	return up, depth
}

func findLCA(a, b int, up [][]int32, depth []int) int {
	if depth[a] < depth[b] {
		a, b = b, a
	}

	difference := depth[a] - depth[b]
	for level := 0; difference > 0; level++ {
		if difference&1 == 1 {
			a = int(up[level][a])
		}
		difference >>= 1
	}

	if a == b {
		return a
	}

	for level := len(up) - 1; level >= 0; level-- {
		if up[level][a] != up[level][b] {
			a = int(up[level][a])
			b = int(up[level][b])
		}
	}
	return int(up[0][a])
}

func solve(s string, queries [][3]int) []string {
	n := len(s)
	reversed := make([]byte, n)
	for i := 0; i < n; i++ {
		reversed[i] = s[n-1-i]
	}

	states, prefixState := buildSAM(string(reversed))
	up, depth := buildLCA(states)

	answers := make([]string, len(queries))
	for i, query := range queries {
		length, a, b := query[0], query[1], query[2]
		first := prefixState[n-a]
		second := prefixState[n-b]
		common := findLCA(first, second, up, depth)
		if states[common].length >= length {
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
