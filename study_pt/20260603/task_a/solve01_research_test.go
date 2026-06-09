package main

import (
	"fmt"
	"testing"
)

func TestResearch(t *testing.T) {
	res := prefixFunctionDebug("abracadabra")
	fmt.Println("Res: ", res) // [0 0 0 1 0 1 0 1 2 3 4]
}

// Time complexity: O(n) — каждый символ сравнивается не более n раз в сумме (k только растёт и откатывается)
// Space complexity: O(n) — храним массив pi длины n
func prefixFunctionDebug(s string) []int {
	n := len(s)
	pi := make([]int, n) // создаёт массив целых чисел длины n (сколько букв в слове), заполненный нулями.
	// В него потом будем записывать результаты префикс-функции для каждой позиции.

	for i := 1; i < n; i++ {
		// k — длина текущего совпадающего префикса
		k := pi[i-1]
		// Пока символы не совпадают, откатываемся
		fmt.Printf("s[i]=%c s[k]=%c \n", s[i], s[k])
		for k > 0 && s[i] != s[k] {
			k = pi[k-1]
		}

		if s[i] == s[k] {
			k++
		}

		pi[i] = k
	}

	return pi
}
