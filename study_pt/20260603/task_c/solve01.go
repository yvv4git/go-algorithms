// Z-функция
// Подход: Классический линейный алгоритм вычисления Z-функции.
// Z[i] — максимальная длина подстроки, начинающейся с позиции i
// и совпадающей с префиксом всей строки. Z[0] = 0.
// Используется [l, r] — отрезок последнего наибольшего совпадения
// для ускорения вычислений.
// Time complexity:  O(n) — каждый символ сравнивается не более одного раза.
// Space complexity: O(n) — массив Z длины n.

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// zFunction возвращает массив Z длины len(s).
func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}
	return z
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	data, _ := io.ReadAll(os.Stdin)
	s := strings.TrimSpace(string(data))

	z := zFunction(s)

	out := make([]string, len(z))
	for i, v := range z {
		out[i] = fmt.Sprint(v)
	}
	fmt.Println(strings.Join(out, " "))
}
