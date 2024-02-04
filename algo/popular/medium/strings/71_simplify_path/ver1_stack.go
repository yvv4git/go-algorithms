package main

import (
	"fmt"
	"strings"
)

// Функция simplifyPath упрощает путь Unix-подобной файловой системы.
func simplifyPath(path string) string {
	/*
		METHOD: Stack
		TIME COMPLEXITY: O(n), где n - количество символов в пути.
		SPACE COMPLEXITY:  O(n), так как в худшем случае мы можем хранить в стеке все символы пути.

		Method
		В этом решении мы используем стек для отслеживания пройденных директорий.
		Мы разделяем путь на части по символу "/", а затем проходим по каждой части. Если часть равна "..",
		мы удаляем последнюю директорию из стека (если стек не пуст). Если часть не равна "." и не пуста, мы добавляем ее в стек.
		В конце мы объединяем все элементы стека с символом "/" и возвращаем результат.
	*/
	// Инициализируем стек для хранения директорий.
	stack := make([]string, 0)
	// Разделяем путь на части по символу "/".
	parts := strings.Split(path, "/")

	// Проходим по каждой части пути.
	for _, part := range parts {
		// Если часть пути равна "..", и стек не пуст, удаляем последнюю директорию из стека.
		if part == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if part != "." && part != "" {
			// Если часть пути не равна "." и не пуста, добавляем ее в стек.
			stack = append(stack, part)
		}
	}

	// Объединяем все элементы стека с символом "/" и возвращаем результат.
	return "/" + strings.Join(stack, "/")
}

func main() {
	fmt.Println(simplifyPath("/home/"))          // Выводит: "/home"
	fmt.Println(simplifyPath("/../"))            // Выводит: "/"
	fmt.Println(simplifyPath("/home//foo/"))     // Выводит: "/home/foo"
	fmt.Println(simplifyPath("/a/./b/../../c/")) // Выводит: "/c"
}
