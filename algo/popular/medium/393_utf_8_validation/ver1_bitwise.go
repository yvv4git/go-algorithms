package main

func validUtf8(data []int) bool {
	/*
		METHOD: Bitwise
		Используется побитовый анализ для определения типа символа UTF-8 и проверки корректности последовательности байтов.
		Проходим по каждому байту и определяем, является ли он началом нового символа или продолжающим байтом.

		TIME COMPLEXITY:
		O(n), где n — количество байтов в массиве data. Мы проходим по каждому байту ровно один раз.

		SPACE COMPLEXITY:
		O(1), так как используется только константное количество дополнительной памяти (переменная remainingBytes).
	*/
	// Количество оставшихся байтов для текущего символа UTF-8
	remainingBytes := 0

	// Проходим по каждому байту в массиве
	for _, byteValue := range data {
		if remainingBytes == 0 {
			// Определяем количество байтов для текущего символа UTF-8
			if byteValue&0b10000000 == 0 {
				// Одно-байтовый символ
				remainingBytes = 0
			} else if byteValue&0b11100000 == 0b11000000 {
				// Двух-байтовый символ
				remainingBytes = 1
			} else if byteValue&0b11110000 == 0b11100000 {
				// Трех-байтовый символ
				remainingBytes = 2
			} else if byteValue&0b11111000 == 0b11110000 {
				// Четырех-байтовый символ
				remainingBytes = 3
			} else {
				// Некорректный первый байт
				return false
			}
		} else {
			// Проверяем, что текущий байт является продолжающим байтом
			if byteValue&0b11000000 != 0b10000000 {
				return false
			}
			// Уменьшаем количество оставшихся байтов
			remainingBytes--
		}
	}

	// Если remainingBytes == 0, значит все символы корректны
	return remainingBytes == 0
}
