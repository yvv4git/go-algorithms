package keyboard_row

import (
	"strings"
)

func findWordsV1(words []string) []string {
	// Сначала надо создать массив и каждому коду символа задать группу, в которою он(символ) входит.
	/*
		Конструкция [r-'a'] позволяет узнать порядковый номер буквы в латинском алфавите.
	*/
	letters := [26]int{}
	for _, r := range "qwertyuiop" {
		letters[r-'a'] = 1
	}
	for _, r := range "asdfghjkl" {
		letters[r-'a'] = 2
	}
	for _, r := range "zxcvbnm" {
		letters[r-'a'] = 3
	}

	result := []string{}
	for _, word := range words {
		w := strings.ToLower(word)
		match := true
		for i := 1; i < len(w); i++ {
			if letters[w[i]-'a'] != letters[w[0]-'a'] {
				match = false
				break
			}
		}
		if match {
			result = append(result, word)
		}
	}

	return result
}
