package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	/*
		METHOD: Iterate and Counting
		TIME COMPLEXITY: O(n), где n - количество символов в строке s. Это связано с тем, что мы проходим по всей строке, разделяя ее на сегменты.
		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем получить n+1 сегментов, если между каждыми двумя словами будет один пробел.
	*/
	// Удаляем пробелы в начале и в конце строки
	s = strings.TrimSpace(s)

	// Если строка пустая, то сегментов нет
	if len(s) == 0 {
		return 0
	}

	// Разделяем строку на сегменты и считаем их количество
	segments := strings.Split(s, " ")
	count := 0
	for _, segment := range segments {
		if segment != "" {
			count++
		}
	}

	return count
}

func main() {
	s := "Hello, my name is John"
	fmt.Println(countSegments(s)) // Выводит: 5
}
