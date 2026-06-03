// Неточное совпадение (поиск с допуском не более одной ошибки)
// Подход: Для каждой возможной позиции вхождения вычисляем длину
// совпавшего суффикса и префикса через Z-функцию. Прямая Z-функция
// (p + '#' + t) даёт LCP с началом p; обратная Z-функция на реверсивных
// строках — LCP с концом p. Если forward + backward >= m-1,
// то расхождение не более чем в одном символе.
// Time complexity:  O(|p| + |t|) — два запуска Z-функции.
// Space complexity: O(|p| + |t|) — массивы Z и обратная строка.

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

// reverse возвращает перевёрнутую копию строки.
func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	data, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) < 2 {
		return
	}
	p := strings.TrimSpace(lines[0])
	t := strings.TrimSpace(lines[1])

	m := len(p)
	n := len(t)

	if m > n {
		fmt.Println(0)
		return
	}

	zFwd := zFunction(p + "#" + t)
	zBwd := zFunction(reverse(p) + "#" + reverse(t))

	var matches []string
	for i := 0; i <= n-m; i++ {
		forward := zFwd[m+1+i]
		if forward >= m {
			matches = append(matches, fmt.Sprint(i+1))
			continue
		}
		revPos := n - i - m
		backward := zBwd[m+1+revPos]
		if forward+backward >= m-1 {
			matches = append(matches, fmt.Sprint(i+1))
		}
	}

	fmt.Println(len(matches))
	if len(matches) > 0 {
		fmt.Println(strings.Join(matches, " "))
	}
}
