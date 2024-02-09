package main

import (
	"bytes"
	"strconv"
)

// Функция для получения k-й перестановки из n чисел
func getPermutationV2(n int, k int) string {
	/*
		METHOD: Math
		TIME COMPLEXITY: O(n^2), где n - количество чисел для перестановки. Это обусловлено двойным циклом, который проходит по всем числам и удаляет выбранное число из слайса.
		SPACE COMPLEXITY: O(n), так как мы храним слайсы для чисел и факториалов, а также строку для результата.
	*/
	// Создаем слайс для хранения доступных чисел
	items := make([]int, n)
	// Создаем слайс для хранения факториалов
	factorials := make([]int, n+1)
	// Базовый случай факториала
	factorials[0] = 1
	// Заполняем слайс чисел и вычисляем факториалы
	for i := 1; i <= n; i++ {
		items[i-1] = i
		factorials[i] = i * factorials[i-1]
	}

	// Буфер для формирования результирующей строки
	var buf bytes.Buffer
	// Пока есть числа для выбора
	for i := 0; i < n; i++ {
		// Вычисляем индекс текущего числа
		digitIdx := (k - 1) / factorials[n-1-i]
		// Добавляем в буфер выбранное число
		buf.WriteString(strconv.Itoa(items[digitIdx]))
		// Удаляем выбранное число из доступных
		items = append(items[:digitIdx], items[digitIdx+1:]...)
		// Обновляем k
		k -= digitIdx * factorials[n-1-i]
	}
	// Возвращаем результат в виде строки
	return buf.String()
}
