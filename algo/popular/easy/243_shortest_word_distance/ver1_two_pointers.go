package main

import (
	"fmt"
	"math"
)

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	/*
		METHOD: Two Pointers
		Мы используем два указателя (индекса) для отслеживания последних позиций, на которых встретились слова word1 и word2.
		При проходе по массиву слов мы обновляем эти указатели и вычисляем текущее расстояние между ними.
		Если текущее расстояние меньше минимального, которое мы запомнили, обновляем минимальное расстояние.

		TIME COMPLEXITY: O(n)
		Мы проходим по массиву слов один раз, что занимает O(n) времени, где n — количество слов в массиве.

		SPACE COMPLEXITY: O(1)
		Мы используем только несколько дополнительных переменных (index1, index2, minDistance), что занимает постоянное количество памяти O(1).
	*/

	// Инициализируем переменные для хранения индексов последних встреч слов word1 и word2
	index1 := -1
	index2 := -1

	// Инициализируем переменную для хранения минимального расстояния
	minDistance := math.MaxInt32

	// Проходим по массиву слов
	for i, word := range wordsDict {
		// Если текущее слово равно word1, обновляем index1
		if word == word1 {
			index1 = i
		}
		// Если текущее слово равно word2, обновляем index2
		if word == word2 {
			index2 = i
		}

		// Если оба индекса установлены (т.е. оба слова уже встречались)
		if index1 != -1 && index2 != -1 {
			// Вычисляем текущее расстояние между индексами
			currentDistance := int(math.Abs(float64(index1 - index2)))
			// Обновляем минимальное расстояние, если текущее расстояние меньше
			if currentDistance < minDistance {
				minDistance = currentDistance
			}
		}
	}

	// Возвращаем минимальное расстояние
	return minDistance
}

func main() {
	// Пример массива слов и слов для поиска
	wordsDict := []string{"practice", "makes", "perfect", "coding", "makes"}
	word1 := "coding"
	word2 := "practice"

	// Вызываем функцию и выводим результат
	result := shortestDistance(wordsDict, word1, word2)
	fmt.Println("Кратчайшее расстояние:", result)
}
