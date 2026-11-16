//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// Структура для хранения информации о лекции
type Lecture struct {
	start int // Время начала
	end   int // Время окончания
}

// Функция для выбора максимального количества непересекающихся лекций
func maxLectures(lectures []Lecture) []Lecture {
	/*
		METHOD: Greedy algorithm
		1. Сортируем лекции по времени окончания.
		2. Проходим по лекциям, выбирая непересекающиеся.

		TIME COMPLEXITY: O(nlogn) (из-за сортировки)
		SPACE COMPLEXITY: O(n) (для хранения результата)
	*/

	// Сортируем лекции по времени окончания
	sort.Slice(lectures, func(i, j int) bool {
		return lectures[i].end < lectures[j].end
	})

	var selected []Lecture
	lastEndTime := 0

	// Проходим по лекциям и выбираем непересекающиеся
	for _, lecture := range lectures {
		if lecture.start >= lastEndTime {
			selected = append(selected, lecture)
			lastEndTime = lecture.end
		}
	}

	return selected
}

// Дано N лекций, каждая начинается в s_i и заканчивается в f_i. Нужно выбрать максимальное количество лекций, которые не пересекаются.
func main() {
	// Задаем список лекций (время начала, время окончания)
	lectures := []Lecture{
		{1, 3}, {2, 5}, {3, 6}, {5, 8}, {8, 9}, {5, 7},
	}

	// Вызываем жадный алгоритм
	selectedLectures := maxLectures(lectures)

	// Выводим результат
	fmt.Println("Выбранные лекции:")
	for _, lecture := range selectedLectures {
		fmt.Printf("Начало: %d, Конец: %d\n", lecture.start, lecture.end)
	}
}
