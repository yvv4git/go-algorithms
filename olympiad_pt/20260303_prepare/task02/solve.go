package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var k int64
	fmt.Fscan(in, &n, &k)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// Сортируем оценки
	slices.Sort(a)

	// Бинарный поиск по ответу
	left := a[0]      // минимальная оценка
	right := a[0] + k // теоретический максимум

	for left < right {
		mid := (left + right + 1) / 2 // берем верхнюю середину
		if canAchieve(a, mid, k) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	fmt.Fprintln(out, left)
}

// canAchieve проверяет, можно ли поднять все оценки до target за k часов
func canAchieve(a []int64, target int64, k int64) bool {
	var need int64 = 0
	for _, v := range a {
		if v < target {
			need += target - v
			if need > k {
				return false
			}
		}
	}

	return need <= k
}
