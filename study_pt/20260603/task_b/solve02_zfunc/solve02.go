// Основание строки (минимальный период) через Z-функцию
// Подход: Строка является префиксом своего многократного повторения.
// Минимальный период — наименьшее i (1..n-1) такое, что i + Z[i] = n,
// т.е. суффикс, начинающийся с i, совпадает с префиксом и покрывает
// остаток строки. Если такого i нет — ответ n (вся строка).
// Time complexity:  O(n) — линейный алгоритм Z-функции.
// Space complexity: O(n) — массив Z длины n.

package main

// zFunction возвращает массив Z длины len(s).
// Z[i] — длина нбольшего общего префикса s и s[i..].
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

// periodZ возвращает минимальную длину периода строки s.
func periodZ(s string) int {
	n := len(s)
	z := zFunction(s)
	for i := 1; i < n; i++ {
		if i+z[i] == n {
			return i
		}
	}
	return n
}
