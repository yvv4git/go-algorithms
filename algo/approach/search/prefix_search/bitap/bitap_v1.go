package main

import "fmt"

func BitapSearch(text, pattern string) []int {
	/*
		METHOD: Алгоритм Bitap (Shift-OR)
		TIME COMPLEXITY: O(n), если m <= 64 (размер машинного слова)
		SPACE COMPLEXITY: O(σ + m), σ — размер алфавита

		Битовый алгоритм нечёткого поиска. Каждый символ алфавита кодируется битовой
		маской: 0 в позиции j означает, что символ совпадает с pattern[j].
		Текущее состояние R — битовая маска: 0 в позиции j означает, что
		префикс pattern[0..j] совпадает с суффиксом текущей подстроки текста.

		R обновляется за O(1) через битовые операции: R = (R << 1) | mask[ch].
		Если после обновления старший бит R = 0 — паттерн найден.
	*/
	if len(pattern) == 0 {
		return []int{}
	}

	m := len(pattern)
	if m > 64 {
		panic("Bitap: паттерн длиннее 64 символов не поддерживается")
	}

	mask := make(map[byte]uint64)
	for i := 0; i < m; i++ {
		mask[pattern[i]] &^= 1 << uint(i)
	}
	for ch := ' '; ch <= '~'; ch++ {
		if _, ok := mask[byte(ch)]; !ok {
			mask[byte(ch)] = ^uint64(0)
		}
	}

	var matches []int
	r := ^uint64(0)
	highBit := uint64(1) << uint(m-1)

	for i := 0; i < len(text); i++ {
		chMask, ok := mask[text[i]]
		if !ok {
			chMask = ^uint64(0)
		}
		r = (r << 1) | chMask

		if r&highBit == 0 {
			matches = append(matches, i-m+1)
		}
	}

	return matches
}

func main() {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	positions := BitapSearch(text, pattern)

	fmt.Printf("Текст:    %s\n", text)
	fmt.Printf("Паттерн:  %s\n", pattern)
	if len(positions) > 0 {
		fmt.Printf("Найден на позициях: %v\n", positions)
	} else {
		fmt.Println("Паттерн не найден в тексте.")
	}
}
