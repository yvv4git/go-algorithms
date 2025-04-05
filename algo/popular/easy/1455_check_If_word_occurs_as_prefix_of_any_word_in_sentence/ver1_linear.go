package main

import (
	"fmt"
	"strings"
)

// Функция проверяет, является ли searchWord префиксом любого слова в предложении.
// Возвращает индекс первого подходящего слова (начиная с 1) или -1, если таких слов нет.
func isPrefixOfWord(sentence string, searchWord string) int {
	/*
		METHOD: Linear Search with String Prefix Check
		- Разбиваем предложение на слова по пробелам.
		- Последовательно проверяем каждое слово, начинается ли оно с searchWord.
		- Возвращаем первое найденное совпадение (1-based индекс).
		- Если совпадений нет, возвращаем -1.

		TIME COMPLEXITY: O(n*m)
		- Разбиение строки: O(n), где n - длина предложения.
		- Проверка каждого слова: O(m) на слово, где m - длина searchWord.
		- В худшем случае проверяем все слова, поэтому общая сложность O(n*m).

		SPACE COMPLEXITY: O(n)
		- Хранение разбитых слов требует O(n) дополнительной памяти.
		- Не используем дополнительную память, пропорциональную размеру входа.
	*/
	// Разделяем предложение на слова по пробелам
	words := strings.Split(sentence, " ")

	// Перебираем каждое слово и его индекс
	for i, word := range words {
		// Проверяем, начинается ли текущее слово с searchWord
		if strings.HasPrefix(word, searchWord) {
			// Возвращаем индекс слова (1-based)
			return i + 1
		}
	}

	// Если ни одно слово не подошло, возвращаем -1
	return -1
}

func main() {
	// Примеры использования функции
	fmt.Println(isPrefixOfWord("i love eating burger", "burg"))           // 4
	fmt.Println(isPrefixOfWord("this problem is an easy problem", "pro")) // 2
	fmt.Println(isPrefixOfWord("i am tired", "you"))                      // -1
}
