package main

func makeGoodV2(s string) string {
	/*
		METHOD: Fixed array
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	stack := make([]rune, len(s)) // Массив фиксированного размера
	top := 0                      // Указатель на вершину стека

	for _, char := range s {
		// Если стек не пуст и текущий символ образует "плохую" пару с верхним элементом стека
		if top > 0 && abs(int(stack[top-1])-int(char)) == 32 {
			top-- // Удаляем верхний элемент стека
		} else {
			// Добавляем текущий символ в стек
			stack[top] = char
			top++
		}
	}

	// Преобразуем массив в строку
	return string(stack[:top])
}

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }
