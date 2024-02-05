package main

import (
	"math"
)

// Функция для вычисления расстояния (различия) между двумя строками
func distance(a, b string) int {
	var count int
	// Проходим по всем символам строк
	for i := 0; i < len(a); i++ {
		// Если символы не совпадают, увеличиваем счетчик
		if a[i] != b[i] {
			count++
		}
	}
	// Возвращаем количество различий (расстояние)
	return count
}

// Функция для обратного отслеживания, реализующая алгоритм поиска
func backTracking(start string, end string, bank []string, steps int) int {
	var newBank []string
	minSteps := math.MaxInt32 // Инициализируем минимальное количество шагов как максимальное значение int32

	// Проходим по всем элементам банка
	for i := 0; i < len(bank); i++ {
		// Если расстояние между начальной строкой и элементом банка равно 1
		if distance(start, bank[i]) == 1 {
			// Если элемент банка совпадает с конечной строкой, возвращаем текущее количество шагов + 1
			if bank[i] == end {
				return steps + 1
			}
			// Создаем новый банк без текущего элемента
			newBank = make([]string, len(bank)-1)
			copy(newBank, bank[:i])
			copy(newBank[i:], bank[i+1:])
			// Рекурсивно вызываем backTracking для нового банка
			result := backTracking(bank[i], end, newBank, steps+1)

			// Если результат не -1 и меньше текущего минимального количества шагов, обновляем минимальное количество шагов
			if result != -1 && result < minSteps {
				minSteps = result
			}
		}
	}
	// Если минимальное количество шагов осталось равным максимальному значению int32, возвращаем -1, так как преобразование невозможно
	if minSteps == math.MaxInt32 {
		return -1
	}
	// Возвращаем минимальное количество шагов
	return minSteps
}

// Функция для поиска минимального количества мутаций
func minMutationV2(start string, end string, bank []string) int {
	// Вызываем функцию backTracking с начальной строкой, конечной строкой, банком и 0 шагов
	return backTracking(start, end, bank, 0)
}

//func main() {
//	start := "AACCGGTT"
//	end := "AACCGGTA"
//	bank := []string{"AACCGGTA"}
//	fmt.Println(minMutationV2(start, end, bank)) // Выводит: 1
//}
