package main

import (
	"bufio"
	"fmt"
	"os"
)

func equalSubstrings(s string, a, b, length int) bool {
	for i := range length {
		if s[a+i] != s[b+i] {
			return false
		}
	}
	return true
}

func solve(s string, queries [][3]int) []string {
	answers := make([]string, len(queries))
	for i, query := range queries {
		if equalSubstrings(s, query[1], query[2], query[0]) {
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
