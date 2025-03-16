//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func restoreString(s string, indices []int) string {
	/*
		APPROACH: Using strings.Builder
		- Создаем `builder` длиной n.
		- Заполняем `builder` пустыми символами.
		- Проходим по `s` и `indices`, записывая символы `s[i]` в `builder` на позицию `indices[i]`.

		TIME COMPLEXITY: O(n)
		- Мы проходим по `s` и `indices` ровно один раз, выполняя O(1) операций для каждого символа.

		SPACE COMPLEXITY: O(n)
		- Используем `strings.Builder`, но без промежуточных срезов, что уменьшает накладные расходы.
	*/
	n := len(s)
	builder := strings.Builder{}
	builder.Grow(n) // Предварительно выделяем память

	result := make([]byte, n) // Создаем срез байтов
	for i, index := range indices {
		result[index] = s[i] // Расставляем символы
	}

	builder.WriteString(string(result)) // Записываем результат
	return builder.String()
}

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(s, indices)) // "leetcode"
}
