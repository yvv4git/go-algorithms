package _59_repeated_substring_pattern

import "strings"

func repeatedSubstringPatternV2(s string) bool {
	/*
		METHOD:
	*/
	// Создаем подстроку, которая начинается со второго символа исходной строки
	sub := s[1:]

	// Пока длина подстроки больше 0
	for len(sub) > 0 {
		// Если длина подстроки делится нацело на длину исходной строки минус длина подстроки и подстрока входит в исходную строку
		if len(s)%len(sub) == 0 && strings.Contains(s, sub) {
			// Создаем новую подстроку, которая начинается со следующего символа исходной подстроки
			sub = sub[1:]
		} else {
			// Иначе возвращаем false
			return false
		}
	}

	// Если мы дошли до этой точки, то подстрока не найдена, возвращаем false
	return false
}
