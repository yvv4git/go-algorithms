package main

import (
	"strings"
)

// isValidSerialization проверяет, является ли данная строка корректной сериализацией бинарного дерева в порядке preorder.
func isValidSerialization(preorder string) bool {
	/*
		METHOD: Stack-based
		Этот подход называется "стековым" (stack-based approach).
		Мы используем стек для отслеживания текущих узлов и проверки их дочерних элементов.
		Основная идея заключается в том, что каждый не-null узел должен иметь два дочерних элемента (которые могут быть null).

		Подробное объяснение
		1. Разбиение строки.
		Мы разбиваем входную строку на отдельные узлы, используя запятую в качестве разделителя.

		2. Использование стека.
		Мы добавляем каждый узел в стек. Если последние два элемента в стеке являются "#" (null узлы),
		а предпоследний элемент — не "#", это означает, что мы завершили поддерево и можем удалить эти три элемента и добавить "#" в стек.

		3. Проверка окончания.
		В конце проверки стек должен содержать только один элемент "#", что означает, что все дерево было корректно завершено.



		TIME COMPLEXITY: O(n), где n — количество узлов в сериализации.
		Мы проходим по каждому узлу один раз, и каждый узел может быть добавлен и удален из стека не более одного раза.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить все узлы в стеке.
	*/
	// Разбиваем строку на отдельные узлы
	nodes := strings.Split(preorder, ",")
	// Стек для отслеживания текущих узлов
	stack := []string{}

	// Проходим по каждому узлу в списке
	for _, node := range nodes {
		// Добавляем текущий узел в стек
		stack = append(stack, node)

		// Пока в стеке есть как минимум три элемента и последние два элемента - "#" (null узлы),
		// а предпоследний элемент - не "#"
		// Использование трех элементов в стеке связано с тем, что каждый не-null узел в бинарном дереве должен иметь два дочерних узла.
		for len(stack) >= 3 && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" && stack[len(stack)-3] != "#" {
			// Удаляем последние три элемента из стека
			stack = stack[:len(stack)-3]
			// Добавляем "#" в стек, чтобы обозначить завершение поддерева
			stack = append(stack, "#")
		}
	}

	// В конце стек должен содержать только один элемент "#", что означает корректное завершение дерева
	return len(stack) == 1 && stack[0] == "#"
}

// Пример использования
func main() {
	println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#")) // True
	println(isValidSerialization("1,#"))                       // False
	println(isValidSerialization("9,#,#,1"))                   // False
}
