package main

func solve(board [][]byte) {
	/*
		METHOD: DFS

		TIME COMPLEXITY: O(n * m), где n — количество строк, m — количество столбцов в матрице.
		Это связано с тем, что каждая ячейка матрицы будет посещена не более двух раз (один раз при маркировке и один раз при замене).

		SPACE COMPLEXITY: O(n * m) в худшем случае, если все 'O' соединены с границей и будут помечены.
		Это связано с тем, что в худшем случае стек вызовов рекурсии может достичь размера всей матрицы.
	*/
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	rows, cols := len(board), len(board[0])

	// Функция для DFS, чтобы пометить все 'O', которые соединены с границей
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'E' // Пометить как 'E' (escaped)
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	// Пройти по всем границам и запустить DFS для 'O'
	for r := 0; r < rows; r++ {
		if board[r][0] == 'O' {
			dfs(r, 0)
		}
		if board[r][cols-1] == 'O' {
			dfs(r, cols-1)
		}
	}
	for c := 0; c < cols; c++ {
		if board[0][c] == 'O' {
			dfs(0, c)
		}
		if board[rows-1][c] == 'O' {
			dfs(rows-1, c)
		}
	}

	// Пройти по всей матрице и заменить 'O' на 'X', а 'E' на 'O'
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == 'E' {
				board[r][c] = 'O'
			}
		}
	}
}
