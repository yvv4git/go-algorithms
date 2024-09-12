package main

import (
	"fmt"
	"strings"
)

// Функция для подсчета частоты букв в строке
func countLetters(s string) [26]int {
	/*
		METHOD: Подсчет частоты букв в строке
		TIME COMPLEXITY: O(n), где n - длина строки s
		SPACE COMPLEXITY: O(1), так как возвращаемый массив фиксированного размера (26 элементов)
	*/
	count := [26]int{}
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			count[char-'a']++
		}
	}
	return count
}

// Функция для проверки, содержит ли слово все буквы из licensePlate в нужном количестве
func isCompletingWord(word string, licenseCount [26]int) bool {
	/*
		METHOD: Проверка соответствия слова требованиям licensePlate
		TIME COMPLEXITY: O(m), где m - длина слова word
		SPACE COMPLEXITY: O(1), так как используем фиксированный массив размера 26
	*/
	wordCount := countLetters(word)
	for i := 0; i < 26; i++ {
		if wordCount[i] < licenseCount[i] {
			return false
		}
	}
	return true
}

// Основная функция решения задачи
func shortestCompletingWord(licensePlate string, words []string) string {
	/*
		METHOD: Iterative
		Поиск самого короткого слова, содержащего все буквы из licensePlate
		TIME COMPLEXITY: O(n + k * m), где n - длина licensePlate, k - количество слов в words, m - средняя длина слов в words
		SPACE COMPLEXITY: O(1), так как используем фиксированные массивы размера 26
	*/
	// Приводим licensePlate к нижнему регистру и удаляем все небуквенные символы
	licensePlate = strings.ToLower(licensePlate)
	licensePlate = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r
		}
		return -1
	}, licensePlate)

	// Подсчитываем частоту букв в licensePlate
	licenseCount := countLetters(licensePlate)

	// Инициализируем переменную для хранения самого короткого слова
	var shortestWord string

	// Проходим по всем словам и ищем самое короткое, которое содержит все буквы из licensePlate
	for _, word := range words {
		if isCompletingWord(word, licenseCount) {
			if shortestWord == "" || len(word) < len(shortestWord) {
				shortestWord = word
			}
		}
	}

	return shortestWord
}

func main() {
	licensePlate := "1s3 PSt"
	words := []string{"step", "steps", "stripe", "stepple"}
	result := shortestCompletingWord(licensePlate, words)
	fmt.Println(result) // Вывод: "steps"
}
