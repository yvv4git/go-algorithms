package main

import (
	"strconv"
)

func parseBoolExpr(expression string) bool {
	// Инициализация стека для хранения символов выражения
	var stack []rune

	// Итерация по символам выражения
	for _, char := range expression {
		// Пропуск символов-разделителей и закрывающей скобки
		if char == ',' || char == ')' {
			continue
		}

		// Добавление логических операторов, констант и открывающей скобки в стек
		if char == 't' || char == 'f' || char == '!' || char == '&' || char == '|' {
			stack = append(stack, char)
		}

		// Обработка открывающей скобки
		if char == '(' {
			// Извлечение последнего символа из стека (оператор)
			operand := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Удаление из стека

			// Выполнение логической операции в зависимости от оператора
			if operand == '!' {
				// Логическое НЕ
				result := !parseBoolExpr(expression[1 : len(expression)-1])
				if result {
					stack = append(stack, 't')
				} else {
					stack = append(stack, 'f')
				}
			} else if operand == '&' {
				// Логическое И
				result := parseBoolExpr(expression[1 : len(expression)-1])
				if result {
					stack = append(stack, 't')
				} else {
					stack = append(stack, 'f')
				}
			} else if operand == '|' {
				// Логическое ИЛИ
				result := parseBoolExpr(expression[1 : len(expression)-1])
				if result {
					stack = append(stack, 't')
				} else {
					stack = append(stack, 'f')
				}
			}
		}
	}

	// Преобразование последнего символа в стеке в булевый результат
	result, _ := strconv.ParseBool(string(stack[0]))
	return result
}
