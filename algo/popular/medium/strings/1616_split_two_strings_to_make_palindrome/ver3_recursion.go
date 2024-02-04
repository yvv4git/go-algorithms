package main

// Функция для проверки, можно ли сделать из двух строк палиндром
func checkPalindromeFormationV3(a string, b string) bool {
	/*
		METHOD: Recursion
		Time complexity: O(n)
		Space complexity: O(1) or O(n) - это связано с тем, что функция recursiveCheckPalindromeFormation использует стек вызовов, который может достигать глубину n в худшем случае.
	*/
	// Получаем длину строки a
	var l int = len(a)
	// Если длина строки меньше или равна 1, то она уже является палиндромом
	if l <= 1 {
		return true
	}
	// Рекурсивно проверяем, можно ли сделать из строк a и b палиндром
	// Или, проверяем, можно ли сделать из строк b и a палиндром
	return recursiveCheckPalindromeFormation(a, l, 0, b, l, l-1) ||
		recursiveCheckPalindromeFormation(b, l, 0, a, l, l-1)
}

// Рекурсивная функция для проверки, можно ли сделать из двух строк палиндром
func recursiveCheckPalindromeFormation(s1 string, l1 int, left int, s2 string, l2 int, right int) bool {
	// Если левая граница больше или равна правой, то строки уже являются палиндромами
	if left != 0 && left >= right {
		return true
	}
	// Если символы в строках s1 и s2 равны, то проверяем, можно ли сделать из оставшихся строк палиндром
	if s1[left] == s2[right] {
		return recursiveCheckPalindromeFormation(s1, l1, left+1, s2, l2, right-1) ||
			recursiveCheckPalindromeFormation(s1, l1, left+1, s1, l1, right-1) ||
			recursiveCheckPalindromeFormation(s2, l2, left+1, s2, l1, right-1)
	}
	// Если символы не равны, то строки не могут быть преобразованы в палиндромы
	return false
}
