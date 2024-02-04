package main

import (
	"fmt"
	"strconv"
)

// getHint функция, которая принимает секретное число и предположение игрока,
// и возвращает строку с количеством "быков" и "коров".
func getHint(secret string, guess string) string {
	/*
		METHOD: Count / Подсчет
		TIME COMPLEXITY: O(n), где n - количество цифр в секретном числе. Мы проходим по каждой цифре ровно один раз.
		SPACE COMPLEXITY: O(1), так как используется фиксированное количество пространства для счетчиков, независимо от размера входных данных.

		Решение основывается на подсчете количества быков и коров для каждой цифры в предположении.
		Если цифра совпадает и находится на правильной позиции, она считается быком.
		Если цифра есть в секретном числе, но находится на неправильной позиции, она считается коровой.
	*/
	// Инициализируем счетчики для быков и коров.
	bulls, cows := 0, 0
	// Создаем два массива для отслеживания количества цифр в секретном числе и предположении.
	secretCounts := make([]int, 10)
	guessCounts := make([]int, 10)

	// Проходим по каждой цифре в секретном числе и предположении.
	for i := 0; i < len(secret); i++ {
		// Если цифры совпадают и находятся на одинаковой позиции, увеличиваем счетчик быков.
		if secret[i] == guess[i] {
			bulls++
		} else {
			// Иначе увеличиваем счетчики для соответствующих цифр в секретном числе и предположении.
			secretCounts[secret[i]-'0']++
			guessCounts[guess[i]-'0']++
		}
	}

	// Проходим по всем цифрам от 0 до 9 и суммируем минимальное количество совпадающих цифр.
	for i := 0; i < 10; i++ {
		cows += min(secretCounts[i], guessCounts[i])
	}

	// Возвращаем ответ в формате "xAyB", где x - количество быков, а y - количество коров.
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}

// Функция для нахождения минимального из двух чисел.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Пример использования функции getHint.
	secret := "1807"
	guess := "7810"
	result := getHint(secret, guess)
	fmt.Println(result) // Вывод: "1A3B"
}
