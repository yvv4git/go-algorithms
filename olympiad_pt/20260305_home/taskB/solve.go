package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return
	}

	if n <= 0 || m <= 0 {
		fmt.Println(-2)
		return
	}

	// Оптимальное решение: ceil(N/2) * ceil(M/2)
	// Пункты выдачи размещаются на пересечениях сетки с шагом 2
	// Покрывают 2x2 квартала каждый

	result := ((n + 1) / 2) * ((m + 1) / 2)
	fmt.Println(result)
}
