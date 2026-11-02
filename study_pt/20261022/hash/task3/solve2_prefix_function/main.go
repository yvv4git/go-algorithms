package main

import (
	"bufio"
	"fmt"
	"os"
)

func prefixFunction(s string) []int {
	pi := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		j := pi[i-1]
		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func minBaseLengthV2(t string) int {
	n := len(t)
	if n == 0 {
		return 0
	}

	pi := prefixFunction(t)
	return n - pi[n-1]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t string
	if _, err := fmt.Fscan(in, &t); err != nil {
		return
	}

	fmt.Fprintln(out, minBaseLengthV2(t))
}
