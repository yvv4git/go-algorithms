//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Тестовые случаи с разными комбинациями цифр
	testCases := []struct {
		input    string
		expected string
	}{
		{"owoztneoer", "012"},   // zero, one, two
		{"fviefuro", "45"},      // four, five
		{"xsis", "66"},          // two six
		{"nnei", "9"},           // nine
		{"zerozero", "00"},      // два zero
		{"twotwo", "22"},        // два two
		{"esnvei", "79"},        // seven, nine
		{"", ""},                // пустая строка
		{"oneoneone", "111"},    // три one
		{"eightwothree", "238"}, // two, three, eight
	}

	// Проверка всех тестовых случаев
	for _, tc := range testCases {
		result := originalDigits(tc.input)
		fmt.Printf("Input: %-15s Expected: %-8s Got: %-8s Correct: %t\n",
			"\""+tc.input+"\"",
			"\""+tc.expected+"\"",
			"\""+result+"\"",
			result == tc.expected)
	}
}

func originalDigits(s string) string {
	/*
		APPROACH: Hashing + Pattern Matching
		- Используем хеш-таблицу для подсчета частот букв (counts[26])
		- Обрабатываем цифры в порядке их уникальных букв:
		  0 (z), 2 (w), 4 (u), 6 (x), 8 (g)
		- Затем обрабатываем зависимые цифры, вычитая уже учтенные буквы:
		  1 (o), 3 (h), 5 (f), 7 (s), 9 (i)
		- Собираем результат в порядке возрастания цифр

		TIME COMPLEXITY: O(N)
		- Подсчет букв: O(n)
		- Обработка цифр: O(1) (фиксированные 10 цифр)
		- Построение результата: O(n) в худшем случае

		SPACE COMPLEXITY: O(1)
		- counts[26] и digits[10] - константные структуры
		- Результирующая строка не учитывается (это output)
	*/
	counts := make([]int, 26)
	for _, c := range s {
		counts[c-'a']++
	}

	digits := make([]int, 10)

	// Обрабатываем цифры с уникальными буквами
	digits[0] = counts['z'-'a'] // zero (z)
	digits[2] = counts['w'-'a'] // two (w)
	digits[4] = counts['u'-'a'] // four (u)
	digits[6] = counts['x'-'a'] // six (x)
	digits[8] = counts['g'-'a'] // eight (g)

	// Обрабатываем зависимые цифры
	digits[1] = counts['o'-'a'] - digits[0] - digits[2] - digits[4] // one (o)
	digits[3] = counts['h'-'a'] - digits[8]                         // three (h)
	digits[5] = counts['f'-'a'] - digits[4]                         // five (f)
	digits[7] = counts['s'-'a'] - digits[6]                         // seven (s)
	digits[9] = counts['i'-'a'] - digits[5] - digits[6] - digits[8] // nine (i)

	// Собираем результат
	var result strings.Builder
	for digit := 0; digit < 10; digit++ {
		for i := 0; i < digits[digit]; i++ {
			result.WriteByte(byte(digit) + '0')
		}
	}

	return result.String()
}
