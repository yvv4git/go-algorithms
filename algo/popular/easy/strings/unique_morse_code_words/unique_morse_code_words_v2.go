package unique_morse_code_words

import "strings"

func uniqueMorseRepresentationsV2(words []string) int {
	wordSet := make(map[string]struct{})

	alphabet := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	for i := 0; i < len(words); i++ {
		var currentWord strings.Builder

		for j := 0; j < len(words[i]); j++ {
			currentWord.WriteString(alphabet[words[i][j]-'a'])
		}

		if _, ok := wordSet[currentWord.String()]; !ok {
			wordSet[currentWord.String()] = struct{}{}
		}
	}

	return len(wordSet)
}
