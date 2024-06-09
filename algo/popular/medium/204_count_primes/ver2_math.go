package main

import (
	"math"
)

// Функция countPrimes принимает целое число n и возвращает количество простых чисел, которые меньше n.
func countPrimesV2(n int) int {
	/*
		METHOD: Math / Решето Аткина

		TIME COMPLEXITY: O(n), где n - входное число. Это результат использования линейной сложности, которая позволяет эффективно перебирать числа и отмечать их как простые или составные.

		SPACE COMPLEXITY: O(n), так как мы используем дополнительный булевый массив размера n для хранения информации о простых числах.
	*/
	if n <= 2 {
		return 0
	}

	// Инициализируем массив для отметки чисел как простые или составные.
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	// Используем Решето Аткина для поиска простых чисел.
	sqrtN := int(math.Sqrt(float64(n)))
	for x := 1; x <= sqrtN; x++ {
		for y := 1; y <= sqrtN; y++ {
			// Используем формулу для генерации простых чисел с помощью Решета Аткина.
			num := 4*x*x + y*y
			if num <= n && (num%12 == 1 || num%12 == 5) {
				isPrime[num] = !isPrime[num]
			}
			num = 3*x*x + y*y
			if num <= n && num%12 == 7 {
				isPrime[num] = !isPrime[num]
			}
			num = 3*x*x - y*y
			if x > y && num <= n && num%12 == 11 {
				isPrime[num] = !isPrime[num]
			}
		}
	}

	// Отмечаем числа вида k*(k+2) как составные.
	for i := 5; i <= sqrtN; i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i * i {
				isPrime[j] = false
			}
		}
	}

	// Подсчитываем количество простых чисел.
	count := 2 // Подсчитываем 2 и 3 вручную, так как шаг цикла начинается с 5.
	for i := 5; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}

//func main() {
//	n := 10
//	fmt.Println(countPrimes(n)) // Выводит: 4
//}
