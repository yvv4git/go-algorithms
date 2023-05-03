package unique_morse_code_words

import "strings"

func uniqueMorseRepresentationsV3(words []string) int {
	/*
		Time complexity : O(n)
		Space complexity : O(n)
	*/
	morseMap := map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
	}

	wordsMap := map[string]bool{}
	for _, word := range words {
		wordsMap[word] = true
	}

	uniqMorseWords := map[string]bool{}
	for word, _ := range wordsMap {
		var code strings.Builder
		for _, letter := range word {
			code.WriteString(morseMap[string(letter)])
		}
		uniqMorseWords[code.String()] = true
	}

	return len(uniqMorseWords)
}
