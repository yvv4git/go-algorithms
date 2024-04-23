package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	/*
		METHOD: Two pointers

		TIME COMPLEXITY: O(n), где n - длина входной строки, так как мы проходим по всей строке ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем получить новую строку, которая совпадает по длине с исходной.
	*/
	// Разделяем строку на слова
	words := strings.Split(s, " ")

	// Обходим каждое слово
	for i, word := range words {
		// Преобразуем слово в массив рун
		runes := []rune(word)
		// Обходим руны слова с двух концов и меняем их местами
		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
			runes[j], runes[k] = runes[k], runes[j]
		}
		// Заменяем слово в массиве слов на обращенное слово
		words[i] = string(runes)
	}

	// Объединяем слова обратно в строку с разделителем-пробелом
	return strings.Join(words, " ")
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s)) // Выводит: "s'teL ekat edoCteeL tsetnoc"
}
