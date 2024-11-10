package main

import (
	"fmt"
)

// Функция для сравнения двух строк после обработки символов backspace
func backspaceCompare(s string, t string) bool {
	/*
		METHOD: Stack
		Мы используем стек для обработки каждой строки. Проходим по каждому символу строки:
		- Если символ не '#', добавляем его в стек.
		- Если символ '#', удаляем последний символ из стека, если стек не пуст.
		После обработки обеих строк сравниваем содержимое стеков.

		Time complexity : O(n)
		- Мы проходим по каждому символу строки один раз, что занимает O(n) времени.
		- Добавление и удаление элементов из стека занимает O(1) времени.
		- Общая временная сложность составляет O(n + m), где n и m — длины строк s и t.

		Space complexity : O(n)
		- Мы используем стек для хранения символов строки, что требует дополнительной памяти.
		- В худшем случае (когда нет символов '#'), стек будет содержать все символы строки.
		- Общая пространственная сложность составляет O(n + m).
	*/
	// Вспомогательная функция для обработки строки с использованием стека
	processString := func(str string) string {
		stack := []rune{}
		for _, char := range str {
			if char != '#' {
				stack = append(stack, char) // Добавляем символ в стек, если это не '#'
			} else if len(stack) > 0 {
				stack = stack[:len(stack)-1] // Удаляем последний символ из стека, если это '#'
			}
		}
		return string(stack) // Возвращаем строку после обработки
	}

	// Обрабатываем обе строки и сравниваем результаты
	return processString(s) == processString(t)
}

func main() {
	s := "ab#c"
	t := "ad#c"
	fmt.Println(backspaceCompare(s, t)) // Вывод: true

	s = "ab##"
	t = "c#d#"
	fmt.Println(backspaceCompare(s, t)) // Вывод: true

	s = "a##c"
	t = "#a#c"
	fmt.Println(backspaceCompare(s, t)) // Вывод: true

	s = "a#c"
	t = "b"
	fmt.Println(backspaceCompare(s, t)) // Вывод: false
}
