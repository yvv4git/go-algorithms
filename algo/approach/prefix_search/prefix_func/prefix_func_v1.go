package main

import "fmt"

func PrefixFunction(s string) []int {
	/*
		METHOD: Префикс-функция (π-функция)
		TIME COMPLEXITY: O(n), где n — длина строки
		SPACE COMPLEXITY: O(n)

		π[i] — длина наибольшего собственного префикса подстроки s[0..i],
		который одновременно является её суффиксом.

		Используется в KMP для оптимизации сдвига при несовпадении,
		а также для поиска периода строки и других задач.

		Алгоритм основан на наблюдении: если мы знаем π[i-1] = len,
		то для вычисления π[i] нужно сравнить s[i] с s[len].
		При совпадении π[i] = len+1, иначе — переходим к π[len-1] и повторяем.
	*/
	n := len(s)
	pi := make([]int, n)

	for i := 1; i < n; i++ {
		j := pi[i-1]

		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}

		if s[i] == s[j] {
			j++
		}

		pi[i] = j
	}

	return pi
}

func main() {
	/*
		Пример вычисления префикс-функции для строки.
		Префикс-функция является основой алгоритма KMP.
	*/
	strings := []string{
		"AAAA",
		"ABABAB",
		"ABCDABC",
		"ABACABAB",
	}

	for _, s := range strings {
		pi := PrefixFunction(s)

		fmt.Printf("Строка: %s\n", s)
		fmt.Printf("π:      %v\n\n", pi)
	}
}
