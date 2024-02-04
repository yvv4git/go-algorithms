package main

import (
	"fmt"
)

func maximumOddBinaryNumber(s string) string {
	/*
		METHOD: Iteration
		Time complexity: O(n), т.к. надо пройтись по строке
		Space complexity: O(1), т.к. не используем доп. памяти, которая зависит от n

		Данный алгоритм решает задачу "Максимальное нечетное двоичное число" следующим образом:
		1. Считает количество единиц (бит, установленных в 1) в исходной строке.
		2. Создает новую строку, где все биты, кроме крайнего правого, устанавливаются в 1, если количество единиц больше 1.
		Если количество единиц равно 1 или меньше, то крайний правый бит также устанавливается в 1.
	*/
	strLen := len(s)     // Длина исходной строки
	var onesBitCount int // Счетчик установленных в 1 бит

	// Подсчитываем количество единиц в исходной строке
	for i := 0; i < strLen; i++ {
		if s[i] == '1' {
			onesBitCount++
		}
	}

	result := make([]byte, 0, strLen) // Создаем слайс байтов для результирующей строки

	// Формируем результирующую строку
	for i := 0; i < strLen; i++ {
		if onesBitCount > 1 || i == strLen-1 {
			// Если есть более одной единицы или это крайний правый бит, устанавливаем его в 1
			result = append(result, '1')
			onesBitCount--
		} else {
			// Иначе устанавливаем бит в 0
			result = append(result, '0')
		}
	}

	return string(result) // Возвращаем результирующую строку в виде строки
}

func main() {
	num := "3245"
	result := maximumOddBinaryNumber(num)
	fmt.Println(result)
}
