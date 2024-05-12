package main

import "fmt"

// Функция для нахождения максимального числа в массиве после операций добавления.
// В функции maxCount аргументы m и n представляют размеры массива, на котором выполняются операции добавления.
// m представляет количество строк в массиве, а n - количество столбцов.
func maxCount(m int, n int, ops [][]int) int {
	/*
		METHOD: Iterate
		Для решения этой задачи мы можем использовать подход, основанный на нахождении пересечения всех диапазонов добавления.
		Поскольку каждая операция добавления может увеличить значение любого элемента в массиве, наибольшее число,
		которое может быть в массиве, будет равно размеру самого маленького диапазона добавления.

		TIME COMPLEXITY: O(k), где k - количество операций добавления.
		Это обусловлено тем, что мы проходим по каждой операции только один раз.

		SPACE COMPLEXITY: O(1), так как мы используем некоторые переменные для хранения минимальных значений,
		но не сохраняем дополнительные структуры данных, которые бы увеличили использование памяти с ростом входных данных.
	*/
	// Если нет операций, то максимальное число - это размер массива
	if len(ops) == 0 {
		return m * n
	}

	// Инициализируем минимальные значения для строк и столбцов
	minRow, minCol := ops[0][0], ops[0][1]

	// Проходим по всем операциям и находим минимальные значения для строк и столбцов
	for _, op := range ops {
		if op[0] < minRow {
			minRow = op[0]
		}
		if op[1] < minCol {
			minCol = op[1]
		}
	}

	// Максимальное число - это произведение минимального количества строк и столбцов
	return minRow * minCol
}

func main() {
	// Пример использования функции
	ops := [][]int{{2, 2}, {3, 3}}
	fmt.Println(maxCount(3, 3, ops)) // Вывод: 4
}
