package keyboard_row

import "strings"

func findWordsV2(words []string) []string {
	var result []string

	for _, word := range words {
		if inSame(word) {
			result = append(result, word)
		}
	}

	return result
}

func inSame(word string) bool {
	info := [26]int{}
	for _, r := range "qwertyuiop" {
		info[r-'a'] = 1
	}
	for _, r := range "asdfghjkl" {
		info[r-'a'] = 2
	}
	for _, r := range "zxcvbnm" {
		info[r-'a'] = 3
	}

	word = strings.ToLower(word)
	row := info[word[0]-'a']
	for i := 1; i < len(word); i++ {
		if info[word[i]-'a'] != row {
			return false
		}
	}

	return true
}
