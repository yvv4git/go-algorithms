package main

import (
	"fmt"
)

// Функция для нахождения наибольшего общего делителя строк
func gcdOfStrings(str1 string, str2 string) string {
	/*
		METHOD: Math
		Используется математический подход, основанный на свойствах НОД и конкатенации строк.

		TIME COMPLEXITY: O(n + m), где n и m — длины строк str1 и str2.
		- Конкатенация строк str1 + str2 и str2 + str1 занимает O(n + m).
		- Вычисление НОД длин строк занимает O(log(min(n, m))).
		- Общая сложность: O(n + m).

		SPACE COMPLEXITY: O(n + m)
		- Для хранения результатов конкатенации строк str1 + str2 и str2 + str1 требуется O(n + m) дополнительной памяти.
		- Вычисление НОД требует O(1) дополнительной памяти.
		- Общая сложность: O(n + m).
	*/
	// Проверяем, могут ли строки быть сформированы из общей строки.
	// Если str1 + str2 не равно str2 + str1, то такой строки не существует.
	if str1+str2 != str2+str1 {
		return ""
	}

	// Находим НОД длин строк str1 и str2.
	// Это будет длина наибольшей общей строки.
	gcdLength := gcd(len(str1), len(str2))

	// Возвращаем подстроку str1 длиной gcdLength.
	return str1[:gcdLength]
}

// Вспомогательная функция для нахождения НОД двух чисел
func gcd(a, b int) int {
	/*
		METHOD: Euclidean Algorithm

		TIME COMPLEXITY: O(log(min(a, b)))
		- Алгоритм Евклида работает за время, пропорциональное логарифму меньшего из чисел a и b.

		SPACE COMPLEXITY: O(1)
		- Используется только несколько переменных, поэтому дополнительная память не требуется.
	*/
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Пример использования
func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	result := gcdOfStrings(str1, str2)
	fmt.Println(result) // Вывод: "ABC"

	str1 = "ABABAB"
	str2 = "ABAB"
	result = gcdOfStrings(str1, str2)
	fmt.Println(result) // Вывод: "AB"

	str1 = "LEET"
	str2 = "CODE"
	result = gcdOfStrings(str1, str2)
	fmt.Println(result) // Вывод: ""
}
