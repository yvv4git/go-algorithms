package main

import (
	"fmt"
	"math"
)

// Функция для подсчёта "хороших" троек
func countGoodTriplets(arr []int, a int, b int, c int) int {
	/*
		METHOD: Brute Force
		Алгоритм работает следующим образом:
		1. Инициализируется переменная count, которая будет использоваться для подсчета количества "хороших" троек.
		2. Вложенные циклы for перебирают все возможные тройки (i, j, k) в массиве arr.
		3. Внутри вложенных циклов проверяется условие |arr[i] - arr[j]| <= a.
		4. Если условие выполня
		5. Внутри вложенных циклов проверяется условие |arr[j] - arr[k]| <= b.
		6. Если условие выполняется, то проверяется условие |arr[i] - arr[k]| <= c.
		7. Если все условия выполняются, то увеличивается счетчик count.
		8. По завершении вложенных циклов возвращается значение count.
		9. Выводится результат на экран.

		Time complexity: O(n^3)
		- В худшем случае алгоритм проходит по всем возможным тройкам в массиве arr, что занимает O(n^3) времени, где n - длина массива arr.
		- Внутри вложенных циклов проверяется условие |arr[j] - arr[k]| <= b.
		- Внутри вложенных циклов проверяется условие |arr[i] - arr[k]| <= c.
		- Временная сложность алгоритма составляет O(n^3).
		- Внутри вложенных циклов проверяется условие |arr[j] - arr[k]| <= b.
		- Внутри вложенных циклов проверяется условие |arr[i] - arr[k]| <= c.


		SPACE COMPLEXITY: O(1)
		- Алгоритм использует только константное количество дополнительной памяти для переменных count, i, j, k, и arr.
		- Пространственная сложность алгоритма составляет O(1).
	*/
	n := len(arr) // Длина массива
	count := 0    // Счётчик "хороших" троек

	// Перебор всех возможных троек (i, j, k)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// Проверка условия |arr[i] - arr[j]| <= a
			if math.Abs(float64(arr[i]-arr[j])) <= float64(a) {
				for k := j + 1; k < n; k++ {
					// Проверка оставшихся условий
					if math.Abs(float64(arr[j]-arr[k])) <= float64(b) &&
						math.Abs(float64(arr[i]-arr[k])) <= float64(c) {
						count++ // Увеличиваем счётчик, если тройка "хорошая"
					}
				}
			}
		}
	}

	return count // Возвращаем количество "хороших" троек
}

func main() {
	// Пример входных данных
	arr := []int{3, 0, 1, 1, 9, 7}
	a, b, c := 7, 2, 3

	// Вызов функции и вывод результата
	result := countGoodTriplets(arr, a, b, c)
	fmt.Println(result) // Ожидаемый вывод: 4
}
