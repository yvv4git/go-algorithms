//go:build ignore

package main

import "fmt"

// Реализация алгоритмов, таких как шифр Цезаря.
func caesarCipher(text string, shift int) string {
	result := ""
	for _, char := range text {
		if char >= 'a' && char <= 'z' {
			result += string('a' + (char-'a'+rune(shift))%26)
		} else if char >= 'A' && char <= 'Z' {
			result += string('A' + (char-'A'+rune(shift))%26)
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	text := "Hello"
	shift := 3
	fmt.Println("Зашифрованный текст:", caesarCipher(text, shift)) // Выведет "Khoor"
}
