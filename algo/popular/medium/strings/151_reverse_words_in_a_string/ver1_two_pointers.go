package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	/*
		METHOD: Iterative & Two pointers
		Time complexity: O(n), где n - количество слов в строке. Это связано с тем, что мы проходим по каждому слову в строке один раз.
		Space complexity: O(n), так как мы создаем новый массив для хранения слов в строке. В худшем случае это будет n, если каждое слово в строке односимвольное.
	*/
	// Разделяем строку на слова
	words := strings.Fields(s)

	// Меняем местами слова в массиве
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединяем слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s)) // Выводит: "blue is sky the"
}
