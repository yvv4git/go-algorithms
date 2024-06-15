package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	/*
		METHOD: Set
		Для решения этой задачи мы будем использовать простой подход с использованием множеств (set) для отслеживания повторяющихся чисел в каждой строке, столбце и квадрате.
		Мы будем проверять каждую строку, столбец и квадрат на наличие повторяющихся чисел. Если находим повторяющееся число, возвращаем false.
		Если мы успешно проверили все строки, столбцы и квадраты и не нашли повторяющихся чисел, возвращаем true.

		TIME COMPLEXITY: O(1), так как независимо от размера входных данных мы всегда будем выполнять одинаковое количество операций.

		SPACE COMPLEXITY: O(1), так как мы используем фиксированное количество дополнительной памяти для отслеживания повторяющихся чисел.
	*/
	for i := 0; i < 9; i++ {
		row := make(map[byte]bool)  // множество для проверки строки
		col := make(map[byte]bool)  // множество для проверки столбца
		cube := make(map[byte]bool) // множество для проверки квадрата

		for j := 0; j < 9; j++ {
			// Проверка строки
			if board[i][j] != '.' {
				if row[board[i][j]] {
					return false
				}
				row[board[i][j]] = true
			}

			// Проверка столбца
			if board[j][i] != '.' {
				if col[board[j][i]] {
					return false
				}
				col[board[j][i]] = true
			}

			// Проверка квадрата
			rowIndex := 3*(i/3) + j/3
			colIndex := 3*(i%3) + j%3
			if board[rowIndex][colIndex] != '.' {
				if cube[board[rowIndex][colIndex]] {
					return false
				}
				cube[board[rowIndex][colIndex]] = true
			}
		}
	}

	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board)) // Вывод: true
}