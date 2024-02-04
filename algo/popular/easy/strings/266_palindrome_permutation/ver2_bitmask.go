package _66_palindrome_permutation

func canPermutePalindromeV2(s string) bool {
	/*
		METHOD: Bitmask, Bitwise
		Time complexity: O(n), где n - длина строки, поскольку мы проходим по всем символам строки.
		Space complexity: O(1), поскольку мы используем фиксированное количество переменных, независимо от размера входных данных.

		В этом коде мы используем побитовую операцию XOR (^) для каждого символа в строке. Если результат XOR равен 0,
		то строка является перестановкой палиндрома. Если результат XOR равен 1, то строка является перестановкой палиндрома,
		если и только если строка имеет четной длину.
	*/
	// Инициализируем переменную bitVector типа int64
	var bitVector int64

	// Проходим по всем символам строки
	for _, char := range s {
		// Для каждого символа выполняем побитовую операцию XOR с bitVector
		// Индекс символа в bitVector вычисляется путем вычитания кода символа 'a'
		bitVector ^= 1 << (char - 'a')
	}

	// Если bitVector равен 0, то строка является перестановкой палиндрома
	if bitVector == 0 {
		return true
	}

	// Если bitVector не равен 0, то проверяем, является ли строка перестановкой палиндрома,
	// если и только если строка имеет четной длину
	return (bitVector & (bitVector - 1)) == 0
}
