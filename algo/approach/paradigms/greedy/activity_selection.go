package main

import (
	"fmt"
	"sort"
)

// Структура для хранения информации о мероприятии
type Activity struct {
	start int // Время начала
	end   int // Время окончания
}

// Функция для выбора максимального количества непересекающихся мероприятий
func activitySelection(activities []Activity) []Activity {
	/*
		METHOD: Greedy algorithm
		1. Сортируем мероприятия по времени окончания (чтобы быстрее освободить место для новых).
		2. Проходим по мероприятиям, выбирая непересекающиеся (т. е. следующее мероприятие должно начинаться после завершения предыдущего).

		TIME COMPLEXITY: O(nlogn)
		SPACE COMPLEXITY: O(n)
	*/
	// Сортируем мероприятия по времени окончания
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].end < activities[j].end
	})

	var selected []Activity
	lastEndTime := 0

	// Проходим по мероприятиям и выбираем непересекающиеся
	for _, activity := range activities {
		if activity.start >= lastEndTime {
			selected = append(selected, activity)
			lastEndTime = activity.end
		}
	}

	return selected
}

// Дано N мероприятий, каждое из которых имеет время начала sᵢ и время окончания fᵢ.
// Требуется выбрать максимальное количество непересекающихся мероприятий, то есть таких, которые не идут одновременно.
func main() {
	// Задаем список мероприятий (время начала, время окончания)
	activities := []Activity{
		{1, 3}, {2, 5}, {3, 9}, {6, 8}, {5, 7}, {8, 9},
	}

	// Вызываем жадный алгоритм
	selectedActivities := activitySelection(activities)

	// Выводим результат
	fmt.Println("Выбранные мероприятия:")
	for _, activity := range selectedActivities {
		fmt.Printf("Начало: %d, Конец: %d\n", activity.start, activity.end)
	}
}
