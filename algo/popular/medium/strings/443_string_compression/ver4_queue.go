package main

import (
	"strconv"
)

//type CharCount struct {
//	char  byte
//	count int
//}

// Функция compressV4 принимает на вход массив байтов chars и сжимает его с использованием очереди.
// Возвращает новую длину массива после сжатия.
func compressV4(chars []byte) int {
	/*
		METHOD: Queue
		TIME COMPLEXITY: O(n), где n - количество элементов в массиве chars. Это происходит потому, что мы проходим по каждому элементу в массиве только один раз.
		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем получить очередь размером, равной количеству уникальных символов в исходном массиве.
	*/
	// Если массив пустой, то сжатие не требуется, возвращаем 0.
	if len(chars) == 0 {
		return 0
	}

	// Инициализируем очередь с первым символом и количеством 1.
	queue := []CharCount{{chars[0], 1}}
	index := 0

	// Проходим по всем символам в массиве, начиная со второго.
	for i := 1; i < len(chars); i++ {
		// Если текущий символ равен символу в конце очереди, увеличиваем его количество.
		if chars[i] == queue[len(queue)-1].char {
			queue[len(queue)-1].count++
		} else {
			// Если символы разные, извлекаем первый элемент из очереди и записываем его в массив.
			chars[index] = queue[0].char
			index++
			// Если количество символов больше 1, записываем количество повторений символа.
			if queue[0].count > 1 {
				for _, digit := range []byte(strconv.Itoa(queue[0].count)) {
					chars[index] = digit
					index++
				}
			}
			// Удаляем первый элемент из очереди.
			queue = queue[1:]
			// Добавляем новый элемент в очередь с текущим символом и количеством 1.
			queue = append(queue, CharCount{chars[i], 1})
		}
	}

	// Обработка последнего символа в очереди.
	chars[index] = queue[0].char
	index++
	if queue[0].count > 1 {
		for _, digit := range []byte(strconv.Itoa(queue[0].count)) {
			chars[index] = digit
			index++
		}
	}

	// Возвращаем новую длину сжатого массива.
	return index
}

//func main() {
//	chars := []byte("aabbccc")
//	length := compressV2(chars)
//	fmt.Println(string(chars[:length])) // Вывод: a2b2c3
//}
