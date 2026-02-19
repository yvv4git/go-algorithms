//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// Read two numbers from input.
	var maxChoosableInteger, desiredTotal int
	fmt.Fscan(in, &maxChoosableInteger)
	fmt.Fscan(in, &desiredTotal)

	result := canIWin(maxChoosableInteger, desiredTotal)
	fmt.Println(result)
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	/*
		INTUITION:
		Need to determine if the first player can guarantee a win
		with optimal play from both participants.

		APPROACH: Backtracking without memoization
		Полный рекурсивный перебор всех возможных комбинаций ходов.
		Для каждого состояния пробуем все доступные числа и проверяем,
		приводит ли хотя бы один ход к победе.

		WHY THIS WAY:
		- Простая и понятная реализация
		- Демонстрирует базовую концепцию поиска по дереву игры
		- Менее эффективна, чем мемоизация, из-за повторных вычислений состояний
		- Экспоненциальная сложность без кэширования

		ОПИСАНИЕ РЕШЕНИЯ:
		Рекурсивный обратный поиск (backtracking) без оптимизации. Игрок
		пробует все доступные числа по очереди. Если выбранное число достаточно
		для достижения целевой суммы - победа. Иначе рекурсивно проверяем,
		сможет ли противник выиграть после нашего хода. Если противник проигрывает -
		мы выигрываем.

		TIME COMPLEXITY: O(n!) in worst case - exponential without memoization
		SPACE COMPLEXITY: O(n) for recursion stack
	*/

	// Base cases
	sum := maxChoosableInteger * (maxChoosableInteger + 1) / 2
	if sum < desiredTotal {
		return false
	}

	if desiredTotal <= 0 {
		return true
	}

	// used[i] = true if number (i+1) has been used
	used := make([]bool, maxChoosableInteger)

	return backtrack(used, desiredTotal, maxChoosableInteger)
}

// backtrack returns true if current player can win
// used - slice tracking which numbers have been used
// remaining - how much more is needed to reach target
// maxInt - maximum available number
func backtrack(used []bool, remaining, maxInt int) bool {
	// Try each available number
	for i := 0; i < maxInt; i++ {
		if used[i] {
			continue
		}

		// If choosing (i+1) is enough to win
		if (i + 1) >= remaining {
			return true
		}

		// Mark number as used
		used[i] = true

		// Recursively check if opponent loses after our move
		// If opponent can't win, we win
		if !backtrack(used, remaining-(i+1), maxInt) {
			used[i] = false // backtrack
			return true
		}

		// Backtrack - unmark the number
		used[i] = false
	}

	// If no move leads to victory, we lose
	return false
}
