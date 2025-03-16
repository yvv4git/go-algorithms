//go:build ignore

package main

import (
	"fmt"
)

func restoreString(s string, indices []int) string {
	/*
		APPROACH: Index Mapping
		- Создаем срез `result` длины n (где n — длина строки s), который будет содержать переставленные символы.
		- Проходим по `s` и `indices`, помещая каждый символ `s[i]` в позицию `indices[i]` в `result`.
		- Преобразуем `result` обратно в строку и возвращаем.

		TIME COMPLEXITY: O(n)
		- Мы проходим по `s` и `indices` ровно один раз, выполняя O(1) операций для каждого элемента.
		- Таким образом, общее время работы составляет O(n), где n — длина строки.

		SPACE COMPLEXITY: O(n)
		- Используется дополнительный массив `result` размером n для хранения переставленных символов.
		- Итоговое использование памяти O(n), так как размер `result` зависит от входных данных.
	*/
	result := make([]byte, len(s)) // Создаем слайс байтов нужной длины
	for i, index := range indices {
		result[index] = s[i] // Расставляем символы на нужные позиции
	}

	return string(result) // Преобразуем байты обратно в строку
}

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(s, indices)) // "leetcode"
}
