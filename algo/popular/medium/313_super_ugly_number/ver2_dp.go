package main

import "math"

func nthSuperUglyNumberV2(n int, primes []int) int {
	/*
		METHOD: Dynamic programming

		TIME COMPLEXITY: O(n⋅p), обусловлена структурой алгоритма и тем, как он обрабатывает каждое число от 2 до n.

		SPACE COMPLEXITY: O(n+p) связана с использованием двух массивов: одного размером n+1 для хранения супер-уродливых чисел и другого размером p для отслеживания индексов простых чисел.
	*/
	// Проверяем базовый случай: если n равно 1, то первое супер-уродливое число - это 1.
	if n == 1 {
		return 1
	}

	// Создаем массив dp для хранения супер-уродливых чисел. Размер массива n+1, так как индексация начинается с 1.
	dp := make([]int, n+1)
	// Определяем количество простых чисел в списке primes.
	p := len(primes)
	// Первое супер-уродливое число - это 1.
	dp[1] = 1
	// Создаем массив primeVal для хранения текущих индексов в массиве dp для каждого простого числа.
	primeVal := make([]int, p)
	// Инициализируем массив primeVal, устанавливая все значения в 1, так как первое супер-уродливое число уже известно.
	for i := 0; i < p; i++ {
		primeVal[i] = 1
	}

	// Начинаем заполнять массив dp супер-уродливыми числами, начиная с 2 до n.
	for i := 2; i <= n; i++ {
		// Инициализируем переменную minP максимальным возможным целым числом, чтобы найти минимальное значение.
		minP := math.MaxInt
		// Проходим по всем простым числам, чтобы найти следующее супер-уродливое число.
		for j := 0; j < p; j++ {
			// Вычисляем потенциальное следующее супер-уродливое число и обновляем minP, если нашли меньшее значение.
			minP = min(minP, dp[primeVal[j]]*primes[j])
		}

		// Сохраняем найденное минимальное значение в массив dp.
		dp[i] = minP

		// Проходим по всем простым числам и обновляем индексы в массиве primeVal, если текущее супер-уродливое число соответствует значению, вычисленному с помощью этого простого числа.
		for j := 0; j < p; j++ {
			if dp[i] == dp[primeVal[j]]*primes[j] {
				primeVal[j]++
			}
		}
	}

	// Возвращаем n-ое супер-уродливое число.
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
