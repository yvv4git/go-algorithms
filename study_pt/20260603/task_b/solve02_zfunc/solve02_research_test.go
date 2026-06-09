package main

import (
	"fmt"
	"testing"
)

func TestResearch(t *testing.T) {
	s := "ababa"
	z := zFunctionDebug(s)
	fmt.Println("Z-function of 'ababa':", z) // [0 0 3 0 1]
}

// zFunctionDebug возвращает массив Z длины len(s).
//
// Z[i] — сколько символов подряд (начиная с позиции i) совпадает с началом строки.
//
// Пример: s = "ababa"
//   - Z[2]: с позиции 2 строка "aba", начало "aba" → совпало 3 → Z[2] = 3
//   - Z[1]: с позиции 1 строка "baba", начало "ab..." → не совпало → Z[1] = 0
//   - Z[3]: с позиции 3 строка "ba", начало "ab..." → не совпало → Z[3] = 0
//   - Z[0] = 0 всегда.
//
// Time complexity:  O(n) — каждый символ сравнивается не более одного раза
// Space complexity: O(n) — массив Z длины n
func zFunctionDebug(s string) []int {
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0

	for i := 1; i < n; i++ {
		// l (left) и r (right) — индексы в строке.
		// Они обозначают отрезок [l, r], внутри которого
		// каждый символ совпадает с началом строки:
		//   s[l]=s[0], s[l+1]=s[1], ..., s[r]=s[r-l].

		// Если i попал в этот отрезок — берём готовое z[i-l],
		// но не дальше конца отрезка (r-i+1).
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}

		// Сравниваем символы: s[z[i]] (из начала) и s[i+z[i]] (от i).
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}

		// Если новый отрезок [i, i+z[i]-1] правее r — запоминаем его.
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}
	return z
}

func TestResearchPeriodZ(t *testing.T) {
	res := periodZDebug("ababa")
	fmt.Println("res: ", res)
}

// periodZ возвращает минимальную длину периода строки s.
// Time complexity:  O(n) — вызов zFunction + проход по массиву
// Space complexity: O(n) — массив Z длины n
func periodZDebug(s string) int {
	n := len(s)
	z := zFunction(s)
	for i := 1; i < n; i++ {
		if i+z[i] == n {
			return i
		}
	}

	return n
}
