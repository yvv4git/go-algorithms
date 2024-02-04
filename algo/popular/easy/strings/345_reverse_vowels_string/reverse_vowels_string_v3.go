package _45_reverse_vowels_string

import "unicode"

func reverseVowelsV3(s string) string {
	/*
		METHOD: Two pointer / Window.
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	vowelsMap := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	result := []rune(s)

	for i, j := 0, len(result)-1; i <= j; {
		if !vowelsMap[unicode.ToLower(result[i])] {
			i++ // Move right
		} else if !vowelsMap[unicode.ToLower(result[j])] {
			j-- // Move left
		} else {
			// It's vowel
			result[i], result[j] = result[j], result[i]
			i++
			j--
		}
	}

	return string(result)
}
