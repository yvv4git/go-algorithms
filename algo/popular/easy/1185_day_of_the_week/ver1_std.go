//go:build ignore

package main

import (
	"fmt"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	/*
		METHOD:
		- Используется встроенный пакет `time` для создания объекта даты и получения дня недели.
		- Создается объект `time.Time` для заданной даты.
		- Метод `Weekday()` используется для получения дня недели, а метод `String()` для преобразования его в строку.

		TIME COMPLEXITY:
		- O(1) — Создание объекта даты и получение дня недели выполняются за константное время, так как это операции, поддерживаемые внутренними структурами данных пакета `time`.

		SPACE COMPLEXITY:
		- O(1) — Используется фиксированное количество памяти для хранения объекта даты и результата.
	*/
	// Создаем объект даты
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	// Получаем день недели
	dayOfWeek := date.Weekday().String()

	return dayOfWeek
}

func main() {
	// Пример использования функции
	day := 31
	month := 8
	year := 2019

	result := dayOfTheWeek(day, month, year)
	fmt.Println(result) // Вывод: Saturday
}
