package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	// Читаем количество элементов
	var n int
	fmt.Fscan(in, &n)

	// Читаем массив words
	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &words[i])
	}

	// Читаем массив groups
	groups := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &groups[i])
	}

	result := getLongestSubsequenceDP(words, groups)
	for i, word := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(word)
	}
	fmt.Println()
}

func getLongestSubsequenceDP(words []string, groups []int) []string {
	/*
		INTUITION:
		Нужно найти самую длинную подпоследовательность из words, где
		соседние элементы имеют разные значения в groups (чередование 0 и 1).

		APPROACH: Dynamic Programming
		Используем DP для отслеживания максимальной длины подпоследовательности,
		заканчивающейся на каждом индексе с определённым значением группы.

		dp[i] = длина максимальной подпоследовательности, заканчивающейся на индексе i
		prev[i] = индекс предыдущего элемента в подпоследовательности (для восстановления)

		ПОЧЕМУ ТАК:
		- Для каждого элемента ищем предыдущий элемент с противоположной группой
		- Выбираем тот, который даёт максимальную длину подпоследовательности
		- Хотя для этой задачи жадный подход оптимален, DP демонстрирует
		  общий подход к задачам о подпоследовательностях

		TIME COMPLEXITY: O(n²) - для каждого элемента проверяем все предыдущие
		SPACE COMPLEXITY: O(n) - для хранения dp и prev массивов
	*/

	n := len(words)
	if n == 0 {
		return []string{}
	}

	// dp[i] - длина максимальной подпоследовательности, заканчивающейся на индексе i
	dp := make([]int, n)
	// prev[i] - индекс предыдущего элемента в подпоследовательности (-1 если нет)
	prev := make([]int, n)

	// Инициализация: каждый элемент сам по себе образует подпоследовательность длины 1
	for i := 0; i < n; i++ {
		dp[i] = 1
		prev[i] = -1
	}

	// Заполняем dp массив
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// Если группы разные и подпоследовательность через j длиннее
			if groups[i] != groups[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
	}

	// Находим индекс конца максимальной подпоследовательности
	maxLen := 0
	endIdx := 0
	for i := 0; i < n; i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
			endIdx = i
		}
	}

	// Восстанавливаем подпоследовательность
	result := make([]string, 0, maxLen)
	for endIdx != -1 {
		result = append([]string{words[endIdx]}, result...)
		endIdx = prev[endIdx]
	}

	return result
}
