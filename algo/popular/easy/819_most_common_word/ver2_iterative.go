package main

import (
	"fmt"
	"sort"
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
		- Создаем словарь для подсчета частоты слов.
		- Создаем список пар (слово, частота) и сортируем его по частоте в порядке убывания.
		- Проходим по отсортированному списку и возвращаем первое слово, которое не входит в список запрещенных слов.

		TIME COMPLEXITY:
		- Удаление пунктуации: O(n), где n — длина текста.
		- Разбиение текста на слова: O(n).
		- Подсчет частоты слов: O(n).
		- Сортировка списка пар: O(n log n).
		- Проверка на запрещенные слова: O(n * m), где m — количество запрещенных слов.
		Общая временная сложность: O(n log n + n * m).

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
	for _, word := range words {
		wordCount[word]++
	}

	// Создаем список пар (слово, частота)
	type wordFreq struct {
		word string
		freq int
	}
	var wordFreqList []wordFreq
	for word, freq := range wordCount {
		wordFreqList = append(wordFreqList, wordFreq{word, freq})
	}

	// Сортируем список по частоте в порядке убывания
	sort.Slice(wordFreqList, func(i, j int) bool {
		return wordFreqList[i].freq > wordFreqList[j].freq
	})

	// Создаем множество для хранения запрещенных слов
	bannedSet := make(map[string]bool)
	for _, word := range banned {
		bannedSet[word] = true
	}

	// Проходим по отсортированному списку и возвращаем первое слово, которое не входит в список запрещенных слов
	for _, wf := range wordFreqList {
		if !bannedSet[wf.word] {
			return wf.word
		}
	}

	// Если все слова запрещены, возвращаем пустую строку
	return ""
}

func main() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}

	result := mostCommonWordV2(paragraph, banned)
	fmt.Println("Наиболее часто встречающееся слово:", result) // Вывод: ball
}
