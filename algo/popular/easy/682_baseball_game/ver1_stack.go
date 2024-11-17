package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	/*
		METHOD: Stack
		- Используем стек для хранения результатов.
		- Проходим по каждой операции в списке и выполняем соответствующие действия:
		  - "+" добавляет сумму двух последних элементов стека.
		  - "D" добавляет удвоенный последний элемент стека.
		  - "C" удаляет последний элемент стека.
		  - Число добавляется в стек после преобразования из строки в целое число.
		- В конце суммируем все элементы в стеке.

		TIME COMPLEXITY: O(n)
		- O(n), где n — количество операций в списке.
		- Каждая операция обрабатывается за постоянное время O(1).

		SPACE COMPLEXITY: O(n)
		- O(n), где n — количество операций в списке.
		- В худшем случае все операции будут числами, и стек будет содержать n элементов.
	*/
	// Создаем стек для хранения результатов
	stack := []int{}

	// Проходим по каждой операции в списке
	for _, op := range operations {
		switch op {
		case "+":
			// Добавляем сумму двух последних элементов стека
			last := len(stack) - 1
			stack = append(stack, stack[last]+stack[last-1])
		case "D":
			// Добавляем удвоенный последний элемент стека
			last := len(stack) - 1
			stack = append(stack, 2*stack[last])
		case "C":
			// Удаляем последний элемент стека
			stack = stack[:len(stack)-1]
		default:
			// Преобразуем строку в число и добавляем в стек
			num, _ := strconv.Atoi(op)
			stack = append(stack, num)
		}
	}

	// Вычисляем сумму всех элементов в стеке
	sum := 0
	for _, val := range stack {
		sum += val
	}

	return sum
}

func main() {
	// Пример использования
	ops := []string{"5", "2", "C", "D", "+"}
	result := calPoints(ops)
	fmt.Println(result) // Вывод: 30
}
