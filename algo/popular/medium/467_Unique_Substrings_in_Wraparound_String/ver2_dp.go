package main

// DP подход: для каждого символа вычисляем длину максимальной непрерывной
// последовательности, заканчивающейся на этот символ.
//
// Пример:
//   s = "zab"
//   d[z] = 1, d[a] = 2 (z->a цикл), d[b] = 3 (a->b подряд)
//   max[z] = 1, max[a] = 2, max[b] = 3
//   Ответ = 1 + 2 + 3 = 6
//
//   s = "cac"
//   d[c] = 1, d[a] = 1 (c->a разрыв), d[c] = 1 (a->c разрыв)
//   max[c] = 1, max[a] = 1
//   Ответ = 1 + 1 = 2
//
// Сложность: O(n) по времени, O(1) по памяти (массив из 26 элементов)

func findSubstringInWraproundStringDP(s string) int {
	/*
		METHOD: Dynamic Programming
		APPROACH: Вычисляем длину непрерывной последовательности для каждого символа и запоминаем максимум для каждой буквы
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)

		Time complexity
		Временная сложность - O(n), где n - длина строки s.
		Мы проходим по строке один раз, вычисляя длину непрерывной последовательности для каждого символа.

		Space complexity
		Пространственная сложность - O(1), так как используем массив из 26 элементов (для каждой буквы алфавита).
		Независимо от длины входной строки, память всегда фиксирована.
	*/
	n := len(s)
	if n == 0 {
		return 0
	}

	// maxLen[char] - максимальная длина непрерывной последовательности,
	// заканчивающейся на символ char
	maxLen := [26]int{}

	// Текущая длина непрерывной последовательности
	currentLen := 0

	for i := 0; i < n; i++ {
		curr := s[i]

		// Если текущий символ продолжает последовательность
		if i > 0 && isContinuous(s[i-1], curr) {
			currentLen++
		} else {
			currentLen = 1
		}

		// Обновляем максимальную длину для этого символа
		charIndex := curr - 'a'
		if currentLen > maxLen[charIndex] {
			maxLen[charIndex] = currentLen
		}
	}

	// Суммируем все максимальные длины
	result := 0
	for i := 0; i < 26; i++ {
		result += maxLen[i]
	}

	return result
}

// isContinuous проверяет, идут ли два символа подряд в алфавите (с учётом цикла)
func isContinuous(prev, curr byte) bool {
	// После z идёт a (цикл)
	if prev == 'z' && curr == 'a' {
		return true
	}
	// Обычный случай: curr = prev + 1
	return curr == prev+1
}
