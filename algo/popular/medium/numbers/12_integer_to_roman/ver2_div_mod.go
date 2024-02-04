package main

func intToRomanV2(num int) string {
	/*
		METHOD: Деления и взятия остатка
		TIME COMPLEXITY: O(log(n)), где n - исходное десятичное число. Это связано с тем, что алгоритм проходит по каждой цифре числа, количество цифр которого равно log(n).
		Space complexity:  O(1), так как алгоритм использует фиксированное количество переменных, не зависящих от размера входных данных.
		Это связано с тем, что алгоритм не использует дополнительную память, которая была бы зависеть от размера входных данных.
	*/
	// Инициализация переменной для хранения результата
	result := ""

	// Переменная k используется для индексации в массиве s
	k := 0

	// Массив s содержит римские цифры в порядке убывания значений
	s := []string{"I", "V", "X", "L", "C", "D", "M"}

	// Цикл продолжается, пока num больше 0
	for num > 0 {
		// Вычисляем последнюю цифру числа
		d := num % 10

		// Переменная v используется для хранения римского числа для текущей цифры
		v := ""

		// Проверяем значение последней цифры
		if d < 4 {
			// Если цифра меньше 4, то добавляем соответствующее количество римских цифр в v
			for j := 0; j < d; j++ {
				v += s[k]
			}
		} else if d == 4 {
			// Если цифра равна 4, то добавляем римскую цифру и следующую большую
			v += s[k] + s[k+1]
		} else if d < 9 {
			// Если цифра меньше 9, то добавляем следующую большую римскую цифру и дополнительные меньшие
			v += s[k+1]
			for j := 0; j < d-5; j++ {
				v += s[k]
			}
		} else if d == 9 {
			// Если цифра равна 9, то добавляем римскую цифру и следующую меньшую
			v += s[k] + s[k+2]
		}

		// Добавляем полученное римское число в начало результата
		result = v + result

		// Увеличиваем k на 2, так как мы обрабатываем две римские цифры за раз
		k += 2

		// Удаляем последнюю цифру из числа
		num /= 10
	}

	return result
}
