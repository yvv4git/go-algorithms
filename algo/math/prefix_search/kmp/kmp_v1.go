package main

import "fmt"

func computeLPS(pattern string) []int {
	/*
		METHOD: LPS (Longest Prefix Suffix)
		TIME COMPLEXITY: O(m), где m — длина pattern
		SPACE COMPLEXITY: O(m) для хранения массива lps

		LPS[i] — длина наибольшего собственного префикса, который одновременно является суффиксом
		для подстроки pattern[0..i]. Используется в KMP, чтобы избежать лишних сравнений.
	*/
	m := len(pattern)
	lps := make([]int, m)
	j := 0

	for i := 1; i < m; {
		if pattern[i] == pattern[j] {
			j++
			lps[i] = j
			i++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func KMPSearch(text, pattern string) []int {
	/*
		METHOD: Алгоритм Кнута-Морриса-Пратта (KMP)
		TIME COMPLEXITY: O(n + m), где n — длина text, m — длина pattern
		SPACE COMPLEXITY: O(m)

		Алгоритм использует префикс-функцию (LPS), чтобы при несовпадении
		сдвигать pattern не на 1 символ, а на максимально возможную величину.
	*/
	if len(pattern) == 0 {
		return []int{}
	}

	lps := computeLPS(pattern)
	var matches []int
	j := 0

	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != pattern[j] {
			j = lps[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == len(pattern) {
			matches = append(matches, i-j+1)
			j = lps[j-1]
		}
	}
	return matches
}

func main() {
	/*
		Пример использования алгоритма Кнута-Морриса-Пратта для поиска подстроки в строке.
	*/
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	positions := KMPSearch(text, pattern)

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
