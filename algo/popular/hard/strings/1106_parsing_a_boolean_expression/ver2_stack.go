package main

// Функция для разбора логического выражения
func parseBoolExprV2(expression string) bool {
	/*
		METHOD: Stack & Recursion
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность функции parseBoolExpr и helper в худшем случае составляет O(n),
		где n - количество символов в выражении. Это связано с тем, что функция проходит по каждому символу в выражении ровно один раз.

		Space complexity
		Пространственная сложность также составляет O(n) в худшем случае.
		Это связано с тем, что в худшем случае функция может создать список элементов размером в 1024,
		который может увеличиться, если в выражении будет больше элементов.
		Однако, в общем случае пространственная сложность будет меньше, так как список элементов будет создан только в том случае,
		если в выражении есть операторы.
	*/
	// Если длина выражения равна 1, то проверяем, является ли он истинным
	if len(expression) == 1 {
		return expression[0] == 't'
	}
	// Вызываем вспомогательную функцию для разбора выражения и получаем результат
	result, _ := helper([]byte(expression), 0)
	return result
}

// Вспомогательная функция для разбора логического выражения
func helper(expression []byte, index int) (bool, int) {
	// expression[index]:   должен быть '&', '|' или '!'
	// expression[index+1]: должен быть '('
	operator := expression[index]
	symbols := make([]bool, 0, 1024)
	retNext := 0
	for i := index + 2; i < len(expression); {
		if expression[i] == ')' {
			// этот оператор завершился
			retNext = i + 1
			break
		}

		if expression[i] == ',' {
			// есть еще элементы
			i++
		}
		if expression[i] == '!' || expression[i] == '&' || expression[i] == '|' {
			// если следующий символ является оператором, вызываем рекурсивно вспомогательную функцию
			result, nextIndex := helper(expression, i)
			symbols = append(symbols, result)
			i = nextIndex
		} else {
			// если следующий символ не является оператором, добавляем его в список элементов
			symbols = append(symbols, expression[i] == 't')
			i++
		}
	}

	// Логические операции
	retResult := false
	switch operator {
	case '!':
		// логическое НЕ
		retResult = !symbols[0]
	case '&':
		// логическое И
		retResult = true
		for _, v := range symbols {
			retResult = retResult && v
		}
	case '|':
		// логическое ИЛИ
		retResult = false
		for _, v := range symbols {
			retResult = retResult || v
		}
	}

	return retResult, retNext
}
