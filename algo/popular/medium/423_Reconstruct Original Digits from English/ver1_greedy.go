package main

import (
	"fmt"
	"strings"
)

func main() {
	// Тестовые случаи
	testCases := []struct {
		input    string
		expected string
	}{
		{"owoztneoer", "012"},   // "zero", "one", "two"
		{"fviefuro", "45"},      // "four", "five"
		{"xsi", "6"},            // "six"
		{"nnei", "9"},           // "nine"
		{"zerozero", "00"},      // два "zero"
		{"twotwo", "22"},        // два "two"
		{"esnvei", "79"},        // "seven", "nine"
		{"", ""},                // пустая строка
		{"oneoneone", "111"},    // три "one"
		{"eightwothree", "238"}, // "two", "three", "eight"
	}

	// Проверка каждого тестового случая
	for _, tc := range testCases {
		result := originalDigits(tc.input)
		fmt.Printf("Ввод: %-15s Ожидаемый результат: %-8s Полученный результат: %-8s Правильно: %t\n",
			"\""+tc.input+"\"",
			"\""+tc.expected+"\"",
			"\""+result+"\"",
			result == tc.expected)
	}
}

func originalDigits(s string) string {
	/*
		APPROACH: Greedy Counting
		- Подсчитываем количество каждой буквы в строке.
		- Используем уникальные буквы для каждой цифры.
		- Обрабатываем цифры с зависимостями (например, три встречается в восьми и трех).
		- Формируем строку с цифрами.
		- Возвращаем строку с цифрами.

		TIME COMPLEXITY: O(N)
		- Один проход по строке.

		SPACE COMPLEXITY: O(1)
		- Массивы фиксированного размера для подсчета букв.
		- O(26+10) -> O(1).
	*/
	counts := make([]int, 26)
	for _, c := range s {
		counts[c-'a']++
	}

	digits := make([]int, 10)

	// Уникальные буквы для цифр
	digits[0] = counts['z'-'a'] // zero (z)
	digits[2] = counts['w'-'a'] // two (w)
	digits[4] = counts['u'-'a'] // four (u)
	digits[6] = counts['x'-'a'] // six (x)
	digits[8] = counts['g'-'a'] // eight (g)

	// Обработка цифр с зависимостями
	digits[3] = counts['h'-'a'] - digits[8]                         // three (h встречается в eight и three)
	digits[5] = counts['f'-'a'] - digits[4]                         // five (f встречается в four и five)
	digits[7] = counts['s'-'a'] - digits[6]                         // seven (s встречается в six и seven)
	digits[1] = counts['o'-'a'] - digits[0] - digits[2] - digits[4] // one (o встречается в zero, two, four)
	digits[9] = counts['i'-'a'] - digits[5] - digits[6] - digits[8] // nine (i встречается в five, six, eight)

	var result strings.Builder
	for i := 0; i < 10; i++ {
		for j := 0; j < digits[i]; j++ {
			result.WriteByte(byte(i) + '0')
		}
	}

	return result.String()
}
