//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// Создаем массив для подсчета частот (разностный массив)
	diff := make([]int, n+2)
	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		diff[l]++
		if r+1 <= n {
			diff[r+1]--
		}
	}

	// Преобразуем разностный массив в массив частот
	freq := make([]int, n)
	current := 0
	for i := 1; i <= n; i++ {
		current += diff[i]
		freq[i-1] = current
	}

	// Сортируем оба массива
	sort.Ints(a)
	sort.Ints(freq)

	// Вычисляем скалярное произведение
	var result int64
	for i := 0; i < n; i++ {
		result += int64(a[i]) * int64(freq[i])
	}

	fmt.Fprintln(out, result)
}
