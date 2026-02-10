package main

// maxRepeatingKMP находит максимальное k для word^k в sequence с помощью KMP.
func maxRepeatingKMP(sequence string, word string) int {
	/*
		INTUITION:
		Алгоритм КМП (Knuth-Morris-Pratt) использует префикс-функцию для
		эффективного поиска подстроки в строке. Префикс-функция π[i] показывает
		длину наибольшего собственного префикса, который также является суффиксом
		для подстроки s[0:i].

		Идея: если word полностью совпадает с subsequence[i-len(word):i],
		то мы можем расширить предыдущую последовательность, используя
		префикс-функцию на строке word + "#" + sequence.

		APPROACH: KMP (Knuth-Morris-Pratt)
		1. Строим префикс-функцию для строки: pattern + "#" + text
		2. Если π[i] >= len(word), значит word заканчивается на позиции i
		3. Используем счетчик для подсчета последовательных повторений

		БАЗА: π[0] = 0 для пустой строки

		Переход:
		- Если символы совпадают: j = π[j-1]
		- Иначе: j уменьшается пока не найдем совпадение или j = 0

		ПОЧЕМУ ТАК:
		- KMP работает за O(n+m) вместо O(n*m) как в наивном поиске
		- Префикс-функция позволяет "помнить" частичные совпадения
		- Эффективен для строк с множеством повторений

		TIME COMPLEXITY: O(n + m) где n = len(sequence), m = len(word)
		SPACE COMPLEXITY: O(n + m) для хранения префикс-функции
	*/

	m := len(word)

	// Строим префикс-функцию для строки word + "#" + sequence.
	combined := word + "#" + sequence
	pi := make([]int, len(combined))

	maxK := 0
	currentK := 0

	for i := 1; i < len(combined); i++ {
		j := pi[i-1]

		// Уменьшаем j пока не найдем совпадение или j = 0.
		for j > 0 && combined[i] != combined[j] {
			j = pi[j-1]
		}

		// Если символы совпадают, увеличиваем j.
		if combined[i] == combined[j] {
			j++
		}

		pi[i] = j

		// Если j >= m, значит word полностью совпал.
		if pi[i] >= m {
			// Это конец word в позиции i.
			currentK++

			// Обновляем максимум.
			if currentK > maxK {
				maxK = currentK
			}
		} else {
			// Сбрасываем счетчик, если текущее совпадение меньше word.
			currentK = 0
		}
	}

	return maxK
}
