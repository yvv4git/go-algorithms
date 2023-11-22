package _45_reverse_vowels_string

import "strings"

func reverseVowelsV2(s string) string {
	/*
		Method: Window loop.
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	low, high := 0, len(s)-1
	vowels := "aeiouAEIOU"
	runeSlice := []rune(s) // Convert the string to a slice of runes

	for low < high {
		// Move the `low` pointer to the right until it points to a vowel
		for low < high && !strings.Contains(vowels, string(runeSlice[low])) {
			low++
		}
		// Move the `high` pointer to the left until it points to a vowel
		for low < high && !strings.Contains(vowels, string(runeSlice[high])) {
			high--
		}

		if low < high {
			// Swap the vowels at positions pointed by `low` and `high`
			runeSlice[low], runeSlice[high] = runeSlice[high], runeSlice[low]
			low++
			high--
		}
	}

	// Convert the rune slice back to a string
	return string(runeSlice)
}
