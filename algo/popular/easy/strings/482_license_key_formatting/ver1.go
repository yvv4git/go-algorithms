package main

import (
	"strings"
	"unicode"
)

// Функция для форматирования ключа лицензии
func licenseKeyFormatting(s string, k int) string {
	/*
		METHOD: Iterative
		Используемый подход - это простой перебор символов в строке в обратном порядке,
		и добавление их в новую строку с учетом правил форматирования.

		TIME COMPLEXITY: O(n), где n - количество символов в строке.
		Это обусловлено тем, что мы проходим по каждому символу в строке ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем сохранить форматированную строку размером n символов
		(если все символы в исходной строке были буквенно-цифровыми).
	*/
	// Удаляем дефисы из строки
	s = strings.ReplaceAll(s, "-", "")

	// Переводим все символы в верхний регистре
	s = strings.ToUpper(s)

	// Проверяем, что строка не пустая
	if len(s) == 0 {
		return ""
	}

	// Инициализируем новую строку для форматированного ключа
	formatted := ""

	// Инициализируем счетчик для отслеживания количества символов в группе
	count := 0

	// Проходим по строке в обратном порядке
	for i := len(s) - 1; i >= 0; i-- {
		// Если текущий символ - буква или цифра
		if unicode.IsLetter(rune(s[i])) || unicode.IsDigit(rune(s[i])) {
			// Если количество символов в текущей группе достигло k,
			// добавляем дефис и сбрасываем счетчик
			if count == k {
				formatted = "-" + formatted
				count = 0
			}

			// Добавляем текущий символ в форматированную строку
			formatted = string(s[i]) + formatted
			count++
		}
	}

	// Возвращаем форматированный ключ лицензии
	return formatted
}

func main() {
	// Пример использования функции
	s := "2-5g-3-J"
	k := 4
	formatted := licenseKeyFormatting(s, k)
	println(formatted) // Вывод: "2-5G3J"
}