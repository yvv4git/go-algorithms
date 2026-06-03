// Основание строки (минимальный период) через префикс-функцию
// Подход: Строка является префиксом своего многократного повторения.
// Минимальный период = n - π[n-1], где π — префикс-функция.
// Time complexity:  O(n) — линейный проход префикс-функции.
// Space complexity: O(n) — массив π длины n.

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// prefixFunction возвращает массив π длины len(s).
func prefixFunction(s string) []int {
	n := len(s)
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		k := pi[i-1]
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

func main() {
	data, _ := io.ReadAll(os.Stdin)
	s := strings.TrimSpace(string(data))

	n := len(s)
	pi := prefixFunction(s)
	period := n - pi[n-1]

	fmt.Println(period)
}
