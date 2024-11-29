package main

func balancedStringSplitV2(s string) int {
	/*
	   METHOD: Stack-based Approach
	   В этом коде мы используем стек для отслеживания баланса символов 'L' и 'R'.

	   В цикле мы проходим по каждому символу в строке s.
	   Если текущий символ совпадает с вершиной стека (или стек пуст), то мы добавляем символ в стек.
	   Если текущий символ не совпадает с вершиной стека, то мы удаляем вершину стека.
	   Если стек становится пустым, то мы нашли сбалансированную подстроку.

	   В конце мы возвращаем количество сбалансированных подстрок.

	   TIME COMPLEXITY: O(n)

	   SPACE COMPLEXITY: O(n)
	*/
	// Инициализируем стек и счетчик сбалансированных подстрок
	stack := make([]rune, 0)
	balancedCount := 0

	// Проходим по каждому символу в строке
	for _, char := range s {
		// Если стек пуст или вершина стека совпадает с текущим символом, добавляем символ в стек
		if len(stack) == 0 || stack[len(stack)-1] == char {
			stack = append(stack, char)
		} else {
			// Иначе, удаляем вершину стека
			stack = stack[:len(stack)-1]
		}

		// Если стек становится пустым, это означает, что мы нашли сбалансированную подстроку
		if len(stack) == 0 {
			balancedCount++
		}
	}

	// Возвращаем количество сбалансированных подстрок
	return balancedCount
}

// func main() {
// 	// Примеры использования функции
// 	s1 := "RLRRLLRLRL"
// 	fmt.Println(balancedStringSplitV2(s1)) // Вывод: 4

// 	s2 := "RLLLLRRRLR"
// 	fmt.Println(balancedStringSplitV2(s2)) // Вывод: 3

// 	s3 := "LLLLRRRR"
// 	fmt.Println(balancedStringSplitV2(s3)) // Вывод: 1
// }
