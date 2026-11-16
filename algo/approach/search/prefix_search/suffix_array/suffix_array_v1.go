package main

import (
	"fmt"
	"sort"
)

type suffix struct {
	index int
	value string
}

func BuildSuffixArray(s string) []int {
	/*
		METHOD: Построение суффиксного массива (наивное, O(n² log n))
		TIME COMPLEXITY: O(n² log n) — наивная сортировка всех суффиксов
		SPACE COMPLEXITY: O(n²) — хранение всех суффиксов как строк

		Суффиксный массив — массив индексов всех суффиксов строки,
		отсортированных в лексикографическом порядке.

		НА ЗАМЕТКУ: существуют алгоритмы построения за O(n log n) и O(n)
		(удвоение префиксов, SA-IS и др.), но для демонстрации сути
		используется наивный вариант.
	*/
	n := len(s)
	suffixes := make([]suffix, n)

	for i := 0; i < n; i++ {
		suffixes[i] = suffix{index: i, value: s[i:]}
	}

	sort.Slice(suffixes, func(i, j int) bool {
		return suffixes[i].value < suffixes[j].value
	})

	sa := make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = suffixes[i].index
	}

	return sa
}

func SearchInSuffixArray(text, pattern string) []int {
	/*
		METHOD: Поиск подстроки через суффиксный массив (бинарный поиск)
		TIME COMPLEXITY: O(m × log n), где m — длина паттерна, n — длина текста
		SPACE COMPLEXITY: O(n)

		Любая подстрока text является префиксом какого-то суффикса.
		Бинарным поиском находим верхнюю и нижнюю границы суффиксов,
		начинающихся с pattern.
	*/
	if len(pattern) == 0 {
		return []int{}
	}

	sa := BuildSuffixArray(text)
	n := len(text)
	m := len(pattern)

	left := sort.Search(n, func(i int) bool {
		idx := sa[i]
		end := idx + m
		if end > n {
			end = n
		}
		return text[idx:end] >= pattern
	})

	right := sort.Search(n, func(i int) bool {
		idx := sa[i]
		end := idx + m
		if end > n {
			end = n
		}
		return text[idx:end] > pattern
	})

	var matches []int
	for i := left; i < right; i++ {
		matches = append(matches, sa[i])
	}

	return matches
}

func main() {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	sa := BuildSuffixArray(text)
	fmt.Printf("Текст:    %s\n", text)
	fmt.Printf("SA:       %v\n", sa)
	fmt.Println("Суффиксы (отсортированы):")
	for _, idx := range sa {
		fmt.Printf("  %2d: %s\n", idx, text[idx:])
	}

	positions := SearchInSuffixArray(text, pattern)
	fmt.Printf("\nПаттерн:  %s\n", pattern)
	if len(positions) > 0 {
		fmt.Printf("Найден на позициях: %v\n", positions)
	} else {
		fmt.Println("Паттерн не найден в тексте.")
	}
}
