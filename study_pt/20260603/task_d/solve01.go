// Гласно-согласный поиск (поиск с классами эквивалентности)
// Подход: Каждая буква заменяется на класс: '0' — гласная, '1' — согласная.
// После замены задача сводится к точному поиску подстроки, который
// решается алгоритмом Кнута — Морриса — Пратта (префикс-функция).
// Time complexity:  O(|T| + |S|) — преобразование + КМП.
// Space complexity: O(|T| + |S|) — хранение преобразованных строк.

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'y'
}

func class(c byte) byte {
	if isVowel(c) {
		return '0'
	}
	return '1'
}

// transform заменяет каждую букву её классом (0 — гласная, 1 — согласная).
func transform(s string) string {
	b := make([]byte, len(s))
	for i := range s {
		b[i] = class(s[i])
	}
	return string(b)
}

// prefixFunction возвращает массив π длины len(s) для алгоритма КМП.
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
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) < 2 {
		return
	}
	t := strings.TrimSpace(lines[0])
	s := strings.TrimSpace(lines[1])

	pt := transform(t)
	ps := transform(s)

	pi := prefixFunction(ps)

	var matches []string
	j := 0
	for i := 0; i < len(pt); i++ {
		for j > 0 && pt[i] != ps[j] {
			j = pi[j-1]
		}
		if pt[i] == ps[j] {
			j++
		}
		if j == len(ps) {
			matches = append(matches, fmt.Sprint(i-len(ps)+1))
			j = pi[j-1]
		}
	}

	if len(matches) == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(strings.Join(matches, " "))
	}
}
