package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция для подсчета частоты слов в файле
func countWordFrequency(filePath string) (map[string]int, error) {
	/*
		METHOD: Frequency
		Используемый в решении подход называется "подсчет частоты встречаемости слов".
		Этот подход основан на подсчете количества вхождений каждого слова в текстовом файле.

		TIME COMPLEXITY: O(n), где n - количество слов в файле, так как мы проходим по каждому слову один раз.

		SPACE COMPLEXITY: O(k), где k - количество уникальных слов в файле, так как мы храним частоту каждого слова в словаре.
	*/
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Инициализируем сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Инициализируем словарь для подсчета слов
	wordFrequency := make(map[string]int)

	// Читаем файл построчно
	for scanner.Scan() {
		// Получаем строку
		line := scanner.Text()

		// Разделяем строку на слова
		words := strings.Fields(line)

		// Подсчитываем частоту каждого слова
		for _, word := range words {
			wordFrequency[word]++
		}
	}

	// Проверяем наличие ошибок при сканировании файла
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return wordFrequency, nil
}

func main() {
	// Путь к файлу
	filePath := "example.txt"

	// Получаем частоту слов
	wordFrequency, err := countWordFrequency(filePath)
	if err != nil {
		fmt.Printf("Ошибка при подсчете частоты слов: %v\n", err)
		return
	}

	// Выводим частоту слов
	for word, frequency := range wordFrequency {
		fmt.Printf("Слово: %s, частота: %d\n", word, frequency)
	}
}
