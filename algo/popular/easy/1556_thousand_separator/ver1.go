package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Примеры использования функции
	fmt.Println(thousandSeparator(987)) // "987"
}

func thousandSeparator(n int) string {
	/*
		APPROACH: Reverse String Construction
		- Преобразуем число в строку.
		- Идем с конца строки и вставляем символ '.' после каждых 3 цифр.
		- Собираем результат в буфер в обратном порядке и переворачиваем в конце.

		TIME COMPLEXITY: O(k)
		- Где k — количество цифр в числе (максимум 10 для int в Go).
		- Один проход по строке и один проход для реверса.

		SPACE COMPLEXITY: O(k)
		- Создаем новый буфер и массив рун для реверса строки.
	*/
	str := strconv.Itoa(n)
	var res strings.Builder

	count := 0
	for i := len(str) - 1; i >= 0; i-- {
		if count > 0 && count%3 == 0 {
			res.WriteByte('.')
		}
		res.WriteByte(str[i])
		count++
	}

	// Переворачиваем строку, так как записывали с конца
	result := []rune(res.String())
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
