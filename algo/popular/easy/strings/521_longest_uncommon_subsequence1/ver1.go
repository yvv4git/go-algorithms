package main

import "fmt"

func findLUSlength(a string, b string) int {
	/*
		METHOD: Compare
		Для решения этой задачи мы можем использовать следующий подход:
		1. Если две строки равны, то необычной подпоследовательностью будет сама строка.
		В этом случае мы возвращаем -1, так как подпоследовательность не должна повторяться.
		2. Если строки не равны, то самая длинная необычная подпоследовательность будет строкой, которая длиннее.

		TIME COMPLEXITY: O(1), так как мы просто сравниваем две строки.

		SPACE COMPLEXITY: O(1), так как мы не используем дополнительную память, которая бы растягивалась с увеличением входных данных.
	*/
	// Если строки равны, то необычной подпоследовательностью нет
	if a == b {
		return -1
	}
	// Если строки не равны, то самая длинная необычная подпоследовательность будет строкой, которая длиннее
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}

func main() {
	fmt.Println(findLUSlength("abc", "abc")) // Выведет: -1
	fmt.Println(findLUSlength("abc", "def")) // Выведет: 3
}
