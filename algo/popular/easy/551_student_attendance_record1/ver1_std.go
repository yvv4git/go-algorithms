//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func checkRecord(s string) bool {
	/*
		METHOD: STD lib

		TIME COMPLEXITY: O(n), так как в Go строки являются неизменяемыми,
		и при копировании подстроки или подсчете символов создается новая строка, потребляющая дополнительное место.

		SPACE COMPLEXITY: O(1), что означает, что алгоритм требует постоянного объема памяти,
		независимо от размера входных данных. Это связано с тем, что алгоритм не использует дополнительную память,
		которая бы растягивалась с увеличением размера входных данных.
	*/
	// Проверяем количество 'A' в записи
	if strings.Count(s, "A") > 1 {
		return false
	}
	// Проверяем наличие подстроки "LLL" в записи
	if strings.Contains(s, "LLL") {
		return false
	}
	return true
}

func main() {
	// Примеры использования
	fmt.Println(checkRecord("PPALLP")) // Вывод: true
	fmt.Println(checkRecord("PPALLL")) // Вывод: false
}
