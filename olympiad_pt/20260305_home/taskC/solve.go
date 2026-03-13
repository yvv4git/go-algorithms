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

	prog := make([][]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(in, &k)
		prog[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(in, &prog[i][j])
		}
	}

	// dp[v] = можем ли иметь v после текущего месяца
	reachable := make([]bool, 10001)
	reachable[0] = true // начальные накопления

	// идём с начала
	for i := 0; i < n; i++ {
		next := make([]bool, 10001)
		for v := 0; v <= 10000; v++ {
			if reachable[v] {
				// Пробуем каждую программу, где a >= v
				for _, a := range prog[i] {
					if a >= v {
						next[a] = true
					}
				}
			}
		}
		reachable = next
	}

	// ищем максимальное достижимое значение
	ans := -1
	for v := 10000; v >= 0; v-- {
		if reachable[v] {
			ans = v
			break
		}
	}

	fmt.Println(ans)
}
