//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// Читаем n из ввода.
	var n int
	fmt.Fscan(in, &n)

	result := getMoneyAmountOptimal(n)
	fmt.Println(result)
}

func getMoneyAmountOptimal(n int) int {
	/*
		INTUITION:
		Задача сводится к нахождению оптимальной стратегии угадывания.
		Для каждого диапазона [low, high] нужно найти число k,
		которое минимизирует максимальную возможную стоимость.

		Математическое обоснование:
		- Если мы угадываем число k, то платим k
		- Далее игра продолжается в диапазоне [low, k-1] или [k+1, high]
		- В худшем случае мы платим больше из двух вариантов
		- Поэтому нужно минимизировать: k + max(dp(low, k-1), dp(k+1, high))

		APPROACH: Оптимальная стратегия (DP с оптимизацией выбора k)
		Используем рекурсию с мемоизацией, где для каждого диапазона
		находим оптимальное число для первой попытки.

		БАЗА:
		- dp(i, i) = 0 — одно число можно угадать бесплатно

		Переход:
		- dp(i, j) = min(k + max(dp(i, k-1), dp(k+1, j))) для k в [i, j]

		Оптимизация:
		- Можно использовать бинарный поиск для нахождения k,
		  так как функция выпуклая (сначала убывает, потом возрастает)
		- Но для n <= 200 достаточно простого перебора

		TIME COMPLEXITY: O(n^3) в худшем случае, O(n^2) с бинарным поиском
		SPACE COMPLEXITY: O(n^2) для мемо-таблицы
	*/

	// Создаём мемо-таблицу
	memo := make([][]int, n+2)
	for i := range memo {
		memo[i] = make([]int, n+2)
	}

	// Рекурсивная функция с мемоизацией
	var dp func(low, high int) int
	dp = func(low, high int) int {
		// Базовый случай: одно число или пустой диапазон
		if low >= high {
			return 0
		}

		// Если уже вычисляли — возвращаем из мемо
		if memo[low][high] != 0 {
			return memo[low][high]
		}

		// Находим оптимальную первую попытку
		minCost := int(^uint(0) >> 1)

		// Перебираем все возможные первые попытки
		for k := low; k <= high; k++ {
			// Стоимость = k + max(стоимость_слева, стоимость_справа)
			cost := k + max(dp(low, k-1), dp(k+1, high))
			if cost < minCost {
				minCost = cost
			}
		}

		memo[low][high] = minCost
		return minCost
	}

	return dp(1, n)
}

// Вспомогательная функция: максимум двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
