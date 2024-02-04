package main

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	/*
		METHOD: Math
		TIME COMPLEXITY: O(n*m), где n и m - длины входных строк num1 и num2 соответственно.
		Space complexity: O(n+m), где n и m - длины входных строк num1 и num2 соответственно.

		Time complexity
		Временная сложность функции multiply составляет O(n*m), где n и m - длины входных строк num1 и num2 соответственно.
		Это обусловлено двумя вложенными циклами, которые проходят по каждому символу в обеих строках.

		Space complexity
		Пространственная сложность составляет O(n+m), где n и m - длины входных строк num1 и num2 соответственно.
		Это обусловлено хранением промежуточных результатов в срезе res, который имеет размер, равный сумме длин входных строк.

		Используемый подход - это ручной метод умножения столбиком, который основан на том,
		что произведение двух чисел можно представить в виде суммы произведений каждого разряда первого числа на каждый разряд второго числа,
		умноженных на соответствующий множитель (1, 10, 100 и т.д.).


	*/
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// Создаем срезы для хранения разрядов чисел и результата
	res := make([]int, len(num1)+len(num2))

	// Умножаем числа столбиком
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			// Вычисляем текущий разряд и перенос
			product := int(num1[i]-'0') * int(num2[j]-'0')
			sum := res[i+j+1] + product
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}

	// Формируем результат в виде строки
	var result strings.Builder
	for i := range res {
		if i == 0 && res[i] == 0 {
			continue // Пропускаем лидирующий ноль
		}
		result.WriteString(strconv.Itoa(res[i]))
	}

	return result.String()
}
