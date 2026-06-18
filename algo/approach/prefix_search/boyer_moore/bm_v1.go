package main

import "fmt"

func buildBadCharTable(pattern string) [256]int {
	/*
		METHOD: Bad Character Heuristic
		TIME COMPLEXITY: O(m + σ), где m — длина pattern, σ — размер алфавита (256)
		SPACE COMPLEXITY: O(σ)

		Для каждого символа алфавита запоминается его последняя позиция в pattern.
		При несовпадении паттерн сдвигается так, чтобы несовпавший символ текста
		совместился с таким же символом в pattern (если он есть).
	*/
	var table [256]int
	for i := range table {
		table[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		table[pattern[i]] = i
	}
	return table
}

func BoyerMooreSearch(text, pattern string) []int {
	/*
		METHOD: Алгоритм Бойера-Мура (Boyer-Moore)
		TIME COMPLEXITY:
		  - Лучший случай: O(n / m) — суб-линейное время
		  - Средний случай: O(n)
		  - Худший случай: O(n × m)
		SPACE COMPLEXITY: O(σ) — таблица bad character (размер алфавита)

		Сравнение символов выполняется справа налево. При несовпадении используется
		правило "плохого символа" (bad character rule): паттерн сдвигается на
		максимально возможное расстояние, определяемое таблицей последних вхождений.
		Благодаря этому алгоритм на практике часто пропускает большие участки текста.
	*/
	if len(pattern) == 0 {
		return []int{}
	}

	n := len(text)
	m := len(pattern)

	if m > n {
		return []int{}
	}

	badChar := buildBadCharTable(pattern)
	var matches []int
	s := 0

	for s <= n-m {
		j := m - 1

		for j >= 0 && pattern[j] == text[s+j] {
			j--
		}

		if j < 0 {
			matches = append(matches, s)
			s++
		} else {
			bcShift := j - badChar[text[s+j]]
			if bcShift < 1 {
				bcShift = 1
			}
			s += bcShift
		}
	}

	return matches
}

func main() {
	/*
		Пример использования алгоритма Бойера-Мура для поиска подстроки в тексте.
		Сравнение идёт справа налево, что позволяет пропускать большие участки текста.
	*/
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	positions := BoyerMooreSearch(text, pattern)

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
