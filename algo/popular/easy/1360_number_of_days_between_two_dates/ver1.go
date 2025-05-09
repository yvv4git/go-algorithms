package main

import (
	"fmt"
	"time"
)

// Функция для вычисления количества дней между двумя датами
func daysBetweenDates(date1 string, date2 string) int {
	/*
		METHOD: Direct calculation
		- Функция принимает две даты в формате строки "YYYY-MM-DD".
		- Используется стандартная библиотека `time` для парсинга строк в объекты `time.Time`.
		- Разница между датами вычисляется с помощью метода `Sub`, который возвращает разницу в виде `time.Duration`.
		- Разница преобразуется в часы, а затем в дни.
		- Результат всегда неотрицательный, так как используется модуль.

		TIME COMPLEXITY: O(1)
		- Парсинг строки в `time.Time` выполняется за константное время.
		- Вычисление разницы между датами и преобразование в дни также выполняется за константное время.
		- Все операции не зависят от размера входных данных.

		SPACE COMPLEXITY: O(1)
		- Используется фиксированное количество переменных (d1, d2, days), которые не зависят от размера входных данных.
		- Дополнительная память не требуется.
	*/
	// Парсим первую дату из строки в формате "YYYY-MM-DD"
	// time.Parse использует макет "2006-01-02" для разбора дат
	d1, err := time.Parse("2006-01-02", date1)
	if err != nil {
		panic(err) // В случае ошибки парсинга программа завершится
	}

	// Парсим вторую дату
	d2, err := time.Parse("2006-01-02", date2)
	if err != nil {
		panic(err)
	}

	// Вычисляем разницу между двумя датами
	// Метод Sub возвращает разницу в виде time.Duration
	// Метод Hours() возвращает количество часов, делим на 24 для получения дней
	days := int(d2.Sub(d1).Hours() / 24)

	// Используем модуль, чтобы результат был неотрицательным
	if days < 0 {
		days = -days
	}

	return days
}

func main() {
	// Пример использования
	date1 := "2019-06-29"
	date2 := "2019-06-30"
	fmt.Println(daysBetweenDates(date1, date2)) // Вывод: 1

	// Другой пример
	date3 := "2020-01-01"
	date4 := "2020-12-31"
	fmt.Println(daysBetweenDates(date3, date4)) // Вывод: 365 (2020 год — високосный)
}
