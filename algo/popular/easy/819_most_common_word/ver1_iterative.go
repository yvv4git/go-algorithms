//go:build ignore

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	/*
		METHOD:
		- Приводим весь текст к нижнему регистру.
		- Разбиваем текст на слова, используя функцию `strings.FieldsFunc`, которая игнорирует символы, не являющиеся буквами.
		- Создаем словарь `wordsMap` для подсчета частоты каждого слова.
		- Создаем множество `bannedWords` для хранения запрещенных слов.
		- Проходим по всем словам в `wordsMap`, проверяем, что слово не запрещено, и определяем слово с максимальной частотой.

		TIME COMPLEXITY:
		- Приведение текста к нижнему регистру: O(n), где n — длина текста.
		- Разбиение текста на слова: O(n).
		- Подсчет частоты слов: O(n).
		- Создание множества запрещенных слов: O(m), где m — количество запрещенных слов.
		- Поиск наиболее частого слова: O(n).
		Общая временная сложность: O(n + m).

		SPACE COMPLEXITY:
		- Хранение слов и их частот: O(n).
		- Хранение запрещенных слов: O(m).
		Общая пространственная сложность: O(n + m).
	*/

	// Приводим весь текст к нижнему регистру
	paragraph = strings.ToLower(paragraph)

	// Разбиваем текст на слова, используя функцию `strings.FieldsFunc`, которая игнорирует символы, не являющиеся буквами
	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	// Создаем словарь `wordsMap` для подсчета частоты каждого слова
	wordsMap := map[string]int{}
	for _, word := range words {
		wordsMap[word]++
	}

	// Создаем множество `bannedWords` для хранения запрещенных слов
	bannedWords := map[string]bool{}
	for _, word := range banned {
		bannedWords[word] = true
	}

	// Структура для хранения наиболее часто встречающегося слова и его частоты
	highest := struct {
		word  string
		count int
	}{}

	// Проходим по всем словам в `wordsMap`, проверяем, что слово не запрещено, и определяем слово с максимальной частотой
	for word, count := range wordsMap {
		if bannedWords[word] {
			continue
		}
		if count > highest.count {
			highest.word = word
			highest.count = count
		}
	}

	return highest.word
}

func main() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}

	result := mostCommonWord(paragraph, banned)
	fmt.Println("Наиболее часто встречающееся слово:", result) // Вывод: ball
}
