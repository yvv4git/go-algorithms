package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil || n < 2 {
		fmt.Println(-2)
		return
	}

	// Читаем программы
	prog := make([][]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(in, &k)
		prog[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(in, &prog[i][j])
		}
	}

	// Greedy: всегда выбираем максимальную программу, которую можем использовать
	money := 0
	for i := 0; i < n; i++ {
		best := -1
		for _, a := range prog[i] {
			if a >= money && a > best {
				best = a
			}
		}
		if best == -1 {
			fmt.Println(-1)
			return
		}
		money = best
	}

	fmt.Println(money)
}
