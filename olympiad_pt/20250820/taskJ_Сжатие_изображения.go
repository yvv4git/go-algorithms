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

	grid := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &grid[i])
	}

	// Находим все возможные делители n и m
	divsN := findDivisors(n)
	divsM := findDivisors(m)

	bestH, bestW := n, m

	// Перебираем все возможные размеры блоков в порядке увеличения размера блока
	// (чтобы сначала найти самые большие блоки, которые дадут наименьшее сжатие)
	for i := len(divsN) - 1; i >= 0; i-- {
		h := divsN[i]
		for j := len(divsM) - 1; j >= 0; j-- {
			w := divsM[j]

			// Пропускаем, если результат будет больше текущего лучшего
			if (n/h)*(m/w) >= bestH*bestW {
				continue
			}

			// Проверяем, можно ли разбить изображение на блоки h x w
			valid := true
			for i := 0; i < n && valid; i += h {
				for j := 0; j < m && valid; j += w {
					firstChar := grid[i][j]
					for di := 0; di < h && valid; di++ {
						for dj := 0; dj < w && valid; dj++ {
							if grid[i+di][j+dj] != firstChar {
								valid = false
							}
						}
					}
				}
			}

			if valid {
				bestH, bestW = n/h, m/w
			}
		}
	}

	// Теперь строим результат только для лучшего варианта
	fmt.Fprintln(out, bestH, bestW)

	// Находим соответствующие размеры блоков
	hBlock := n / bestH
	wBlock := m / bestW

	for i := 0; i < bestH; i++ {
		row := make([]byte, bestW)
		for j := 0; j < bestW; j++ {
			row[j] = grid[i*hBlock][j*wBlock]
		}
		fmt.Fprintln(out, string(row))
	}
}

func findDivisors(x int) []int {
	divisors := []int{}
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
