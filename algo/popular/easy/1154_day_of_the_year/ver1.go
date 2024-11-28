package main

import (
	"fmt"
	"time"
)

func main() {
	// Пример даты
	date := "15.03.2023"

	// Вызываем функцию для определения дня года
	dayOfYear := getDayOfYear(date)

	// Выводим результат
	fmt.Printf("День года для даты %s: %d\n", date, dayOfYear)
}

// Функция для определения дня года
func getDayOfYear(date string) int {
	/*
		METHOD: Loop
		TIME COMPLEXITY: O(1)
		SPACE COMPLEXITY: O(1)
	*/
	// Парсим дату в формате "день.месяц.год"
	t, err := time.Parse("02.01.2006", date)
	if err != nil {
		fmt.Println("Ошибка при парсинге даты:", err)
		return -1
	}

	// Получаем первый день года
	yearStart := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.UTC)

	// Вычисляем разницу в днях между датой и первым днем года
	daysDifference := t.Sub(yearStart).Hours() / 24

	// Возвращаем день года
	return int(daysDifference) + 1
}
