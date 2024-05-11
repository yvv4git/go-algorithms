package main

import "fmt"

// Функция для проверки, является ли строка B циклическим сдвигом строки A
func rotateStringV3(s string, goal string) bool {
	/*
		METHOD: KMP / Кнута-Морриса-Пратта
		Алгоритм Кнута-Морриса-Пратта (KMP) позволяет эффективно искать подстроку в строке,
		используя предварительно обработанные данные.

		TIME COMPLEXITY: O(n), где n - длина строки s.
		Это связано с тем, что алгоритм KMP проходит по строке s и goal ровно один раз,
		поэтому его временная сложность линейно зависит от длины входных данных.

		SPACE COMPLEXITY: O(n), так как мы храним префикс-функцию для строки goal, которая имеет длину n.
		Таким образом, пространственная сложность также линейно зависит от размера входных данных.
	*/
	// Проверяем, что длины строк равны
	if len(s) != len(goal) {
		return false
	}

	// Если строки равны, то они являются циклическим сдвигом друг друга
	if s == goal {
		return true
	}

	// Создаем новую строку, состоящую из s и s
	newA := s + s

	// Используем KMP для поиска подстроки goal в новой строке newA
	prefix := computePrefixFunction(goal)
	j := 0

	// Проходим по всем символам строки newA
	for i := 0; i < len(newA); i++ {
		// Пока j больше 0 и символы goal[j] и newA[i] не совпадают,
		// мы сдвигаем j на позицию, указанную в префикс-функции для предыдущего символа
		for j > 0 && goal[j] != newA[i] {
			j = prefix[j-1]
		}
		// Если символы goal[j] и newA[i] совпадают, увеличиваем j на 1
		if goal[j] == newA[i] {
			j++
		}
		// Если j равен длине строки goal, то мы нашли всю строку goal в newA,
		// что означает, что строка goal является циклическим сдвигом строки s,
		// и функция возвращает true
		if j == len(goal) {
			return true
		}
	}

	return false
}

// Функция для создания префикс-функции для строки pattern
func computePrefixFunction(pattern string) []int {
	// Получаем длину строки pattern
	n := len(pattern)
	// Создаем слайс для хранения префикс-функции, инициализируем его с длиной n
	prefix := make([]int, n)
	// Первый элемент префикс-функции всегда равен 0
	prefix[0] = 0
	// Переменная j используется для отслеживания длины максимального префикса, который также является суффиксом
	j := 0

	// Проходим по всем символам строки pattern, начиная со второго символа
	for i := 1; i < n; i++ {
		// Пока j больше 0 и символы pattern[i] и pattern[j] не совпадают,
		// мы сдвигаем j на позицию, указанную в префикс-функции для предыдущего символа
		for j > 0 && pattern[i] != pattern[j] {
			j = prefix[j-1]
		}
		// Если символы pattern[i] и pattern[j] совпадают, увеличиваем j на 1
		if pattern[i] == pattern[j] {
			j++
		}
		// Записываем длину максимального префикса, который также является суффиксом для подстроки pattern[:i+1]
		prefix[i] = j
	}

	// Возвращаем полученную префикс-функцию
	return prefix
}

func main() {
	A := "abcde"
	B := "cdeab"
	fmt.Println(rotateString(A, B)) // Вывод: true
}
