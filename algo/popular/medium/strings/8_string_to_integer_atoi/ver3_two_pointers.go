package main

import "unicode"

// Функция myAtoiV3 преобразует строку в число.
func myAtoiV3(s string) int {
	/*
		METHOD: Two pointers
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1) - в общем случае, в худшем случае, если строка состоит только из цифр, то n записей придется копировать.


		В данном случае, использование двух указателей (two pointers) выбранное решение, поскольку оно эффективно и эффективно.
		Двумя указателями здесь реализуется функция, которая проходит по строке слева направо и проверяет каждый символ.
		Если символ является цифрой, то он добавляется к результату. Если символ не является цифрой,
		то происходит проверка на наличие знака минус или плюс в начале строки.
		Если знак присутствует, то он также добавляется к результату.
	*/
	// Инициализируем переменные sign, res и started.
	// sign - это знак числа, res - это результат преобразования, started - это флаг, показывающий, что преобразование началось.
	sign := 1
	res := 0
	started := false

	// Проходимся по строке.
	for _, c := range s {
		// Если символ - пробел, то пропускаем его.
		// Если преобразование уже началось, то прекращаем его.
		if unicode.IsSpace(c) {
			if started {
				break
			} else {
				continue
			}
		}

		// Если символ - '-', то меняем знак числа на отрицательный.
		// Если преобразование уже началось, то прекращаем его.
		if c == '-' {
			if started {
				break
			} else {
				sign = -1
				started = true
				continue
			}
		}

		// Если символ - '+', то пропускаем его.
		// Если преобразование уже началось, то прекращаем его.
		if c == '+' {
			if started {
				break
			} else {
				started = true
				continue
			}
		}

		// Если символ - цифра, то добавляем ее в результат.
		// Если результат выходит за пределы int32, то возвращаем ближайшее к этому число, которое вмещается в int32.
		if unicode.IsDigit(c) {
			started = true
			res = res*10 + int(c-'0')
			if sign*res < -(1 << 31) {
				return -(1 << 31)
			} else if sign*res > (1<<31)-1 {
				return (1 << 31) - 1
			}
		} else {
			break
		}
	}

	// Возвращаем результат преобразования.
	return sign * res
}
