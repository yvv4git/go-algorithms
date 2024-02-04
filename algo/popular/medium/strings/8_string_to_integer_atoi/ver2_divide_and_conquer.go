package main

// Функция myAtoiV2 преобразует строку в число.
func myAtoiV2(s string) int {
	/*
		METHOD: Divide and conquer
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)

		Разделяй и властвуй:
		1. Разделяй:
		Данный метод разделяет исходную задачу на несколько меньших подзадач.
		В данном случае, строка преобразуется в число, которое является основной задачей.

		2. Властвуй:
		Каждая из полученных подзадач решается отдельно. В данном случае, каждая цифра в строке преобразуется в число.

		3. Объединяй:
		Решения подзадач объединяются, чтобы получить решение исходной задачи. В данном случае, цифры объединяются, чтобы получить число.
	*/

	// Переменные для хранения промежуточных результатов
	var c string                // строка для хранения цифр
	var a, d, cnt int = 0, 1, 0 // a - количество символов '-' или '+', d - знак числа, cnt - количество обработанных символов

	// Цикл по строке
	for _, i := range s {
		// Если символ - пробел, то проверяем, были ли обработаны символы до этого
		if i == ' ' {
			if cnt > 0 {
				break
			}
			continue
		}
		cnt++
		// Если символ - '-' или '+', то проверяем, были ли обработаны символы до этого и количество символов '-' или '+'
		if i == '-' || i == '+' {
			if len(c) > 0 {
				break
			}
			a++
			if a > 1 {
				break
			}
			if i == '-' {
				d = -1
			}
		} else if i >= '0' && i <= '9' {
			// Если символ - цифра, то добавляем ее в строку c
			c += string(i)
		} else {
			// Если символ - не цифра и не '-' и не '+', то выходим из цикла
			break
		}
	}

	n := 0
	// Цикл по строке c
	for i := 0; i < len(c); i++ {
		// Умножаем текущее число на 10 и добавляем следующую цифру
		n *= 10
		n += int(c[i]) - 48
		// Если текущее число выходит за пределы int32, то возвращаем ближайшее к этому предельное значение
		if n*d > ((1 << 31) - 1) {
			return ((1 << 31) - 1)
		}
		if n*d < -(1 << 31) {
			return -(1 << 31)
		}
	}

	// Возвращаем результат, умноженный на знак числа
	return n * d
}
