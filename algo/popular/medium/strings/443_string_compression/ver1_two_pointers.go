package main

import (
	"fmt"
	"strconv"
)

// Функция compress принимает на вход массив байтов chars и сжимает его.
// Возвращает новую длину массива после сжатия.
func compress(chars []byte) int {
	/*
		METHOD: Two pointers
		TIME COMPLEXITY: O(n), где n - количество элементов в массиве chars. Это происходит потому, что мы проходим по каждому элементу в массиве только один раз.
		SPACE COMPLEXITY: O(1), в целях экономии памяти, мы модифицируем исходный массив, поэтому пространственная сложность составляет O(1), если не учитывать входящий массив.
	*/
	// Если массив пустой, то сжатие не требуется, возвращаем 0.
	if len(chars) == 0 {
		return 0
	}

	// index - текущая позиция для записи в сжатый массив.
	// count - счетчик повторяющихся символов.
	index := 0
	count := 1

	// Проходим по всем символам в массиве.
	for i := 1; i < len(chars); i++ {
		// Если текущий символ равен предыдущему, увеличиваем счетчик.
		if chars[i] == chars[i-1] {
			count++
		} else {
			// Если символы разные, записываем предыдущий символ в сжатый массив.
			chars[index] = chars[i-1]
			index++
			// Если счетчик больше 1, то записываем количество повторений символа.
			if count > 1 {
				for _, c := range strconv.Itoa(count) {
					chars[index] = byte(c)
					index++
				}
			}
			// Сбрасываем счетчик.
			count = 1
		}
	}

	// Записываем последний символ в сжатый массив.
	chars[index] = chars[len(chars)-1]
	index++
	// Если счетчик больше 1, то записываем количество повторений символа.
	if count > 1 {
		for _, c := range strconv.Itoa(count) {
			chars[index] = byte(c)
			index++
		}
	}

	// Возвращаем новую длину сжатого массива.
	return index
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	length := compress(chars)
	fmt.Println(string(chars[:length]))
}
