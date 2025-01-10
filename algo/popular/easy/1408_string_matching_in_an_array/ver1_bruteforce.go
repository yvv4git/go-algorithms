package main

import "strings"

func stringMatching(words []string) []string {
	/*
		METHOD: Brute Force

		TIME COMPLEXITY: O(n^2 * m), где n - количество слов в массиве, m - средняя длина слов.
		SPACE COMPLEXITY: O(n), так как используется массив `result` для хранения результата.
	*/

	// Результирующий массив для хранения подстрок
	result := make([]string, 0)

	// Проходим по всем словам
	for i := 0; i < len(words); i++ {
		// Для каждого слова проверяем, является ли оно подстрокой других слов
		for j := 0; j < len(words); j++ {
			// Пропускаем сравнение слова с самим собой
			if i == j {
				continue
			}

			// Если текущее слово является подстрокой другого слова
			if strings.Contains(words[j], words[i]) {
				// Добавляем его в результат и прерываем внутренний цикл
				result = append(result, words[i])
				break
			}
		}
	}

	return result
}
