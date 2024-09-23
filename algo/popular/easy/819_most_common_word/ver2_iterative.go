//go:build ignore

package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Функция для удаления пунктуации из строки
func removePunctuation(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, s)
}

// Функция для поиска наиболее часто встречающегося слова в тексте, не входящего в список запрещенных слов
func mostCommonWordV2(paragraph string, banned []string) string {
	/*
		METHOD:
		- Удаляем пунктуацию из текста.
		- Разбиваем текст на слова и приводим их к нижнему регистру.
		- Создаем словарь для подсчета частоты слов, игнорируя запрещенные слова.
		- Проходим по всем словам и обновляем счетчик частоты, если слово не запрещено.
		- Определяем слово с максимальной частотой.

		TIME COMPLEXITY:
		- Удаление пунктуации: O(n), где n — длина текста.
		- Разбиение текста на слова: O(n).
		- Подсчет частоты слов: O(n).
		- Проверка на запрещенные слова: O(n * m), где m — количество запрещенных слов.
		Общая временная сложность: O(n * m).

		SPACE COMPLEXITY:
		- Хранение слов и их частот: O(n).
		- Хранение запрещенных слов: O(m).
		Общая пространственная сложность: O(n + m).
	*/

	// Удаляем пунктуацию из текста
	paragraph = removePunctuation(paragraph)

	// Разбиваем текст на слова и приводим их к нижнему регистру
	words := strings.Fields(strings.ToLower(paragraph))

	// Создаем словарь для подсчета частоты слов
	wordCount := make(map[string]int)

	// Создаем множество для хранения запрещенных слов
	bannedSet := make(map[string]bool)
	for _, word := range banned {
		bannedSet[word] = true
	}

	// Проходим по всем словам и обновляем счетчик частоты, если слово не запрещено
	for _, word := range words {
		if !bannedSet[word] {
			wordCount[word]++
		}
	}

	// Определяем слово с максимальной частотой
	maxCount := 0
	mostCommon := ""
	for word, count := range wordCount {
		if count > maxCount {
			maxCount = count
			mostCommon = word
		}
	}

	return mostCommon
}

func main() {
	paragraph := "a, a, a, a, b,b,b,c, c"
	banned := []string{"a"}

	result := mostCommonWordV2(paragraph, banned)
	fmt.Println("Наиболее часто встречающееся слово:", result) // Вывод: b
}
