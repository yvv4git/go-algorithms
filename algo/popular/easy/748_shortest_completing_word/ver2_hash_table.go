package main

import (
	"strings"
)

// Функция для подсчета частоты букв в строке с использованием хеш-таблицы
func countLettersMap(s string) map[rune]int {
	count := make(map[rune]int)
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			count[char]++
		}
	}
	return count
}

// Функция для проверки, содержит ли слово все буквы из licensePlate в нужном количестве с использованием хеш-таблицы
func isCompletingWordMap(word string, licenseCount map[rune]int) bool {
	wordCount := countLettersMap(word)
	for char, requiredCount := range licenseCount {
		if wordCount[char] < requiredCount {
			return false
		}
	}
	return true
}

// Основная функция решения задачи с использованием хеш-таблиц
func shortestCompletingWordV2(licensePlate string, words []string) string {
	/*
		METHOD: Hash table
		Использование хеш-таблиц для подсчета частоты букв и проверки соответствия слов
		TIME COMPLEXITY: O(n + k * m), где n - длина licensePlate, k - количество слов в words, m - средняя длина слов в words
		SPACE COMPLEXITY: O(1), так как используем фиксированные хеш-таблицы размера 26
	*/
	// Приводим licensePlate к нижнему регистру и удаляем все небуквенные символы
	licensePlate = strings.ToLower(licensePlate)
	licensePlate = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r
		}
		return -1
	}, licensePlate)

	// Подсчитываем частоту букв в licensePlate с использованием хеш-таблицы
	licenseCount := countLettersMap(licensePlate)

	// Инициализируем переменную для хранения самого короткого слова
	var shortestWord string

	// Проходим по всем словам и ищем самое короткое, которое содержит все буквы из licensePlate
	for _, word := range words {
		if isCompletingWordMap(word, licenseCount) {
			if shortestWord == "" || len(word) < len(shortestWord) {
				shortestWord = word
			}
		}
	}

	return shortestWord
}

//func main() {
//	licensePlate := "1s3 PSt"
//	words := []string{"step", "steps", "stripe", "stepple"}
//	result := shortestCompletingWordV2(licensePlate, words)
//	fmt.Println(result) // Вывод: "steps"
//}
