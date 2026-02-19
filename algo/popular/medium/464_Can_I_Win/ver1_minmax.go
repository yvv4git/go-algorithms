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

		APPROACH: Minimax with memoization
		Рекурсивный поиск с запоминанием результатов для каждого состояния.
		Состояние кодируется битовой маской - каждый бит показывает,
		было ли соответствующее число (1..maxChoosableInteger) использовано.

		WHY THIS WAY:
		- Minimax моделирует оптимальную игру обоих игроков
		- Мемоизация позволяет избежать повторных вычислений идентичных состояний
		- Битовые маски эффективно представляют наборы доступных чисел

		ОПИСАНИЕ РЕШЕНИЯ:
		Два игрока поочередно выбирают числа из диапазона 1..maxChoosableInteger.
		Каждое число может быть использовано только один раз. Игрок, первым
		достигший суммы >= desiredTotal, побеждает. Используем рекурсивный
		алгоритм minimax с мемоизацией для определения выигрышной стратегии.

		TIME COMPLEXITY: O(2^n * n) where n = maxChoosableInteger
		SPACE COMPLEXITY: O(2^n) for memoization storage
	*/

	// Base cases
	// If sum of all numbers is less than desiredTotal, no one can win
	sum := maxChoosableInteger * (maxChoosableInteger + 1) / 2
	if sum < desiredTotal {
		return false
	}

	// If desiredTotal <= 0, first player already won
	if desiredTotal <= 0 {
		return true
	}

	// Memoization: map[bitmask] -> result
	memo := make(map[int]bool)

	// Start recursion with empty set of used numbers
	return minimax(0, desiredTotal, maxChoosableInteger, memo)
}

// minimax returns true if current player can win
// used - bitmask of used numbers
// remaining - how much more is needed
// maxInt - maximum available number
func minimax(used, remaining, maxInt int, memo map[int]bool) bool {
	// If already computed for this state
	if result, ok := memo[used]; ok {
		return result
	}

	// Try to choose each available number
	for i := 1; i <= maxInt; i++ {
		// Check if number i has already been used
		bit := 1 << (i - 1)
		if used&bit != 0 {
			continue
		}

		// If choosing i reaches or exceeds remaining - we win
		if i >= remaining {
			memo[used] = true
			return true
		}

		// Recursively check: if opponent loses after our move,
		// then we win
		newUsed := used | bit
		if !minimax(newUsed, remaining-i, maxInt, memo) {
			memo[used] = true
			return true
		}
	}

	// If no move leads to victory - we lose
	memo[used] = false
	return false
}
