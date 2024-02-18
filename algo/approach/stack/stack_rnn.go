package main

import (
	"strconv"
)

// Функция для вычисления выражения в постфиксной нотации
func evaluateRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		// Проверяем, является ли токен числом
		if num, err := strconv.Atoi(token); err == nil {
			// Если токен число, то добавляем его в стек
			stack = append(stack, num)
		} else {
			// Если токен оператор, то извлекаем два последних числа из стека
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2] // Удаляем из стека последние два числа

			// Выполняем операцию и добавляем результат обратно в стек
			switch token {
			case "+":
				stack = append(stack, operand1+operand2)
			case "-":
				stack = append(stack, operand1-operand2)
			case "*":
				stack = append(stack, operand1*operand2)
			case "/":
				stack = append(stack, operand1/operand2)
			}
		}
	}

	// В конце стек должен содержать единственное число - результат вычисления
	return stack[0]
}

//func main() {
//	// Пример использования функции
//	expression := "2 3 + 4 *"
//	tokens := strings.Split(expression, " ")
//	result := evaluateRPN(tokens)
//	fmt.Printf("The result of the expression '%s' is: %d\n", expression, result)
//}
