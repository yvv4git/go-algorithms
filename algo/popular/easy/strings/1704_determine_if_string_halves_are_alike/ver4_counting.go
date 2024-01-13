package main

func halvesAreAlikeV4(s string) bool {
	/*
		Method: Count
		Time complexity: O(n), где n - длина строки, потому что мы проходим по каждому символу строки ровно один раз.
		Space complexity: O(1), потому что мы используем фиксированное количество дополнительной памяти, независимо от размера входных данных.

		Метод, используемый в этом решении, заключается в подсчете количества гласных букв в первой половине строки
		и вычитании количества гласных букв во второй половине строки. Если количество гласных букв равно нулю,
		то строки одинаковы.
	*/

	// Вычисляем середину строки.
	mid := len(s) / 2

	// Инициализируем счетчик гласных букв.
	count := 0

	// Проходим по всем символам строки.
	for i := 0; i < len(s); i++ {
		// Проверяем, является ли символ гласной буквой.
		switch rune(s[i]) {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			// Если символ находится в первой половине строки,
			// то увеличиваем счетчик.
			if i < mid {
				count++
			} else {
				// Если символ находится во второй половине строки,
				// то уменьшаем счетчик.
				count--
			}
		}
	}

	// Если счетчик равен нулю, то возвращаем true,
	// иначе - false.
	return count == 0
}