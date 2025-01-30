package main

import "fmt"

func numRookCaptures(board [][]byte) int {
	/*
		METHO: Linear scan
		1. Найти положение ладьи на доске
		2. Проверить все четыре направления: вверх, вниз, влево, вправо
		3. Подсчитать захваты в каждом направлении
		4. Вернуть количество захватов

		TIME COMPLEXITY: O(1)
		SPACE COMPLEXITY: O(1)
	*/
	// Найдем положение ладьи на доске
	var rookX, rookY int
	found := false
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				rookX, rookY = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	captures := 0

	// Функция для проверки захватов в одном направлении
	checkDirection := func(dx, dy int) {
		x, y := rookX, rookY
		for {
			x += dx
			y += dy
			if x < 0 || x >= 8 || y < 0 || y >= 8 {
				break // Вышли за пределы доски
			}
			if board[x][y] == 'p' {
				captures++ // Зацепили пешку
				break
			} else if board[x][y] != '.' {
				break // Встретили другую фигуру
			}
		}
	}

	// Проверим все четыре направления: вверх, вниз, влево, вправо
	checkDirection(-1, 0) // Вверх
	checkDirection(1, 0)  // Вниз
	checkDirection(0, -1) // Влево
	checkDirection(0, 1)  // Вправо

	return captures
}

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}

	fmt.Println(numRookCaptures(board)) // Выведет 3
}
