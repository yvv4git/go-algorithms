//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func decompressRLElist(nums []int) []int {
	/*
		APPROACH: Bytes buffer
		1. Создаём буфер buf
		2. Проходим по nums с шагом 2 (i увеличивается на 2)
		3. Частота повторения freq = nums[i]
		4. Значение val = nums[i+1]
		5. Записываем val freq раз в буфер buf
		6. Преобразуем буфер buf в строку
		7. Разбиваем строку обратно в []int (если нужно)
		8. Возвращаем результирующий список

		TIME COMPLEXITY: O(n)
		- Мы проходим по nums один раз, выполняя константные операции на каждом шаге.

		SPACE COMPLEXITY: O(n)
		- Мы используем память для хранения буфера.
		- Размер буфера равен сумме частот повторений.
	*/
	var buf bytes.Buffer

	for i := 0; i < len(nums); i += 2 {
		freq, val := nums[i], nums[i+1]
		for j := 0; j < freq; j++ {
			buf.WriteString(strconv.Itoa(val))
			buf.WriteByte(' ') // Разделитель (можно убрать)
		}
	}

	// Разбиваем строку обратно в []int (если нужно)
	// (Здесь просто для примера, на практике лучше избегать конвертаций)
	parts := strings.Fields(buf.String())
	result := make([]int, len(parts))
	for i, s := range parts {
		num, _ := strconv.Atoi(s)
		result[i] = num
	}

	return result
}

func main() {
	// Пример 1
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(decompressRLElist(nums1)) // [2 4 4 4]

	// Пример 2
	nums2 := []int{2, 1, 3, 5}
	fmt.Println(decompressRLElist(nums2)) // [1 1 5 5 5]
}
