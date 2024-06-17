package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	/*
		METHOD: Math

		TIME COMPLEXITY: O(k), где k — длина повторяющейся части десятичной дроби.
		В худшем случае, когда дробь имеет бесконечно повторяющуюся часть, временная сложность может быть O(n),
		где n — количество цифр в повторяющейся части.

		SPACE COMPLEXITY: O(k), где k — длина повторяющейся части десятичной дроби.
		Это связано с использованием карты для отслеживания остатков.
	*/
	// Если числитель равен 0, результатом будет "0"
	if numerator == 0 {
		return "0"
	}

	// Инициализация результата
	result := ""

	// Обработка знака результата
	if (numerator < 0) != (denominator < 0) {
		result += "-"
	}

	// Приведение числителя и знаменателя к положительным числам
	num := abs(numerator)
	den := abs(denominator)

	// Вычисление целой части
	integerPart := num / den
	result += strconv.Itoa(integerPart)

	// Вычисление остатка
	remainder := num % den
	if remainder == 0 {
		return result
	}

	// Добавление десятичной точки
	result += "."

	// Создание карты для отслеживания остатков и их позиций
	remainderMap := make(map[int]int)

	// Вычисление дробной части
	for remainder != 0 {
		// Если остаток уже встречался, значит, начинается повторение
		if pos, found := remainderMap[remainder]; found {
			result = result[:pos] + "(" + result[pos:] + ")"
			break
		}

		// Сохранение позиции остатка
		remainderMap[remainder] = len(result)

		// Вычисление следующей цифры дробной части
		remainder *= 10
		result += strconv.Itoa(remainder / den)
		remainder %= den
	}

	return result
}

// Вспомогательная функция для вычисления абсолютного значения
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(fractionToDecimal(1, 3)) // Вывод: "0.(3)"
	fmt.Println(fractionToDecimal(2, 7)) // Вывод: "0.(285714)"
	fmt.Println(fractionToDecimal(1, 6)) // Вывод: "0.1(6)"
	fmt.Println(fractionToDecimal(1, 2)) // Вывод: "0.5"
}
