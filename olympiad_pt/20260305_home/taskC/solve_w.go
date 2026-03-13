package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil || n < 2 {
		fmt.Println(-2)
		return
	}

	// Читаем программы, сортируем
	prog := make([][]int, n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(in, &k)
		prog[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(in, &prog[i][j])
		}
		sort.Ints(prog[i])
	}

	// Обратная динамика: идём с конца
	need := 0
	for i := n - 1; i >= 0; i-- {
		ok := false
		for _, a := range prog[i] {
			if a >= need {
				need = a
				ok = true
				break
			}
		}
		if !ok {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(need)
}
