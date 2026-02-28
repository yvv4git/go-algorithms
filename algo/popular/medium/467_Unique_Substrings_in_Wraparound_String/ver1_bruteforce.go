package main

// Задача: найти количество уникальных подстрок строки s, присутствующих в бесконечной строке base.
// base = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz..." (бесконечный цикл a-z)
//
// Пример:
//   s = "zab"
//   Подстроки s: "z", "a", "b", "za", "ab", "zab"
//   Валидные подстроки в base: все 6 подстрок (z->a это цикл, a->b идёт по алфавиту)
//   Ответ: 6
//
//   s = "cac"
//   Подстроки s: "c", "a", "c", "ca", "ac", "cac"
//   Валидные: "a", "c" (ca не подряд, ac не подряд, cac не подряд)
//   Ответ: 2

// findSubstringInWraproundString возвращает количество уникальных подстрок s,
// которые присутствуют в бесконечной строке base
func findSubstringInWraproundString(s string) int {
	/*
		METHOD: Brute Force
		APPROACH: Генерируем все подстроки и проверяем каждую на валидность
		TIME COMPLEXITY: O(n^2)
		SPACE COMPLEXITY: O(n^2)

		Time complexity
		Временная сложность - O(n^2), где n - длина строки s.
		Мы генерируем все подстроки (n*(n+1)/2 штук) и для каждой проверяем валидность за O(n) в худшем случае.

		Space complexity
		Пространственная сложность - O(n^2), так как в map может храниться до n*(n+1)/2 уникальных подстрок.
		В худшем случае (например, строка "abcdefghijklmnopqrstuvwxyz") каждая подстрока уникальна.
	*/
	n := len(s)
	if n == 0 {
		return 0
	}

	// Используем map для хранения уникальных подстрок
	uniqueSubstrings := make(map[string]bool)

	// Генерируем все подстроки
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			substr := s[i:j]
			// Проверяем, является ли подстрока валидной для base
			if isValidInBase(substr) {
				uniqueSubstrings[substr] = true
			}
		}
	}

	return len(uniqueSubstrings)
}

// isValidInBase проверяет, является ли подстрока подстрокой бесконечной строки base
// Подстрока валидна, если символы идут подряд по алфавиту (с учётом цикла)
//
// Примеры:
//
//	"abc" -> true  (a->b->c идут подряд)
//	"za"  -> true  (после z идёт a - цикл)
//	"abz" -> false (b->z не подряд)
//	"ace" -> false (c->e не подряд)
func isValidInBase(substr string) bool {
	if len(substr) <= 1 {
		return true
	}

	for i := 1; i < len(substr); i++ {
		prev := substr[i-1]
		curr := substr[i]

		// Проверяем, идут ли символы подряд по алфавиту
		// После z должен идти a (цикл)
		if !((prev == 'z' && curr == 'a') || (curr == prev+1)) {
			return false
		}
	}

	return true
}
