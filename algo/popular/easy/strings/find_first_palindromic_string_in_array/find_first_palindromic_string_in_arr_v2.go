package find_first_palindromic_string_in_array

func firstPalindromeV2(words []string) string {
	const emptyValue = ""

	var ok bool

	// Идем по всем словам - O(n)
	for _, w := range words {
		ok = true

		// Проверяем конкретную строку на палиндром - O(n)
		for i := 0; i < len(w)/2; i++ {
			if w[i] != w[len(w)-i-1] {
				ok = false
				break
			}
		}

		// Если строка оказалась палиндромом возвращаем ее.
		if ok {
			return w
		}
	}

	// Если ни одна стока не оказалась палиндромом - вернем пустую строку.
	return emptyValue
}
