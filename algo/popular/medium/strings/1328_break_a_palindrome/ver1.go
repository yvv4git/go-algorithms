package _328_break_a_palindrome

func breakPalindromeV1(palindrome string) string {
	/*
		METHOD: Iterative
		TIME COMPLEXITY: O(n), где n - длина строки. Это связано с тем, что мы проходим по строке только один раз.
		SPACE COMPLEXITY: O(1), так как мы не используем дополнительное пространство, кроме входных данных.

		Метод решения заключается в том, чтобы заменить первый символ строки, который не является 'a', на 'a'.
		Если такого символа нет, то мы заменяем последний символ на 'b'.
		Это гарантирует, что полученная строка будет лексикографически меньше исходной, а также не будет являться палиндромом.

		Если исходная строка состоит из одного символа, то мы не можем изменить ее, поэтому возвращаем пустую строку.
	*/
	// Получаем длину строки
	n := len(palindrome)

	// Если длина строки равна 1, то возвращаем пустую строку, так как невозможно изменить символ в односимвольной строке
	if n == 1 {
		return ""
	}

	// Проходим по первой половине строки
	for i := 0; i < n/2; i++ {
		// Если символ не равен 'a', то меняем его на 'a' и возвращаем измененную строку
		if palindrome[i] != 'a' {
			return palindrome[:i] + "a" + palindrome[i+1:]
		}
	}

	// Если все символы в первой половине строки равны 'a', то меняем последний символ на 'b' и возвращаем измененную строку
	return palindrome[:n-1] + "b"
}
