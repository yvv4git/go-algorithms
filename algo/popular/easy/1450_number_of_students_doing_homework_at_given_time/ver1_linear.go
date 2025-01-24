package main

import "fmt"

// Функция для подсчёта количества студентов, выполняющих домашнюю работу в заданное время
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	/*
		METHOD: Linear scan
		- Мы проходим по каждому студенту и проверяем, попадает ли queryTime в интервал [startTime[i], endTime[i]].
		- Если queryTime находится в этом интервале, увеличиваем счётчик.
		- В конце возвращаем значение счётчика.

		TIME COMPLEXITY:
		- O(n), где n — количество студентов.
		- Мы проходим по каждому элементу массивов startTime и endTime ровно один раз, выполняя константные операции для каждого студента.

		SPACE COMPLEXITY:
		- O(1), так как мы используем только константное количество дополнительной памяти (переменная count).
		- Мы не создаём никаких дополнительных структур данных, которые зависят от размера входных данных.
	*/
	count := 0 // Счётчик студентов

	// Проходим по каждому студенту
	for i := 0; i < len(startTime); i++ {
		// Проверяем, находится ли queryTime в интервале [startTime[i], endTime[i]]
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			count++ // Увеличиваем счётчик, если условие выполняется
		}
	}

	return count // Возвращаем итоговое количество студентов
}

func main() {
	// Пример входных данных
	startTime := []int{1, 2, 3}
	endTime := []int{3, 2, 7}
	queryTime := 4

	// Вызов функции и вывод результата
	result := busyStudent(startTime, endTime, queryTime)
	fmt.Println(result) // Ожидаемый результат: 1
}
