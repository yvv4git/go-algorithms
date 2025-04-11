package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	/*
		APPROACH: Sliding Window with Hash Maps
		1. Проверяем крайние случаи (пустые входные данные).
		2. Вычисляем длину одного слова и общую длину всех слов.
		3. Создаем эталонную мапу для подсчета частот слов.
		4. Перебираем все возможные начальные позиции подстроки:
		   - Для каждой позиции проверяем подстроку длиной totalLen
		   - Разбиваем подстроку на слова фиксированной длины
		   - Сравниваем частоты слов с эталоном
		5. Если частоты совпадают, добавляем индекс в результат.

		TIME COMPLEXITY: O(n * m * k)
		- n - длина строки s
		- m - количество слов в words
		- k - длина каждого слова
		- Внешний цикл выполняется O(n) раз
		- Внутренний цикл выполняется O(m) раз для каждой позиции
		- Извлечение подстроки занимает O(k)
		- Итоговая сложность O(n * m * k)

		SPACE COMPLEXITY: O(m)
		- wordCount мапа занимает O(m) памяти
		- seen мапа в худшем случае занимает O(m)
		- Результирующий слайс в худшем случае O(n/m)
		- Итоговая сложность O(m) (доминирующий фактор)
	*/
	// Если строка или слова пустые, возвращаем пустой результат
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	// Длина одного слова (все слова одинаковой длины)
	wordLen := len(words[0])
	// Общая длина всех слов вместе
	totalLen := wordLen * len(words)
	// Длина строки
	sLen := len(s)
	// Результат - массив начальных индексов
	result := []int{}

	// Если общая длина слов больше длины строки, возвращаем пустой результат
	if totalLen > sLen {
		return result
	}

	// Создаем мапу для подсчета слов (эталон)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// Перебираем все возможные начальные позиции
	for i := 0; i <= sLen-totalLen; i++ {
		// Создаем временную мапу для проверки текущей подстроки
		seen := make(map[string]int)
		j := 0
		// Проверяем все слова в текущей подстроке
		for j < len(words) {
			// Вычисляем начальную позицию текущего слова
			start := i + j*wordLen
			// Извлекаем текущее слово
			currentWord := s[start : start+wordLen]
			// Если слово есть в эталонной мапе
			if count, exists := wordCount[currentWord]; exists {
				seen[currentWord]++
				// Если встретили слово больше раз, чем в эталоне - прерываем проверку
				if seen[currentWord] > count {
					break
				}
			} else {
				// Если слова нет в эталоне - прерываем проверку
				break
			}
			j++
		}
		// Если все слова совпали с эталоном, добавляем индекс в результат
		if j == len(words) {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	// Пример использования
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words)) // Вывод: [0 9]

	s = "wordgoodgoodgoodbestword"
	words = []string{"word", "good", "best", "word"}
	fmt.Println(findSubstring(s, words)) // Вывод: []

	s = "barfoofoobarthefoobarman"
	words = []string{"bar", "foo", "the"}
	fmt.Println(findSubstring(s, words)) // Вывод: [6 9 12]
}
