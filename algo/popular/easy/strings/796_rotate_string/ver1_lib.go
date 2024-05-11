package main

import (
	"strings"
)

// Функция для проверки, является ли строка B циклическим сдвигом строки A
func rotateString(s string, goal string) bool {
	/*
		METHOD: Lib strings
		Выбранный подход основан на использовании функции strings.Contains.
		Мы проверяем, является ли строка B циклическим сдвигом строки A, просто дублируя строку A и проверяя,
		входит ли B в полученную строку. Если B входит, то она является циклическим сдвигом A.
		Этот подход работает, потому что любой циклический сдвиг строки A будет находиться в строке A+A.

		TIME COMPLEXITY: O(n), где n - длина строк A и B.
		Это связано с тем, что функция strings.Contains в Go работает за линейное время в худшем случае, когда подстрока не найдена.

		SPACE COMPLEXITY: O(n), так как мы создаем дополнительную строку, которая содержит дубликат строки A.
	*/
	// Проверяем, что длины строк равны
	if len(s) != len(goal) {
		return false
	}

	// Если строки равны, то они являются циклическим сдвигом друг друга
	if s == goal {
		return true
	}

	// Используем функцию strings.Contains для проверки, является ли строка goal циклическим сдвигом строки s
	// Строка goal является циклическим сдвигом строки s, если она входит в саму себя, просто сдвинутую на какое-то количество символов
	return strings.Contains(s+s, goal)
}

//func main() {
//	A := "abcde"
//	B := "cdeab"
//	fmt.Println(rotateString(A, B)) // Вывод: true
//}
