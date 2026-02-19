//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// Chitaem dva chisla iz vvoda.
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

		APPROACH: DP with bitmask (iterative version)
		Итеративный перебор всех возможных состояний игры, начиная с конца.
		Состояние представляется битовой маской использованных чисел.
		Для каждого состояния определяем, может ли текущий игрок выиграть.

		WHY THIS WAY:
		- Итеративный подход позволяет избежать накладных расходов рекурсии
		- Заполняем DP-массив от конца к началу
		- dp[mask] = true, если игрок, делающий ход в этом состоянии, побеждает

		ОПИСАНИЕ РЕШЕНИЯ:
		Динамическое программирование с битовой маской. Перебираем все возможные
		состояния (2^n вариантов) и для каждого вычисляем, может ли игрок,
		который сейчас должен ходить, обеспечить себе победу. Предварительно
		вычисляем суммы для каждой маски для оптимизации.

		TIME COMPLEXITY: O(2^n * n) where n = maxChoosableInteger
		SPACE COMPLEXITY: O(2^n) for dp array
	*/

	// Base cases
	sum := maxChoosableInteger * (maxChoosableInteger + 1) / 2
	if sum < desiredTotal {
		return false
	}

	if desiredTotal <= 0 {
		return true
	}

	n := maxChoosableInteger
	totalStates := 1 << n // number of possible states

	// dp[mask] = true if with current set of used numbers
	// current player can win
	dp := make([]bool, totalStates)

	// Calculate sum of numbers for each mask
	// sumMask[mask] = sum of all numbers set in mask
	sumMask := make([]int, totalStates)
	for mask := 1; mask < totalStates; mask++ {
		// Find lowest set bit and add corresponding number
		lowestBit := mask & -mask
		bitPos := 0
		temp := lowestBit
		for temp > 1 {
			temp >>= 1
			bitPos++
		}
		// Number = bitPos + 1
		sumMask[mask] = sumMask[mask^lowestBit] + (bitPos + 1)
	}

	// Iterate through all states in reverse order (from larger mask to smaller)
	// This allows using already computed values
	for mask := totalStates - 1; mask >= 0; mask-- {
		currentSum := sumMask[mask]
		remaining := desiredTotal - currentSum

		// If remaining <= 0, game is already over
		// This shouldn't happen as we can't be in a state
		// where game is over but we're still making a move
		if remaining <= 0 {
			continue
		}

		// Check each possible move
		for i := 0; i < n; i++ {
			// If number (i+1) hasn't been used yet
			bit := 1 << i
			if mask&bit == 0 {
				// If choosing (i+1) is enough to win
				if (i + 1) >= remaining {
					dp[mask] = true
					break
				}
				// Otherwise check if opponent loses
				// after our move
				newMask := mask | bit
				if !dp[newMask] {
					dp[mask] = true
					break
				}
			}
		}
		// If no move leads to victory, dp[mask] remains false
	}

	return dp[0]
}
