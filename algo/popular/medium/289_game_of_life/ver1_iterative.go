package main

import (
	"fmt"
)

// Функция gameOfLife обновляет состояние сетки в соответствии с правилами "Игры жизни"
func gameOfLife(board [][]int) {
	/*
		METHOD:  Iterative
		Итеративный обновление состояния сетки с использованием временной копии"
		или "Императивный подход с использованием двумерных массивов и временной копии".
		Этот подход является классическим для решения задач, связанных с клеточными автоматами,
		и хорошо подходит для языка Go благодаря его эффективной работе с массивами и циклами.

		TIME COMPLEXITY: O(n*m), где n — количество строк, m — количество столбцов в сетке.
		Это связано с тем, что мы обходим каждую клетку сетки и для каждой клетки считаем живых соседей.

		SPACE COMPLEXITY: O(n*m), так как мы создаем новую сетку того же размера для хранения следующего состояния.
	*/
	rows := len(board)
	if rows == 0 {
		return
	}
	cols := len(board[0])

	// Функция для подсчета живых соседей
	countLiveNeighbors := func(x, y int) int {
		directions := [][2]int{
			{-1, -1}, {-1, 0}, {-1, 1},
			{0, -1}, {0, 1},
			{1, -1}, {1, 0}, {1, 1},
		}
		count := 0
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newX < rows && newY >= 0 && newY < cols && board[newX][newY] == 1 {
				count++
			}
		}
		return count
	}

	// Создаем копию сетки для хранения следующего состояния
	nextBoard := make([][]int, rows)
	for i := range nextBoard {
		nextBoard[i] = make([]int, cols)
	}

	// Обходим каждую клетку сетки
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			liveNeighbors := countLiveNeighbors(x, y)
			// Применяем правила игры жизни
			if board[x][y] == 1 && (liveNeighbors == 2 || liveNeighbors == 3) {
				nextBoard[x][y] = 1
			} else if board[x][y] == 0 && liveNeighbors == 3 {
				nextBoard[x][y] = 1
			}
		}
	}

	// Копируем следующее состояние обратно в исходную сетку
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			board[x][y] = nextBoard[x][y]
		}
	}
}

func main() {
	// Пример начального состояния сетки
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	// Выводим начальное состояние
	fmt.Println("Initial state:")
	for _, row := range board {
		fmt.Println(row)
	}

	// Обновляем состояние сетки
	gameOfLife(board)

	// Выводим следующее состояние
	fmt.Println("Next state:")
	for _, row := range board {
		fmt.Println(row)
	}
}
