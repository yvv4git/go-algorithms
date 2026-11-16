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

	prices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &prices[i])
	}

	if n == 0 {
		fmt.Println(0)
		fmt.Println("0 0")
		return
	}

	// Максимальное возможное количество купонов
	maxCoupons := n + 1

	// dp[i][c] = минимальная стоимость за первые i дней с c купонами в конце
	dp := make([][]int, n+1)
	parent := make([][]int, n+1) // -1 - купон не использовали, -2 - использовали купон

	for i := range dp {
		dp[i] = make([]int, maxCoupons)
		parent[i] = make([]int, maxCoupons)
		for j := range dp[i] {
			dp[i][j] = 1_000_000_000
			parent[i][j] = -3
		}
	}

	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		price := prices[i-1]

		for c := 0; c < maxCoupons; c++ {
			// Вариант 1: Покупаем обед за деньги
			newCoupons := c
			if price > 100 {
				newCoupons = c + 1
			}

			if newCoupons < maxCoupons {
				if dp[i-1][c]+price < dp[i][newCoupons] {
					dp[i][newCoupons] = dp[i-1][c] + price
					parent[i][newCoupons] = c // пришли из состояния c
				}
			}

			// Вариант 2: Используем купон (если он есть)
			if c > 0 {
				newCoupons := c - 1
				if price > 100 {
					newCoupons = c
				}

				if newCoupons < maxCoupons && dp[i-1][c] < dp[i][newCoupons] {
					dp[i][newCoupons] = dp[i-1][c]
					parent[i][newCoupons] = -2 // метка, что использовали купон
				}
			}
		}
	}

	// Находим минимальную стоимость
	minCost := 1_000_000_000
	bestC := 0

	for c := 0; c < maxCoupons; c++ {
		if dp[n][c] < minCost {
			minCost = dp[n][c]
			bestC = c
		} else if dp[n][c] == minCost && c > bestC {
			bestC = c
		}
	}

	// Восстанавливаем ответ
	couponsLeft := bestC
	usedDays := []int{}

	// Восстанавливаем использование купонов
	c := bestC
	for i := n; i > 0; i-- {
		prev := parent[i][c]

		if prev == -2 {
			// Использовали купон
			usedDays = append(usedDays, i)
			// Нужно найти, из какого состояния пришли
			for prevC := 0; prevC < maxCoupons; prevC++ {
				// Проверяем, могли ли мы прийти из состояния prevC
				// с использованием купона
				if dp[i-1][prevC] == dp[i][c] {
					price := prices[i-1]
					expectedC := c
					if price > 100 {
						expectedC = c + 1
					}
					if prevC == expectedC {
						c = prevC
						break
					}
				}
			}
		} else if prev >= 0 {
			// Не использовали купон
			c = prev
		}
	}

	// Разворачиваем дни в правильном порядке
	for i, j := 0, len(usedDays)-1; i < j; i, j = i+1, j-1 {
		usedDays[i], usedDays[j] = usedDays[j], usedDays[i]
	}

	K2 := len(usedDays)
	K1 := couponsLeft

	fmt.Println(minCost)
	fmt.Printf("%d %d\n", K1, K2)
	for _, day := range usedDays {
		fmt.Println(day)
	}
}
