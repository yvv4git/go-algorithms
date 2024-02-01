package main

import (
	"strconv"
	"strings"
)

// Функция diffWaysToComputeV2 принимает строку-выражение и возвращает слайс целых чисел.
func diffWaysToComputeV2(expression string) []int {
	// Разбиваем выражение на числа и знаки операций, заменяя знаки на пробелы.
	numStr := strings.Fields(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(expression, "+", " "), "-", " "), "*", " "))
	numbers := make([]int, 0)
	for _, v := range numStr {
		n, _ := strconv.Atoi(v)
		numbers = append(numbers, n)
	}
	// Если чисел нет, возвращаем пустой слайс.
	if len(numbers) == 0 {
		return []int{}
	}
	// Создаем двумерный слайс для хранения результатов вычислений.
	dp := make([][][]int, len(numbers))
	for i := range dp {
		dp[i] = make([][]int, len(numbers))
	}
	// Заполняем главную диагональ слайса числами из выражения.
	for i := range dp {
		dp[i][i] = []int{numbers[i]}
	}
	// Создаем слайс для хранения знаков операций.
	sign := make([]rune, 0, len(numbers)-1)
	for _, v := range expression {
		switch v {
		case '+', '-', '*':
			sign = append(sign, v)
		}
	}
	// Вызываем функцию calc для вычисления всех возможных комбинаций значений.
	return calc(numbers, sign, dp, 0, len(numbers)-1)
}

// Функция calc вычисляет все возможные комбинации значений для выражения.
func calc(numbers []int, sign []rune, dp [][][]int, i, j int) []int {
	// Если результат уже вычислен, возвращаем его.
	if len(dp[i][j]) != 0 {
		return dp[i][j]
	}
	// Итеративно вычисляем значения для всех возможных комбинаций.
	for m := i; m < j; m++ {
		a := calc(numbers, sign, dp, i, m)
		b := calc(numbers, sign, dp, m+1, j)
		for _, va := range a {
			for _, vb := range b {
				// В зависимости от знака операции, выполняем соответствующую операцию.
				switch sign[m] {
				case '+':
					dp[i][j] = append(dp[i][j], va+vb)
				case '-':
					dp[i][j] = append(dp[i][j], va-vb)
				case '*':
					dp[i][j] = append(dp[i][j], va*vb)
				}
			}
		}
	}
	// Возвращаем результат вычислений.
	return dp[i][j]
}
