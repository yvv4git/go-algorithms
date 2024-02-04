package main

import "strings"

func reverseWordsV2(s string) string {
	/*
		METHOD: Iterative
		Time complexity: O(n), где n - количество слов в строке. Это связано с тем, что мы проходим по каждому слову в строке один раз.
		Space complexity: O(n), так как мы создаем новый массив для хранения слов в строке. В худшем случае это будет n, если каждое слово в строке односимвольное.
	*/

	// Добавляем пробел в начало строки, чтобы учесть случай, когда строка начинается с пробела
	s = " " + s

	// Используем strings.Builder для эффективного создания строки
	var res strings.Builder

	// Инициализируем правый указатель на последний символ строки
	right := len(s) - 1

	// Идем справа налево по строке
	for left := len(s) - 2; left >= 0; left-- {
		// Если текущий символ - пробел, а следующий - не пробел, значит, начинается новое слово
		if s[left] == ' ' && s[left+1] != ' ' {
			// Если уже есть слова в результате, добавляем пробел перед новым словом
			if res.Len() > 0 {
				res.WriteString(" ")
			}

			// Добавляем новое слово в результат
			res.WriteString(s[left+1 : right+1])
		}

		// Если текущий символ - не пробел, а следующий - пробел, значит, текущее слово закончилось
		if s[left] != ' ' && s[left+1] == ' ' {
			// Обновляем правый указатель на начало текущего слова
			right = left
		}
	}

	// Возвращаем перевернутую строку
	return res.String()
}
