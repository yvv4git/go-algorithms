package main

import (
	"fmt"
	"unicode"
)

func reverseOnlyLetters(s string) string {
	/*
		METHOD: Two pointers
		В этом коде мы используем функцию unicode.IsLetter для проверки, является ли символ буквой.
		Если символ не является буквой, мы просто переходим к следующему или предыдущему символу.
		Если символ является буквой, мы меняем его местами с символом, находящимся на противоположной стороне от начала или конца строки.
		Таким образом, буквы в строке перевернуты, а другие символы остаются на своих местах.

		TIME COMPLEXITY: O(n), где n - длина входной строки, так как мы проходим по всей строке только один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем получить новую строку, которая совпадает с исходной.
	*/
	// Конвертируем строку в массив рун для работы с символами
	runes := []rune(s)

	// Инициализируем два указателя: один на начало строки, другой на конец
	left, right := 0, len(runes)-1

	// Проходим по строке с двух концов
	for left < right {
		// Если левый символ не является буквой, переходим к следующему
		if !unicode.IsLetter(runes[left]) {
			left++
			continue
		}

		// Если правый символ не является буквой, переходим к предыдущему
		if !unicode.IsLetter(runes[right]) {
			right--
			continue
		}

		// Меняем местами левый и правый символы
		runes[left], runes[right] = runes[right], runes[left]

		// Перемещаем указатели к центру
		left++
		right--
	}

	// Конвертируем массив рун обратно в строку
	return string(runes)
}

func main() {
	// Пример использования функции
	s := "ab-cd"
	result := reverseOnlyLetters(s)
	fmt.Println(result) // Вывод: "dc-ba"
}
