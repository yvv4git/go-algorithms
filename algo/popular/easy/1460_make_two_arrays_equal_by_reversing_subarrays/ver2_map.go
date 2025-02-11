//go:build ignore

package main

import (
	"fmt"
)

// Функция для решения задачи с помощью хеш-таблиц
func canBeEqualByHashMap(target, arr []int) bool {
	/*
		METHOD: MAP / HASH-TABLE
		Пусть есть два массива: target и arr.
		Цель: найти такой порядок переворота подмассивов arr, что они будут равны target.
		Подмассивы arr могут быть переворачиваемыми только в том случае, если они содержат все элементы target.
		Поэтому мы можем составить хеш-таблицу, в которой будем хранить частоту элементов target и arr.
		Затем, для каждого элемента target, мы будем уменьшать частоту элемента в arr на 1.
		Если в каком-то месте arr уменьшается частота элемента target на 1, то это означает, что мы не можем переворачивать этот подмассив.
		Если же все элементы target имеют частоту 0 в arr, то все подмассивы arr могут быть переворачиваемыми.
		Таким образом, мы можем проверить, можно ли переворачивать подмассивы arr, чтобы они стали равными target.

		TIME COMPLEXITY: O(n), где n - длина массивов target и arr.
		SPACE COMPLEXITY: O(n), где n - длина массивов target и arr.
	*/
	// Создаем хеш-таблицы для хранения частоты элементов
	countTarget := make(map[int]int)
	countArr := make(map[int]int)

	// Заполняем хеш-таблицы для обоих массивов
	for _, num := range target {
		countTarget[num]++
	}
	for _, num := range arr {
		countArr[num]++
	}

	// Сравниваем хеш-таблицы
	for key, value := range countTarget {
		if countArr[key] != value {
			return false
		}
	}
	return true
}

func main() {
	// Пример данных
	target := []int{1, 2, 3, 4}
	arr := []int{2, 4, 3, 1}

	// Вызов функции для проверки
	if canBeEqualByHashMap(target, arr) {
		fmt.Println("Массивы могут быть равными после переворота подмассивов")
	} else {
		fmt.Println("Массивы не могут быть равными после переворота подмассивов")
	}
}
