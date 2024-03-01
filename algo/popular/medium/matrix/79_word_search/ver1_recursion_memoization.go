package main

import "fmt"

func exist(board [][]byte, word string) bool {
	/*
		METHOD: Recursion with memoization
		Мы будем проверять каждую ячейку на соответствие первой букве слова, и, если она совпадает,
		мы будем проверять остальные буквы слова в соседних ячейках.

		TIME COMPLEXITY: O(N), где N - общее количество ячеек в двумерном массиве.
		Это связано с тем, что мы проходим по каждой ячейке массива только один раз.

		SPACE COMPLEXITY: O(N), так как в худшем случае мы можем посетить каждую ячейку массива.
		Это происходит, когда слово состоит из одинаковых букв, и мы можем вернуться к предыдущим ячейкам.
	*/
	// Проверяем, что board и word не пустые
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}

	// Создаем двумерный массив для отслеживания посещенных ячеек
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	// Проходим по каждой ячейке и ищем слово
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0, visited) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, i, j, index int, visited [][]bool) bool {
	// Если индекс равен длине слова, значит, мы нашли слово
	if index == len(word) {
		return true
	}

	// Если координаты выходят за пределы массива или ячейка уже посещена или буквы не совпадают, возвращаем false
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || visited[i][j] || board[i][j] != word[index] {
		return false
	}

	// Помечаем ячейку как посещенную
	visited[i][j] = true

	// Проверяем соседние ячейки
	if dfs(board, word, i-1, j, index+1, visited) ||
		dfs(board, word, i+1, j, index+1, visited) ||
		dfs(board, word, i, j-1, index+1, visited) ||
		dfs(board, word, i, j+1, index+1, visited) {
		return true
	}

	// Если слово не найдено, отмечаем ячейку как непосещенную и возвращаем false
	visited[i][j] = false
	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word)) // Выводит: true
}
