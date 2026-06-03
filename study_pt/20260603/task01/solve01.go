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

	pi := prefixFunction(s)

	out := make([]string, len(pi))
	for i, v := range pi {
		out[i] = fmt.Sprint(v)
	}
	fmt.Println(strings.Join(out, " "))
}
