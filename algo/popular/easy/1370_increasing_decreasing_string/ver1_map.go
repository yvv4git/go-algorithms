package main

import (
	"fmt"
)

func sortString(s string) string {
	/*
		METHOD: Map and Sort

		Пусть есть строка s. Необходимо сортировать эту строку в порядке возрастания.

		Подход:
		1. Создать массив count[26] для подсчета количества каждого символа.
		2. Подсчитать количество каждого символа в строке.
		3. Создать слайс result для хранения результата.
		4. Пока длина результата меньше длины исходной строки:
			- Добавлять символы в порядке возрастания (от 'a' до 'z').
			- Добавлять символы в порядке убывания (от 'z' до 'a').
		5. Преобразовать слайс result в строку и возвращать результат.

		TIME COMPLEXITY: O(n), где n — длина строки.
		SPACE COMPLEXITY: O(1), не зависит от длины строки.
	*/
	// Создаем массив для подсчета количества каждого символа
	// Индекс 0 соответствует 'a', индекс 1 — 'b', и так далее до 'z' (индекс 25)
	count := make([]int, 26)

	// Подсчитываем количество каждого символа в строке
	for _, char := range s {
		count[char-'a']++
	}

	// Создаем слайс для хранения результата
	result := []rune{}

	// Пока длина результата меньше длины исходной строки
	for len(result) < len(s) {
		// Добавляем символы в порядке возрастания (от 'a' до 'z')
		for i := 0; i < 26; i++ {
			if count[i] > 0 { // Если символ еще остался
				result = append(result, rune(i+'a')) // Добавляем символ в результат
				count[i]--                           // Уменьшаем количество оставшихся символов
			}
		}

		// Добавляем символы в порядке убывания (от 'z' до 'a')
		for i := 25; i >= 0; i-- {
			if count[i] > 0 { // Если символ еще остался
				result = append(result, rune(i+'a')) // Добавляем символ в результат
				count[i]--                           // Уменьшаем количество оставшихся символов
			}
		}
	}

	// Преобразуем слайс рун в строку и возвращаем результат
	return string(result)
}

func main() {
	// Пример использования
	s := "aaaabbbbcccc"
	fmt.Println("Исходная строка:", s)
	fmt.Println("Результат:", sortString(s))
}
