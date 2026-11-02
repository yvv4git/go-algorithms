package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func buildSuffixArray(s string) []int {
	n := len(s)
	if n == 0 {
		return nil
	}

	sa := make([]int, n)
	rank := make([]int, n)
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = i
		rank[i] = int(s[i])
	}

	for k := 1; k < n; k <<= 1 {
		sort.Slice(sa, func(i, j int) bool {
			a, b := sa[i], sa[j]
			if rank[a] != rank[b] {
				return rank[a] < rank[b]
			}
			ra, rb := -1, -1
			if a+k < n {
				ra = rank[a+k]
			}
			if b+k < n {
				rb = rank[b+k]
			}
			return ra < rb
		})

		tmp[sa[0]] = 0
		for i := 1; i < n; i++ {
			a, b := sa[i-1], sa[i]
			same := rank[a] == rank[b]
			if same {
				ra, rb := -1, -1
				if a+k < n {
					ra = rank[a+k]
				}
				if b+k < n {
					rb = rank[b+k]
				}
				same = ra == rb
			}
			if same {
				tmp[b] = tmp[a]
			} else {
				tmp[b] = tmp[a] + 1
			}
		}
		copy(rank, tmp)
		if rank[sa[n-1]] == n-1 {
			break
		}
	}
	return sa
}

func buildLCP(s string, sa []int) []int {
	n := len(s)
	if n <= 1 {
		return nil
	}

	rank := make([]int, n)
	for i, position := range sa {
		rank[position] = i
	}

	lcp := make([]int, n-1)
	common := 0
	for i := 0; i < n; i++ {
		r := rank[i]
		if r == n-1 {
			common = 0
			continue
		}
		j := sa[r+1]
		for i+common < n && j+common < n && s[i+common] == s[j+common] {
			common++
		}
		lcp[r] = common
		if common > 0 {
			common--
		}
	}
	return lcp
}

func buildSparseTable(values []int) ([][]int, []int) {
	m := len(values)
	if m == 0 {
		return nil, nil
	}

	logs := make([]int, m+1)
	for i := 2; i <= m; i++ {
		logs[i] = logs[i/2] + 1
	}

	levels := logs[m] + 1
	table := make([][]int, levels)
	table[0] = make([]int, m)
	copy(table[0], values)
	for level := 1; level < levels; level++ {
		size := m - (1 << level) + 1
		table[level] = make([]int, size)
		half := 1 << (level - 1)
		for i := 0; i < size; i++ {
			table[level][i] = min(table[level-1][i], table[level-1][i+half])
		}
	}
	return table, logs
}

func rangeMinimum(table [][]int, logs []int, left, right int) int {
	level := logs[right-left+1]
	return min(table[level][left], table[level][right-(1<<level)+1])
}

func solve(s string, queries [][3]int) []string {
	n := len(s)
	sa := buildSuffixArray(s)
	rank := make([]int, n)
	for i, position := range sa {
		rank[position] = i
	}
	lcp := buildLCP(s, sa)
	table, logs := buildSparseTable(lcp)

	answers := make([]string, len(queries))
	for i, query := range queries {
		length, a, b := query[0], query[1], query[2]
		if a == b {
			answers[i] = "yes"
			continue
		}

		left, right := rank[a], rank[b]
		if left > right {
			left, right = right, left
		}
		common := rangeMinimum(table, logs, left, right-1)
		if common >= length {
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
