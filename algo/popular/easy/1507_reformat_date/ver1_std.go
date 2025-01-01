package main

import (
	"fmt"
	"strings"
)

func reformatDate(date string) string {
	/*
		METHOD: Split, Map, Join
		- Split: Разделение строки на компоненты (день, месяц, год) по пробелам.
		- Map: Использование словаря для сопоставления названия месяца с его числовым значением.
		- Join: Объединение года, месяца и дня в строку формата YYYY-MM-DD с помощью fmt.Sprintf.

		TIME COMPLEXITY: O(1)
		- Разделение строки: O(1), так как строка фиксированной длины.
		- Удаление суффикса и добавление ведущего нуля: O(1).
		- Поиск месяца в словаре: O(1).
		- Форматирование результата: O(1).

		SPACE COMPLEXITY: O(1)
		- Используется фиксированное количество дополнительной памяти:
		  - Словарь для месяцев (константный размер).
		  - Временные переменные для дня, месяца и года.
	*/
	// Разделяем строку на компоненты: день, месяц, год
	parts := strings.Split(date, " ")
	day, month, year := parts[0], parts[1], parts[2]

	// Убираем суффикс из дня (последние два символа)
	day = day[:len(day)-2]
	// Добавляем ведущий ноль, если день состоит из одной цифры
	if len(day) == 1 {
		day = "0" + day
	}

	// Сопоставляем месяц с его числовым значением
	monthMap := map[string]string{
		"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04",
		"May": "05", "Jun": "06", "Jul": "07", "Aug": "08",
		"Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12",
	}
	monthNum := monthMap[month]

	// Собираем результат в формате YYYY-MM-DD
	return fmt.Sprintf("%s-%s-%s", year, monthNum, day)
}

func main() {
	// Пример использования
	date := "20th Oct 2052"
	formattedDate := reformatDate(date)
	fmt.Println(formattedDate) // Вывод: 2052-10-20
}
