package main

func makeGoodV3(s string) string {
	/*
		METHOD: Two pointer / Window.
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	result := []rune(s) // Преобразуем строку в срез рун для удобства работы
	writePointer := 0   // Указатель для записи "хороших" символов

	for readPointer := 0; readPointer < len(result); readPointer++ {
		// Если writePointer > 0 и текущий символ образует "плохую" пару с последним записанным символом
		if writePointer > 0 && abs(int(result[writePointer-1])-int(result[readPointer])) == 32 {
			writePointer-- // Откатываем указатель записи на один шаг назад
		} else {
			// Записываем текущий символ в результат
			result[writePointer] = result[readPointer]
			writePointer++
		}
	}

	// Возвращаем строку, состоящую из "хороших" символов
	return string(result[:writePointer])
}
