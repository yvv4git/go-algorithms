package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func prefixFunction(s string) []int {
	n := len(s)
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		k := pi[i-1]
		for k > 0 && s[i] != s[k] {
			k = pi[k-1]
		}
		if s[i] == s[k] {
			k++
		}
		pi[i] = k
	}
	return pi
}

func main() {
	data, _ := io.ReadAll(os.Stdin)
	s := strings.TrimSpace(string(data))

	n := len(s)
	pi := prefixFunction(s)
	period := n - pi[n-1]

	fmt.Println(period)
}
