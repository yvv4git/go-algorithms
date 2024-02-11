package main

import (
	"strconv"
)

type CharCount struct {
	char  byte
	count int
}

// Функция compressV3 принимает на вход массив байтов chars и сжимает его с использованием стека.
// Возвращает новую длину массива после сжатия.
func compressV3(chars []byte) int {
	/*
		METHOD: Stack
		TIME COMPLEXITY: O(n), где n - количество элементов в массиве chars. Это происходит потому, что мы проходим по каждому элементу в массиве только один раз.
		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем получить стек размером, равным количеству уникальных символов в исходном массиве.
	*/
	// Если массив пустой, то сжатие не требуется, возвращаем 0.
	if len(chars) == 0 {
		return 0
	}

	// Инициализируем стек с первым символом и количеством 1.
	stack := []CharCount{{chars[0], 1}}

	// Проходим по всем символам в массиве, начиная со второго.
	for i := 1; i < len(chars); i++ {
		// Если текущий символ равен символу на вершине стека, увеличиваем его количество.
		if chars[i] == stack[len(stack)-1].char {
			stack[len(stack)-1].count++
		} else {
			// Если символы разные, добавляем новый элемент в стек с текущим символом и количеством 1.
			stack = append(stack, CharCount{chars[i], 1})
		}
	}

	// Индекс для записи в сжатый массив.
	index := 0
	// Проходим по стеку и записываем символы и их количества в сжатый массив.
	for _, cc := range stack {
		chars[index] = cc.char
		index++
		// Если количество символов больше 1, записываем количество повторений символа.
		if cc.count > 1 {
			for _, digit := range []byte(strconv.Itoa(cc.count)) {
				chars[index] = digit
				index++
			}
		}
	}

	// Возвращаем новую длину сжатого массива.
	return index
}

//func main() {
//	chars := []byte("aabbccc")
//	length := compress(chars)
//	fmt.Println(string(chars[:length])) // Вывод: a2b2c3
//}
