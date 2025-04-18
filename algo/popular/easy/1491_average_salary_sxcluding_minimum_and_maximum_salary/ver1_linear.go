package main

import (
	"fmt"
)

func average(salary []int) float64 {
	/*
		APPROACH: Linear scan
		- Мы проходим по массиву один раз, чтобы:
			1. Найти минимальную зарплату (minSalary)
			2. Найти максимальную зарплату (maxSalary)
			3. Посчитать сумму всех зарплат
		- Затем вычитаем minSalary и maxSalary из общей суммы и вычисляем среднее арифметическое.

		TIME COMPLEXITY: O(n)
		- Мы выполняем один проход по массиву (один цикл for), где n — количество элементов в `salary`.
		- В каждой итерации выполняются только константные операции (сравнения, сложение).
		- Поэтому временная сложность составляет O(n).

		SPACE COMPLEXITY: O(1)
		- Мы используем только несколько переменных (`minSalary`, `maxSalary`, `sum`), независимо от размера входных данных.
		- Мы **не создаем** дополнительный массив или структуру данных.
		- Следовательно, пространственная сложность равна O(1).
	*/
	// Инициализируем переменные для поиска минимальной и максимальной зарплаты
	minSalary, maxSalary := salary[0], salary[0]
	sum := 0

	// Проходим по всем зарплатам
	for _, s := range salary {
		// Находим минимальное значение
		if s < minSalary {
			minSalary = s
		}
		// Находим максимальное значение
		if s > maxSalary {
			maxSalary = s
		}
		// Суммируем все зарплаты
		sum += s
	}

	// Вычисляем сумму без минимального и максимального значения
	sumExcludingMinMax := sum - minSalary - maxSalary

	// Возвращаем среднее арифметическое
	return float64(sumExcludingMinMax) / float64(len(salary)-2)
}

func main() {
	// Пример входных данных
	salary := []int{4000, 3000, 1000, 2000}

	// Вычисляем среднюю зарплату и выводим результат
	fmt.Println(average(salary)) // Ожидаемый результат: 2500.0
}
