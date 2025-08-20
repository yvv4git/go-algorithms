//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		grid[i] = []byte(s)
	}

	// Создаем массивы для длин непрерывных последовательностей в каждом направлении
	left := make([][]int, n)
	right := make([][]int, n)
	up := make([][]int, n)
	down := make([][]int, n)

	for i := 0; i < n; i++ {
		left[i] = make([]int, m)
		right[i] = make([]int, m)
		up[i] = make([]int, m)
		down[i] = make([]int, m)
	}

	// Вычисляем длины последовательностей влево
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				left[i][j] = 0
			} else {
				if j == 0 {
					left[i][j] = 1
				} else {
					left[i][j] = left[i][j-1] + 1
				}
			}
		}
	}

	// Вычисляем длины последовательностей вправо
	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			if grid[i][j] == '0' {
				right[i][j] = 0
			} else {
				if j == m-1 {
					right[i][j] = 1
				} else {
					right[i][j] = right[i][j+1] + 1
				}
			}
		}
	}

	// Вычисляем длины последовательностей вверх
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if grid[i][j] == '0' {
				up[i][j] = 0
			} else {
				if i == 0 {
					up[i][j] = 1
				} else {
					up[i][j] = up[i-1][j] + 1
				}
			}
		}
	}

	// Вычисляем длины последовательностей вниз
	for j := 0; j < m; j++ {
		for i := n - 1; i >= 0; i-- {
			if grid[i][j] == '0' {
				down[i][j] = 0
			} else {
				if i == n-1 {
					down[i][j] = 1
				} else {
					down[i][j] = down[i+1][j] + 1
				}
			}
		}
	}

	maxSize := -1

	// Для каждой ячейки находим максимальный возможный размер плюса
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				// Находим минимальную длину во всех четырех направлениях
				minArm := min(left[i][j], right[i][j], up[i][j], down[i][j])

				// Размер плюса вычисляется по формуле 4k + 1
				// где k = minArm - 1 (так как включаем центральную ячейку)
				if minArm > 0 {
					k := minArm - 1
					size := 4*k + 1
					if size > maxSize {
						maxSize = size
					}
				}
			}
		}
	}

	fmt.Fprintln(out, maxSize)
}

func min(a, b, c, d int) int {
	minVal := a
	if b < minVal {
		minVal = b
	}
	if c < minVal {
		minVal = c
	}
	if d < minVal {
		minVal = d
	}
	return minVal
}
