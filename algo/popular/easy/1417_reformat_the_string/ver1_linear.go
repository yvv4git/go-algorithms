package main

import (
	"fmt"
	"unicode"
)

func reformat(s string) string {
	/*
		METHOD: Linear scan
		1. Создаем два слайса для хранения букв и цифр.
		2. Проходим по каждому символу строки, разделяя их на буквы и цифры.
		3. Проверяем, можно ли переформатировать строку:
		   - Если разница в количестве букв и цифр больше 1, возвращаем пустую строку.
		4. Создаем слайс для результата.
		5. Определяем, с чего начать (буквы или цифры), в зависимости от того, чего больше.
		6. Чередуем элементы двух слайсов с помощью функции interleave.
		7. Преобразуем слайс рун в строку и возвращаем результат.

		TIME COMPLEXITY: O(n)
		   - Мы проходим по строке один раз для разделения символов (O(n)).
		   - Чередование символов также выполняется за O(n).
		   - Итоговая сложность: O(n).

		SPACE COMPLEXITY: O(n)
		   - Мы используем дополнительные слайсы для хранения букв и цифр, которые в худшем случае занимают O(n) памяти.
		   - Результирующая строка также занимает O(n) памяти.
		   - Итоговая сложность: O(n).
	*/
	// Создаем два слайса для хранения букв и цифр
	var letters, digits []rune

	// Разделяем символы строки на буквы и цифры
	for _, char := range s {
		if unicode.IsLetter(char) {
			letters = append(letters, char)
		} else if unicode.IsDigit(char) {
			digits = append(digits, char)
		}
	}

	// Проверяем, можно ли переформатировать строку
	// Если разница в количестве букв и цифр больше 1, возвращаем пустую строку
	if abs(len(letters)-len(digits)) > 1 {
		return ""
	}

	// Создаем слайс для результата
	var result []rune

	// Определяем, с чего начать (буквы или цифры)
	if len(letters) > len(digits) {
		// Если букв больше, начинаем с буквы
		result = interleave(letters, digits)
	} else {
		// Иначе начинаем с цифры
		result = interleave(digits, letters)
	}

	// Преобразуем слайс рун в строку и возвращаем результат
	return string(result)
}

// Функция для чередования элементов двух слайсов
func interleave(first, second []rune) []rune {
	var result []rune
	for i := 0; i < len(first); i++ {
		result = append(result, first[i])
		if i < len(second) {
			result = append(result, second[i])
		}
	}
	return result
}

// Функция для вычисления абсолютного значения разницы
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Примеры использования
	fmt.Println(reformat("a0b1c2"))     // "a0b1c2"
	fmt.Println(reformat("leetcode"))   // ""
	fmt.Println(reformat("1229857369")) // ""
	fmt.Println(reformat("covid2019"))  // "c2o0v1i9d" или другой допустимый вариант
}
