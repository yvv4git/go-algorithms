package main

func halvesAreAlikeV2(s string) bool {
	/*
		Method: Count
		Time complexity: O(n), где n - длина строки, потому что мы проходим по каждому символу строки ровно один раз.
		Space complexity: O(1), потому что мы используем фиксированное количество дополнительной памяти, независимо от размера входных данных.

		Метод, используемый в этом решении, заключается в подсчете количества гласных букв в первой половине строки
		и вычитании количества гласных букв во второй половине строки. Если количество гласных букв равно нулю,
		то строки одинаковы.
	*/
	// Создаем множество гласных букв.
	vowelsSet := make(map[byte]struct{})

	// Заполняем множество гласных букв.
	for _, b := range []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'} {
		vowelsSet[b] = struct{}{}
	}

	// Инициализируем счетчики гласных букв в левой и правой половинах строки.
	leftCount, rightCount := 0, 0

	// Вычисляем смещение для правой половины строки.
	offset := len(s) / 2

	// Проходим по левой половине строки.
	for i := 0; i < len(s)/2; i++ {
		// Если текущий символ является гласной буквой,
		// то увеличиваем счетчик гласных букв в левой половине.
		if _, ok := vowelsSet[s[i]]; ok {
			leftCount++
		}

		// Если текущий символ является гласной буквой,
		// то увеличиваем счетчик гласных букв в правой половине.
		if _, ok := vowelsSet[s[i+offset]]; ok {
			rightCount++
		}
	}

	// Если количество гласных букв в левой и правой половинах строки равно,
	// то возвращаем true, иначе - false.
	return leftCount == rightCount
}