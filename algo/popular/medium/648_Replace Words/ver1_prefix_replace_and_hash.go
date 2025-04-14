package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	// Создаем множество корней для быстрого поиска
	rootSet := make(map[string]bool)
	for _, root := range dictionary {
		rootSet[root] = true
	}

	// Разбиваем предложение на отдельные слова
	words := strings.Split(sentence, " ")

	// Обрабатываем каждое слово
	for i, word := range words {
		// Проверяем все возможные префиксы слова от самых коротких
		for l := 1; l <= len(word); l++ {
			prefix := word[:l]
			// Если префикс есть в множестве корней
			if rootSet[prefix] {
				// Заменяем слово на корень
				words[i] = prefix
				break // Прерываем цикл, так как нашли самый короткий корень
			}
		}
	}

	// Собираем слова обратно в предложение
	return strings.Join(words, " ")
}

func main() {
	// Тест 1
	dictionary1 := []string{"cat", "bat", "rat"}
	sentence1 := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary1, sentence1)) // Ожидаемый результат: "the cat was rat by the bat"

	// Тест 2
	dictionary2 := []string{"a", "b", "c"}
	sentence2 := "aadsfasf absbs bbab cadsfafs"
	fmt.Println(replaceWords(dictionary2, sentence2)) // Ожидаемый результат: "a a b c"

	// Тест 3 (дополнительный)
	dictionary3 := []string{"go", "golang", "prog"}
	sentence3 := "gopher went to program golang code"
	fmt.Println(replaceWords(dictionary3, sentence3)) // Ожидаемый результат: "go went to prog go code"
}
