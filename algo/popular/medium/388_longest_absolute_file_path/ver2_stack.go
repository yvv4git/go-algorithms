package main

import (
	"strings"
)

// Функция lengthLongestPathV2 принимает строку input, представляющую иерархию файлов и директорий,
// и возвращает длину самого длинного абсолютного пути к файлу.
func lengthLongestPathV2(input string) int {
	/*
		METHOD:
		Используется стек для отслеживания длин путей до каждого уровня вложенности.
		Каждая строка разбирается на имя и уровень вложенности, после чего обновляется стек.
		Если встречается файл, вычисляется длина текущего пути и обновляется максимальная длина, если текущая длина больше.

		TIME COMPLEXITY:
		O(n), где n — длина входной строки. Мы проходим по каждому символу входной строки не более одного раза.

		SPACE COMPLEXITY:
		O(m), где m — количество строк после разделения. Это связано с тем, что мы храним стек для длин путей.
	*/
	// Разделяем входную строку по символу новой строки '\n'
	lines := strings.Split(input, "\n")

	// Создаем стек для хранения длины пути до каждого уровня вложенности
	stack := make([]int, len(lines)+1)

	// Инициализируем переменную для хранения максимальной длины пути
	maxLength := 0

	// Проходим по каждой строке
	for _, line := range lines {
		// Убираем символы табуляции для получения имени директории или файла
		name := strings.TrimLeft(line, "\t")

		// Определяем уровень вложенности по количеству символов табуляции
		depth := len(line) - len(name)

		// Текущая длина пути до текущего уровня
		currentLength := stack[depth] + len(name)

		// Если имя содержит точку, значит это файл
		if strings.Contains(name, ".") {
			// Обновляем максимальную длину, если текущая длина больше
			if currentLength > maxLength {
				maxLength = currentLength
			}
		} else {
			// Если это директория, обновляем длину пути до следующего уровня
			// Добавляем 1 для символа '/' между директориями
			stack[depth+1] = currentLength + 1
		}
	}

	return maxLength
}
