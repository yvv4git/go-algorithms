package main

import "fmt"

func computePrefixFunction(pattern string) []int {
	/*
		Строит префикс-функцию π для паттерна. Используется для построения
		конечного автомата.
	*/
	m := len(pattern)
	pi := make([]int, m)

	for i := 1; i < m; i++ {
		j := pi[i-1]

		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}

		if pattern[i] == pattern[j] {
			j++
		}

		pi[i] = j
	}

	return pi
}

func BuildAutomaton(pattern string) [][256]int {
	/*
		METHOD: Построение конечного автомата по префикс-функции
		TIME COMPLEXITY: O(m × σ), где m — длина паттерна, σ — размер алфавита (256)
		SPACE COMPLEXITY: O(m × σ)

		Автомат имеет m+1 состояний (0..m). Переход δ(q, ch) определяет длину
		наибольшего префикса pattern, который является суффиксом прочитанного текста
		(pattern[0..q-1] + ch). Поиск по тексту сводится к проходу по автомату
		без откатов: O(n).
	*/
	m := len(pattern)
	pi := computePrefixFunction(pattern)

	aut := make([][256]int, m+1)

	for q := 0; q <= m; q++ {
		for ch := 0; ch < 256; ch++ {
			if q < m && byte(ch) == pattern[q] {
				aut[q][ch] = q + 1
			} else {
				if q == 0 {
					aut[q][ch] = 0
				} else {
					aut[q][ch] = aut[pi[q-1]][ch]
				}
			}
		}
	}

	return aut
}

func AutomatonSearch(text, pattern string) []int {
	/*
		METHOD: Поиск с помощью конечного автомата
		TIME COMPLEXITY: O(n + m × σ), где n — длина текста, m — длина паттерна
		SPACE COMPLEXITY: O(m × σ)

		Проход по тексту без откатов. Автомат всегда в одном из m+1 состояний.
		Как только достигнуто состояние m — паттерн найден.
	*/
	if len(pattern) == 0 {
		return []int{}
	}

	aut := BuildAutomaton(pattern)
	var matches []int
	q := 0

	for i := 0; i < len(text); i++ {
		q = aut[q][text[i]]

		if q == len(pattern) {
			matches = append(matches, i-len(pattern)+1)
		}
	}

	return matches
}

func main() {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	positions := AutomatonSearch(text, pattern)

	fmt.Printf("Текст:    %s\n", text)
	fmt.Printf("Паттерн:  %s\n", pattern)
	if len(positions) > 0 {
		fmt.Printf("Найден на позициях: %v\n", positions)
	} else {
		fmt.Println("Паттерн не найден в тексте.")
	}
}
