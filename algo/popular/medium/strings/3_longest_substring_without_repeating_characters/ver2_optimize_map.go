package __longest_substring_without_repeating_characters

// Функция для поиска самой длинной подстроки без повторяющихся символов
func lengthOfLongestSubstringV2(s string) int {
	/*
		METHOD: Optimize map
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY:  O(min(n, m)), где m - размер алфавита, а n - количество символов в окне.
	*/
	// Инициализируем map для хранения последнего встречаемого символа
	m := make(map[rune]int)
	// Инициализируем указатели на начало и конец подстроки
	start, maxLength := 0, 0
	for i, char := range s {
		// Если символ уже встречался и его позиция больше начала текущей подстроки
		if _, ok := m[char]; ok && m[char] >= start {
			// Обновляем начало подстроки
			start = m[char] + 1
		}
		// Обновляем позицию символа
		m[char] = i
		// Обновляем максимальную длину, если необходимо
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}
	return maxLength
}
