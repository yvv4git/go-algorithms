//go:build ignore

package main

import (
	"fmt"
)

func dayOfTheWeek(day int, month int, year int) string {
	/*
		METHOD:
		- Используется формула Зеллера (Zeller's Congruence) для определения дня недели.
		- Формула учитывает год, месяц и день для вычисления дня недели.

		TIME COMPLEXITY:
		- O(1) — Вычисления по формуле Зеллера выполняются за константное время.

		SPACE COMPLEXITY:
		- O(1) — Используется фиксированное количество памяти для хранения промежуточных результатов и конечного результата.
	*/

	// Корректировка месяца и года для формулы Зеллера
	if month < 3 {
		month += 12
		year -= 1
	}

	// Вычисление столетия и года в столетии
	century := year / 100
	yearInCentury := year % 100

	// Формула Зеллера
	h := (day + 13*(month+1)/5 + yearInCentury + yearInCentury/4 + century/4 + 5*century) % 7

	// Преобразование результата в день недели
	daysOfWeek := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	return daysOfWeek[h]
}

func main() {
	// Пример использования функции
	day := 31
	month := 8
	year := 2019

	result := dayOfTheWeek(day, month, year)
	fmt.Println(result) // Вывод: Saturday
}
