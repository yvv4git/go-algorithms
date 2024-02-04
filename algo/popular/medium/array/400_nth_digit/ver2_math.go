package _00_nth_digit

import "math"

func findNthDigitV2(n int) int {
	/*
		METHOD: Math
		TIME COMPLEXITY: O(log n)
		SPACE COMPLEXITY: O(1)
	*/

	//1-9 		-> 9 * length-1 number
	//10-99 	-> 90 * length-2 number
	//100-999 	-> 900 * length-3 number
	// Если n меньше или равно 9, то n-ая цифра равна n.
	if n <= 9 {
		return n
	}

	// Инициализация переменных
	nine := 9      // Количество цифр в числах с i цифрами
	curDigits := 0 // Общее количество цифр в числах с меньше i цифрами
	curNum := 0    // Наибольшее число с i-1 цифрами
	i := 1         // Количество цифр в текущем числе

	// Цикл по числам с i цифрами
	for ; i <= 10; i++ {
		// Обновление общего количества цифр и наибольшего числа с i-1 цифрами
		curDigits += nine * i
		curNum = curNum*10 + 9
		nine *= 10

		// Если общее количество цифр плюс количество цифр в следующем числе больше n, то выходим из цикла
		if curDigits+(i+1)*nine > n {
			i++
			break
		}
	}

	// Вычисление разницы в значении между n-ой цифрой и наибольшим числом с i-1 цифрами
	dif := (n - curDigits) / i
	// Вычисление позиции n-ой цифры в числе
	digit := (n - curDigits) % i
	// Вычисление числа, которое содержит n-ую цифру
	result := curNum + dif

	// Если позиция n-ой цифры не равна 0, то переходим к следующему числу
	if digit != 0 {
		result += 1
	} else {
		// Если позиция n-ой цифры равна 0, то возвращаем последнюю цифру числа
		return result % 10
	}

	// Возвращаем n-ую цифру в числе
	return result / int(math.Pow(float64(10), float64(i-digit))) % 10
}
