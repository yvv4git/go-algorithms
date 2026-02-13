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

	result := getLongestSubsequence(words, groups)
	for i, word := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(word)
	}
	fmt.Println()
}

func getLongestSubsequence(words []string, groups []int) []string {
	/*
		INTUITION:
		Нужно найти самую длинную подпоследовательность из words, где
		соседние элементы имеют разные значения в groups (чередование 0 и 1).

		APPROACH: Greedy
		Проходим по массиву слева направо и добавляем элемент в результат,
		если его группа отличается от группы последнего добавленного элемента.

		ПОЧЕМУ ТАК:
		- Значения в groups бинарные (только 0 или 1)
		- Каждый раз, когда мы видим элемент с отличающейся группой,
		  мы можем безопасно его добавить
		- Это максимально увеличивает длину подпоследовательности
		- Жадный выбор локально оптимален и приводит к глобальному оптимуму

		TIME COMPLEXITY: O(n) - один проход по массиву
		SPACE COMPLEXITY: O(n) - для хранения результата
	*/

	n := len(words)
	if n == 0 {
		return []string{}
	}

	// Добавляем первый элемент в результат
	result := []string{words[0]}
	// Запоминаем группу последнего добавленного элемента
	lastGroup := groups[0]

	// Проходим по оставшимся элементам
	for i := 1; i < n; i++ {
		// Если группа текущего элемента отличается от последнего добавленного
		if groups[i] != lastGroup {
			result = append(result, words[i])
			lastGroup = groups[i]
		}
	}

	return result
}
