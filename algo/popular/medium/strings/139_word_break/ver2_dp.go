package main

// Функция wordBreakV2 принимает строку s и словарь wordDict, а затем проверяет,
// можно ли разбить строку на слова из словаря.
func wordBreakV2(s string, wordDict []string) bool {
	/*
		Method: Dynamic programming & Memoization
		Time complexity: O(n^2), где n - длина строки s. Это связано с двумя вложенными циклами, которые проходят по всем подстрокам строки s.
		Space complexity: O(n), так как мы используем дополнительный массив dp для хранения результатов подзадач.
	*/
	// Создаем множество wordSet для быстрого поиска слов в словаре.
	wordSet := make(map[string]struct{}, len(wordDict))
	for _, str := range wordDict {
		wordSet[str] = struct{}{}
	}

	// Создаем словарь memory для хранения уже вычисленных результатов.
	memory := map[string]bool{}

	// Запускаем цикл, который начинает с конца строки s и движется к началу.
	for offset := len(s) - 1; offset > 0; offset-- {
		// Для каждой подстроки subStr вызываем функцию workHorse,
		// которая проверяет, можно ли разбить подстроку на слова из словаря.
		subStr := s[offset:]
		memory[subStr] = workHorse(subStr, wordSet, memory)
	}

	// В конце вызываем функцию workHorse для исходной строки s,
	// чтобы определить, можно ли разбить ее на слова из словаря.
	return workHorse(s, wordSet, memory)
}

// Функция workHorse принимает строку s, множество wordSet со словами из словаря и словарь solutions для хранения уже вычисленных результатов.
// Она проверяет, можно ли разбить строку s на слова из словаря.
func workHorse(s string, wordSet map[string]struct{}, solutions map[string]bool) bool {
	// Запускаем цикл, который проходит по всем префиксам строки s.
	for i := 1; i <= len(s); i++ {
		// Извлекаем префикс s[:i] как возможное слово.
		word := s[:i]

		// Проверяем, присутствует ли префикс в словаре wordSet.
		if _, inWords := wordSet[word]; inWords {
			// Если префикс есть в словаре, проверяем, является ли оставшаяся часть строки s[i:] пустой строкой или уже разбиваемой на слова.
			// Если это так, возвращаем true, так как строка s можно разбить на слова из словаря.
			if i == len(s) || solutions[s[i:]] {
				return true
			}
		}
	}

	// Если мы прошли весь цикл и не нашли разбиения строки на слова из словаря, возвращаем false.
	return false
}
