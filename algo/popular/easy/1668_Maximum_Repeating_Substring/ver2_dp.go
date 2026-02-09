//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// Читаем две строки из ввода.
	var sequence, word string
	fmt.Fscan(in, &sequence)
	fmt.Fscan(in, &word)

	result := maxRepeatingDP(sequence, word)
	fmt.Println(result)
}

func maxRepeatingDP(sequence string, word string) int {
	/*
		INTUITION:
		Нужно найти максимальное k, при котором word, повторенное k раз,
		является подстрокой sequence.

		Идея DP: dp[i] = максимальное количество повторений word,
		которое заканчивается в позиции i sequence (exclusive).

		Если word полностью совпадает с subsequence[i-len(word):i],
		то мы можем расширить предыдущую последовательность.

		APPROACH: Dynamic Programming (одномерный массив)
		dp[i] — максимальное k, при котором word^k заканчивается на позиции i.

		БАЗА: dp[0] = 0 (пустая строка не содержит word)
		Все элементы dp инициализированы нулями.

		Переход:
		- Если sequence[i-len(word):i] == word и i-len(word) >= 0:
		  то dp[i] = dp[i-len(word)] + 1
		- Иначе dp[i] = 0

		ПОЧЕМУ ТАК:
		- Каждая позиция i хранит информацию о законченных повторениях word
		- Если текущий сегмент равен word, мы расширяем предыдущую цепочку
		- Ответ — максимальное значение в массиве dp

		TIME COMPLEXITY: O(n * m) где n = len(sequence), m = len(word)
		SPACE COMPLEXITY: O(n) для хранения dp-таблицы
	*/

	n := len(sequence)
	m := len(word)

	// dp[i] — количество последовательных повторений word, которые заканчиваются на позиции i.
	// Например: sequence="ababab", word="ab", i=6 -> dp[6]=3 ("ab" + "ab" + "ab" заканчивается здесь).
	dp := make([]int, n+1)

	// maxK — хранит максимальное количество повторений, найденное на данный момент.
	maxK := 0

	for i := m; i <= n; i++ {
		// Проверяем, заканчивается ли word на позиции i.
		if sequence[i-m:i] == word {
			// Расширяем предыдущую цепочку.
			dp[i] = dp[i-m] + 1
			// Обновляем максимальное значение.
			if dp[i] > maxK {
				maxK = dp[i]
			}
		}
	}

	return maxK
}
