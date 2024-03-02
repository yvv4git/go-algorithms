package main

// Структура position используется для хранения координат ячейки в матрице.
type position struct {
	y, x int
}

// Функция generateMatrixV2 генерирует двумерную матрицу размера n x n, заполненную по спирали.
func generateMatrixV2(n int) [][]int {
	/*
		METHOD: Iterative
		TIME COMPLEXITY: O(n^2)
		SPACE COMPLEXITY: O(n^2)
	*/
	// Инициализация матрицы заполненной нулями.
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	// Максимальное значение для заполнения матрицы.
	maxValue := n * n
	// Текущая позиция для заполнения.
	pos := position{0, 0}
	// Направление движения по спирали.
	direction := "right"

	// Заполнение матрицы по спирали.
	for i := 1; i <= maxValue; i++ {
		// Проверка, можно ли продолжить движение в текущем направлении.
		if canContinue(&matrix, pos.y, pos.x) {
			matrix[pos.y][pos.x] = i
		}

		// Изменение позиции и направления движения в зависимости от текущего направления.
		switch direction {
		case "up":
			// Проверка, можно ли двигаться вверх.
			if !canContinue(&matrix, pos.y-1, pos.x) {
				direction = "right"
				pos.x++
				continue
			}
			pos.y--
		case "right":
			// Проверка, можно ли двигаться вправо.
			if !canContinue(&matrix, pos.y, pos.x+1) {
				direction = "down"
				pos.y++
				continue
			}
			pos.x++
		case "down":
			// Проверка, можно ли двигаться вниз.
			if !canContinue(&matrix, pos.y+1, pos.x) {
				direction = "left"
				pos.x--
				continue
			}
			pos.y++
		case "left":
			// Проверка, можно ли двигаться влево.
			if !canContinue(&matrix, pos.y, pos.x-1) {
				direction = "up"
				pos.y--
				continue
			}
			pos.x--
		}
	}

	return matrix
}

// Функция canContinue проверяет, можно ли продолжить движение в заданной позиции.
func canContinue(matrix *[][]int, posY, posX int) bool {
	n := len(*matrix)

	// Проверка, находится ли позиция в пределах матрицы и ячейка не заполнена.
	if posY < 0 || posY == n || posX < 0 || posX == n || (*matrix)[posY][posX] != 0 {
		return false
	}

	return true
}

//func main() {
//	// Размер матрицы.
//	n := 3
//	// Получение заполненной по спирали матрицы.
//	result := generateMatrixV2(n)
//	// Вывод матрицы на экран.
//	for _, row := range result {
//		fmt.Println(row)
//	}
//}
