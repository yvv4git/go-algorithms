package main

import (
	"fmt"
	"strconv"
)

func selfDividingNumbers(left int, right int) []int {
	/*
		METHOD: Iterative
		Для решения этой задачи мы будем использовать следующий подход:
		1. Пройдемся по всем числам в заданном диапазоне.
		2. Для каждого числа проверим, является ли оно саморазделившимся.
		3. Если число саморазделившееся, добавим его в результирующий список.

		Таким образом, мы получим список всех саморазделившихся чисел в заданном диапазоне.

		TIME COMPLEXITY: O(n), где n - количество чисел в заданном диапазоне.
		Это связано с тем, что мы проходимся по каждому числу в диапазоне.

		SPACE COMPLEXITY: O(1), так как мы не используем дополнительное пространство, кроме списка для хранения результата.
	*/
	var result []int
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			result = append(result, i)
		}
	}
	return result
}

func isSelfDividing(num int) bool {
	strNum := strconv.Itoa(num)
	for _, digit := range strNum {
		digitInt, _ := strconv.Atoi(string(digit))
		if digitInt == 0 || num%digitInt != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
