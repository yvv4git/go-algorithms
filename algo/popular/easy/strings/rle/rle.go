package main

import (
	"strconv"
)

/*
Алгоритм RLE — это алгоритм сжатия данных, который поддерживается большинством форматов файлов растровых изображений: TIFF, BMP и PCX.
RLE подходит для сжатия любого типа данных независимо от его информационного содержимого,
но содержание данных влияет на коэффициент сжатия. Несмотря на то, что большинство алгоритмов RLE не могут обеспечить высокие коэффициенты сжатия более сложных методов,
данный инструмент легко реализовать и быстро выполнить, что делает его хорошей альтернативой.
*/

func rleString(in string) string {
	var result string
	var counter int
	var last rune

	if len(in) == 0 {
		return ""
	}

	for _, c := range in {
		counter += 1

		if (last != c) && counter == 1 {
			result += string(c)
			last = c
		} else if (last != c) && counter > 1 {
			result += strconv.Itoa(counter - 1)
			counter = 1
			result += string(c)
			last = c
		}
	}

	// Обработаем последний символ
	result += strconv.Itoa(counter)
	return result
}
