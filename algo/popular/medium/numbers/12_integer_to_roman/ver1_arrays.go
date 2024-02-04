package main

import "fmt"

func intToRoman(num int) string {
	/*
		METHOD: Use arrays & math
		TIME COMPLEXITY: O(1), так как независимо от входного числа мы всегда выполняем ограниченное
		SPACE COMPLEXITY: O(1), так как мы используем фиксированное количество памяти для хранения массивов.
	*/
	// Массивы для римских цифр и их значений
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""
	// Проходим по всем значениям
	for i := 0; i < len(values); i++ {
		// Пока значение меньше или равно num, вычитаем его из num и добавляем соответствующую римскую цифру в результат
		for num >= values[i] {
			num -= values[i]
			result += romans[i]
		}
	}
	return result
}

func main() {
	fmt.Println(intToRoman(3))    // Вывод: "III"
	fmt.Println(intToRoman(4))    // Вывод: "IV"
	fmt.Println(intToRoman(9))    // Вывод: "IX"
	fmt.Println(intToRoman(58))   // Вывод: "LVIII"
	fmt.Println(intToRoman(1994)) // Вывод: "MCMXCIV"
}
