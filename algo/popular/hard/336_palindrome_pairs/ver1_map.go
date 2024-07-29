package main

import (
	"fmt"
	"strings"
)

// Функция для проверки, является ли строка палиндромом
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// Основная функция для поиска всех пар индексов, образующих палиндром
func palindromePairs(words []string) [][]int {
	/*
		METHOD: Map / Dict
		Я буду использовать метод, основанный на хэш-таблице (map) для быстрого доступа к строкам и их индексам.
		Этот метод позволяет эффективно проверять возможные пары строк, которые могут образовать палиндром.
		Метод решения:
		1. Хэш-таблица (map): Используется для быстрого доступа к индексам строк. Это позволяет проверять наличие строк и их инверсий за O(1) времени.
		2. Проверка палиндромов: Для каждой строки и её подстрок проверяется, являются ли они палиндромами.
		3. Обратные строки: Проверяются обратные строки и их наличие в хэш-таблице.

		TIME COMPLEXITY: O(n * k^2), где n — количество строк, а k — средняя длина строки.
		Это связано с тем, что для каждой строки мы проверяем все возможные подстроки и их инверсии.

		SPACE COMPLEXITY: O(n), где n - количество строк. Это связано с использованием хэш-таблицы для хранения строк и их индексов.
	*/
	// Создаем словарь для хранения строк и их индексов
	wordMap := make(map[string]int)
	for i, word := range words {
		wordMap[word] = i
	}

	// Результирующий список пар индексов
	var result [][]int

	// Проверяем каждую строку и её возможные подстроки
	for i, word := range words {
		// Случай 1: Пустая строка и строка-палиндром
		if idx, ok := wordMap[""]; ok && idx != i && isPalindrome(word) {
			result = append(result, []int{i, idx})
			result = append(result, []int{idx, i})
		}

		// Случай 2: Обратная строка существует
		reverseWord := reverseString(word)
		if idx, ok := wordMap[reverseWord]; ok && idx != i {
			result = append(result, []int{i, idx})
		}

		// Случай 3: Проверка всех возможных подстрок
		for j := 1; j < len(word); j++ {
			prefix := word[:j]
			suffix := word[j:]

			// Проверка префикса
			if isPalindrome(prefix) {
				reverseSuffix := reverseString(suffix)
				if idx, ok := wordMap[reverseSuffix]; ok && idx != i {
					result = append(result, []int{idx, i})
				}
			}

			// Проверка суффикса
			if isPalindrome(suffix) {
				reversePrefix := reverseString(prefix)
				if idx, ok := wordMap[reversePrefix]; ok && idx != i {
					result = append(result, []int{i, idx})
				}
			}
		}
	}

	return result
}

// Функция для реверса строки
func reverseString(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func main() {
	words := []string{"abcd", "dcba", "lls", "s", "sssll"}
	result := palindromePairs(words)
	fmt.Println(result) // Вывод: [[0 1] [1 0] [3 2] [2 4]]
}
