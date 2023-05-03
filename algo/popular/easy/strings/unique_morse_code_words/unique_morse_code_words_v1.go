package unique_morse_code_words

import "fmt"

func uniqueMorseRepresentationsV1(words []string) int {
	/*
		Time complexity:100%
		Space complexity:28%
	*/
	morseMap := map[byte]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}
	morseSet := make(map[string]bool)

	for _, word := range words {
		//fmt.Println("--------->", word)
		morseWord := ""

		for _, letter := range []byte(word) {
			morseWord += morseMap[letter]
			fmt.Printf("morseWord[%s]: %v \n", string(letter), morseWord)
		}

		morseSet[morseWord] = true
	}

	//fmt.Printf("===> morseSet: %v \n", morseSet)
	return len(morseSet)
}
