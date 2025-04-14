//go:build ignore

package main

import (
	"fmt"
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	/*
		APPROACH: Sorted Roots with Early Termination
		- Сортируем корни по длине (от коротких к длинным)
		- Разбиваем предложение на слова
		- Для каждого слова проверяем корни в порядке сортировки
		- При первом совпадении заменяем слово и прекращаем проверку
		- Собираем обработанные слова в новое предложение

		TIME COMPLEXITY: O(DlogD + N*M)
		- DlogD - сортировка словаря
		- N - количество слов в предложении
		- M - среднее количество проверяемых корней для слова
		- В лучшем случае (ранний выход) сложность ближе к O(DlogD + N)

		SPACE COMPLEXITY: O(D + N)
		- D - хранение отсортированного словаря
		- N - хранение разбитых слов
	*/
	// Сортируем корни по длине (от самых коротких к длинным)
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})

	words := strings.Split(sentence, " ")

	for i, word := range words {
		// Проверяем корни в порядке возрастания длины
		for _, root := range dictionary {
			if strings.HasPrefix(word, root) {
				words[i] = root // Заменяем на первый подходящий корень
				break           // Прекращаем проверку для этого слова
			}
		}
	}

	return strings.Join(words, " ")
}

func main() {
	// Тест 1: базовый случай
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Printf("Input: %v\nOutput: %q\n\n", dictionary, replaceWords(dictionary, sentence))

	// Тест 2: несколько возможных корней
	dictionary = []string{"a", "aa", "aaa", "aaaa"}
	sentence = "a aa aaa aaaa aaaaa"
	fmt.Printf("Input: %v\nOutput: %q\n\n", dictionary, replaceWords(dictionary, sentence))

	// Тест 3: случай без замен
	dictionary = []string{"abc", "def", "ghi"}
	sentence = "hello world go"
	fmt.Printf("Input: %v\nOutput: %q\n\n", dictionary, replaceWords(dictionary, sentence))
}
