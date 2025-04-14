//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	/*
		APPROACH: Prefix Replacement with Hash Set
		- Создаем хэш-множество корней для быстрого поиска O(1)
		- Разбиваем предложение на слова
		- Для каждого слова проверяем все возможные префиксы от коротких к длинным
		- При нахождении первого совпадающего префикса заменяем слово
		- Собираем обработанные слова в новое предложение

		TIME COMPLEXITY: O(N*M)
		- N - количество слов в предложении
		- M - средняя длина слова
		- В худшем случае для каждого слова проверяем все префиксы

		SPACE COMPLEXITY: O(K + N)
		- K - размер словаря (храним в хэш-множестве)
		- N - длина предложения (храним разбитые слова)
	*/
	// Создаем множество корней для быстрого поиска
	rootSet := make(map[string]struct{})
	for _, root := range dictionary {
		rootSet[root] = struct{}{}
	}

	words := strings.Split(sentence, " ")

	for i, word := range words {
		// Проверяем все возможные префиксы слова
		for l := 1; l <= len(word); l++ {
			prefix := word[:l]
			if _, exists := rootSet[prefix]; exists {
				words[i] = prefix // Заменяем на найденный корень
				break             // Используем первый найденный (самый короткий)
			}
		}
	}

	return strings.Join(words, " ")
}

func main() {
	// Тест 1 из условия задачи
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary, sentence)) // "the cat was rat by the bat"

	// Тест 2 из условия задачи
	dictionary = []string{"a", "b", "c"}
	sentence = "aadsfasf absbs bbab cadsfafs"
	fmt.Println(replaceWords(dictionary, sentence)) // "a a b c"

	// Дополнительный тест
	dictionary = []string{"go", "gol", "prog"}
	sentence = "gopher went to program golang code"
	fmt.Println(replaceWords(dictionary, sentence)) // "go went to prog gol code"
}
