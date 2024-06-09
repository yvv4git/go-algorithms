package main

import (
	"strconv"
	"strings"
)

// Функция largestNumberV2 реализует алгоритм построения наибольшего числа из массива чисел.
// Она использует кучу (heap) для эффективного нахождения наибольшего числа.
func largestNumberV2(nums []int) string {
	/*
		METHOD: Heap
		TIME COMPLEXITY: O(n log n)
		SPACE COMPLEXITY: O(n), где n - количество чисел в массиве nums

		Time complexity
		1. Построение кучи: В цикле for _, num := range nums мы добавляем каждый элемент в кучу, что требует O(log n) времени для просеивания вверх.
		Таким образом, построение кучи требует O(n log n) времени.
		2. Извлечение элементов из кучи: В цикле for len(heap) > 0 мы извлекаем каждый элемент из кучи, что также требует O(log n) времени для просеивания вниз.
		Таким образом, извлечение всех элементов из кучи требует O(n log n) времени.

		largestNumberV2 использует кучу (heap) для построения наибольшего числа из массива чисел.
		Куча - это структура данных, которая позволяет быстро получать доступ к максимальному (или минимальному) элементу,
		а также добавлять новые элементы с логарифмической сложностью.
	*/
	// Инициализация кучи строк
	var heap []string

	// Перебираем каждое число в массиве
	for _, num := range nums {
		// Добавляем число в виде строки в кучу
		heap = append(heap, strconv.Itoa(num))
		// Поднимаем последний элемент кучи
		heapUp(heap, len(heap)-1)
	}

	// Инициализация строителя для результата
	var res strings.Builder

	// Пока куча не пуста
	for len(heap) > 0 {
		// Если первый элемент в куче равен "0" и результат равен "0", то прерываем цикл
		if heap[0] == "0" && res.String() == "0" {
			break
		}

		// Добавляем первый элемент кучи в результат
		res.WriteString(heap[0])

		// Перемещаем последний элемент кучи на место первого
		heap[0] = heap[len(heap)-1]
		// Удаляем последний элемент из кучи
		heap = heap[:len(heap)-1]
		// Просеиваем первый элемент кучи вниз
		heapDown(heap, 0, len(heap)-1)
	}

	// Возвращаем результат в виде строки
	return res.String()
}

// Функция isLarger сравнивает две строки и определяет, является ли первая строка больше второй.
// Она используется для определения порядка элементов в куче.
func isLarger(str1, str2 string) bool {
	// Если строки равны, то они не могут быть больше или меньше
	if str1 == str2 {
		return false
	}

	// Определяем длину наименьшей строки
	length := len(str1)
	if len(str2) < length {
		length = len(str2)
	}

	// Сравниваем символы строк до тех пор, пока не достигнем конца наименьшей строки
	i := 0
	for ; i < length; i++ {
		if str1[i] > str2[i] {
			return true
		} else if str1[i] < str2[i] {
			return false
		}
	}

	// Если достигнут конец первой строки, то она больше
	if i == len(str1) {
		return isLarger(str1, str2[i:])
	}

	// Иначе вторая строка больше
	return isLarger(str1[i:], str2)
}

// Функция heapDown просеивает элемент вниз в куче.
func heapDown(heap []string, p, limit int) {
	// Вычисляем индексы левого и правого потомков
	l, r := 2*p+1, 2*p+2
	// Инициализация индекса наибольшего элемента как текущего элемента
	larger := p

	// Если левый потомок существует и он больше текущего элемента, то обновляем индекс наибольшего элемента
	if l <= limit && isLarger(heap[l], heap[larger]) {
		larger = l
	}

	// Если правый потомок существует и он больше текущего элемента, то обновляем индекс наибольшего элемента
	if r <= limit && isLarger(heap[r], heap[larger]) {
		larger = r
	}

	// Если наибольший элемент не является текущим элементом, то меняем их местами и просеиваем наибольший элемент вниз
	if larger != p {
		heap[p], heap[larger] = heap[larger], heap[p]
		heapDown(heap, larger, limit)
	}
}

// Функция heapUp просеивает элемент вверх в куче.
func heapUp(heap []string, p int) {
	// Вычисляем индекс родителя
	parent := (p - 1) / 2

	// Если родитель существует и текущий элемент больше родителя, то меняем их местами и просеиваем родителя вверх
	if parent >= 0 && isLarger(heap[p], heap[parent]) {
		heap[parent], heap[p] = heap[p], heap[parent]
		heapUp(heap, parent)
	}
}