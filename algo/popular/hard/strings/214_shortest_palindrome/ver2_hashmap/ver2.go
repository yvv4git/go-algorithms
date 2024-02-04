package ver2_hashmap

import "fmt"

func shortestPalindrome(s string) string {
	/*
		METHOD: HashMap
		TIME COMPLEXITY: O(n^2), где n - длина входной строки.
		SPACE COMPLEXITY: O(n), где n - длина входной строки.
	*/
	// Если длина слова меньше 2, то самое короткое палиндромное слово - это само слово.
	if len(s) < 2 {
		return s
	}

	// Переменная r отвечает за индекс, начиная с которого следует добавлять символы в начало слова.
	r := 0
	// Переменная longestPalindrome хранит самое длинное найденное палиндромное слово.
	longestPalindrome := fmt.Sprintf("%c", s[0])

	// Функция isPalindromic проверяет, является ли слово палиндромом.
	var isPalindromic func(string) bool
	isPalindromic = func(str string) bool { // Функция будет выполняться за O(n)
		// Если длина слова меньше 2, то оно является палиндромом.
		if len(s) < 2 {
			return true
		}

		// Проверяем, является ли слово палиндромом.
		l, r := 0, len(str)-1
		for l < r {
			if str[l] != str[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	// Создаем карту, где ключ - символ, а значение - список индексов, на которых этот символ встречается в слове.
	m := make(map[byte][]int)

	// Проходим по всем символам слова.
	for i := 0; i < len(s); i++ { // O(n)
		// Добавляем индекс текущего символа в карту.
		m[s[i]] = append(m[s[i]], i)

		// Если символ встречается более одного раза, проверяем все его вхождения.
		if len(m[s[i]]) > 1 {
			for _, val := range m[s[i]] {
				// Проверяем, является ли слог между текущим символом и предыдущим символом палиндромом.
				if val == 0 {
					if i != val && isPalindromic(s[val:i+1]) { // Функция isPalindromic будет выполняться за O(n), но она вызывается внутри цикла, что дает O(n^2)
						// Если слог является палиндромом и его длина больше самого длинного найденного палиндрома, обновляем самое длинное палиндромное слово и индекс r.
						if i-val+1 > len(longestPalindrome) {
							longestPalindrome = s[val : i+1]
							r = i + 1
						}
					}
				}
			}
		}
	}

	// Если r равно 0, обновляем его на 1.
	if r == 0 {
		r = 1
	}

	// Добавляем символы в начало слова, чтобы получить самое короткое палиндромное слово.
	for i := r; i < len(s); i++ {
		longestPalindrome = fmt.Sprintf("%c", s[i]) + longestPalindrome + fmt.Sprintf("%c", s[i])
	}

	// Возвращаем самое короткое палиндромное слово.
	return longestPalindrome
}
