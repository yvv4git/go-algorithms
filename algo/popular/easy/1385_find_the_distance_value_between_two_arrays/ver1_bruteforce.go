package main

import (
	"fmt"
	"math"
)

// Функция findTheDistanceValue принимает два массива целых чисел arr1 и arr2,
// а также целое число d. Она возвращает количество элементов в arr1,
// которые не имеют пары в arr2, такой что абсолютная разница между элементами пары больше d.
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	/*
		METHOD: Brute Force

		TIME COMPLEXITY: O(n * m), где n - количество элементов в arr1, а m - количество элементов в arr2.
		Это обусловлено двумя вложенными циклами, которые проходят по всем элементам массивов.
		В худшем случае, каждый элемент arr1 будет сравниваться со всеми элементами arr2.

		SPACE COMPLEXITY: O(1)
		Временная сложность алгоритма составляет O(1), так как не используется дополнительная память,
		кроме нескольких переменных (count, valid, x, y).
	*/
	// Инициализируем счетчик подходящих элементов
	count := 0

	// Проходим по каждому элементу x в arr1
	for _, x := range arr1 {
		// Инициализируем флаг, который будет показывать, подходит ли текущий элемент x
		valid := true

		// Проходим по каждому элементу y в arr2
		for _, y := range arr2 {
			// Если абсолютная разница между x и y меньше или равна d,
			// то текущий элемент x не подходит
			if math.Abs(float64(x-y)) <= float64(d) {
				valid = false
				// Прерываем внутренний цикл, так как элемент x уже не подходит
				break
			}
		}

		// Если элемент x подходит (т.е. для всех y в arr2 выполняется условие |x - y| > d),
		// увеличиваем счетчик
		if valid {
			count++
		}
	}

	// Возвращаем количество подходящих элементов
	return count
}

func main() {
	// Пример использования функции
	arr1 := []int{4, 5, 8}
	arr2 := []int{10, 9, 1, 8}
	d := 2

	result := findTheDistanceValue(arr1, arr2, d)
	fmt.Println(result) // Вывод: 2
}
