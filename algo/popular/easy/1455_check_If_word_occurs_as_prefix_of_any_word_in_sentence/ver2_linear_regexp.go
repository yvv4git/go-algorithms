//go:build ignore

package main

import (
	"regexp"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	/*
	   METHOD: Regular Expression
	   - Используем регулярное выражение для поиска слов с нужным префиксом.
	   - Менее эффективен по производительности, но демонстрирует альтернативный подход.
	*/
	re := regexp.MustCompile(`\b` + regexp.QuoteMeta(searchWord) + `\w*\b`)
	words := strings.Split(sentence, " ")

	for i, word := range words {
		if re.FindStringIndex(word) != nil && strings.HasPrefix(word, searchWord) {
			return i + 1
		}
	}

	return -1
}

func main() {
	// Примеры использования функции
	fmt.Println(isPrefixOfWord("i love eating burger", "burg"))           // 4
	fmt.Println(isPrefixOfWord("this problem is an easy problem", "pro")) // 2
	fmt.Println(isPrefixOfWord("i am tired", "you"))                      // -1
}
