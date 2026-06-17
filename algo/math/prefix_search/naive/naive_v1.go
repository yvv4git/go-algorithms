package main

import "fmt"

func NaiveSearch(text, pattern string) []int {
	/*
		METHOD: Наивный поиск подстроки (brute force)
		TIME COMPLEXITY: O(n × m), где n — длина text, m — длина pattern
		SPACE COMPLEXITY: O(1) без учёта массива результатов

		Последовательно проверяет каждую позицию в тексте на совпадение с паттерном.
		При несовпадении сдвигается на 1 символ вправо.
	*/
	if len(pattern) == 0 {
		return []int{}
	}

	var matches []int
	n := len(text)
	m := len(pattern)

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && text[i+j] == pattern[j] {
			j++
		}
		if j == m {
			matches = append(matches, i)
		}
	}
	return matches
}

func main() {
	/*
		Пример использования наивного алгоритма поиска подстроки.
	*/
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	positions := NaiveSearch(text, pattern)

	fmt.Printf("Текст:    %s\n", text)
	fmt.Printf("Паттерн:  %s\n", pattern)
	if len(positions) > 0 {
		fmt.Printf("Найден на позициях: %v\n", positions)
		for _, pos := range positions {
			fmt.Printf("  %s\n", text[:pos]+"["+pattern+"]"+text[pos+len(pattern):])
		}
	} else {
		fmt.Println("Паттерн не найден в тексте.")
	}
}
