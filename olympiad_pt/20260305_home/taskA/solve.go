package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	if n < 0 {
		fmt.Println(-2)
		return
	}

	bestIdx := 1
	bestK := int64(0) // точка окупаемости лучшего станка

	for i := 1; i <= n; i++ {
		var a, b, c int64
		fmt.Fscan(in, &a, &b, &c)

		// Точка окупаемости: минимальное k: a + k*b <= k*c
		// a <= k * (c - b)
		// k >= ceil(a / (c - b))
		delta := c - b               // > 0 по условию b < c
		k := (a + delta - 1) / delta // ceil(a / delta)

		if i == 1 || k < bestK || (k == bestK && i < bestIdx) {
			bestK = k
			bestIdx = i
		}
	}

	fmt.Println(bestIdx)
}
