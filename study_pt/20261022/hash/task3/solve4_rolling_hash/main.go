package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	hashMod1 int64 = 1_000_000_007
	hashMod2 int64 = 1_000_000_009
	hashBase int64 = 911_382_323
)

type rollingHash struct {
	mod  int64
	pref []int64
	pow  []int64
}

func newRollingHash(s string, mod, base int64) rollingHash {
	n := len(s)
	pref := make([]int64, n+1)
	pow := make([]int64, n+1)
	pow[0] = 1
	for i := 0; i < n; i++ {
		pref[i+1] = (pref[i]*base + int64(s[i])) % mod
		pow[i+1] = pow[i] * base % mod
	}
	return rollingHash{mod: mod, pref: pref, pow: pow}
}

func (h rollingHash) substr(left, length int) int64 {
	value := h.pref[left+length] - h.pref[left]*h.pow[length]%h.mod
	if value < 0 {
		value += h.mod
	}
	return value
}

func minBaseLengthV4(t string) int {
	/*
		TASK: Найти минимально возможную длину строки S, из повторений которой получен данный префикс t.
		Строка S писалась несколько раз подряд, после чего от результата взяли префикс t.
		Нужно вернуть минимальную длину S.
		Пример: t = "zzz" - это префикс "zzzz...", поэтому S = "z" и ответ 1.
		Пример: t = "bcabcab" - это префикс "bcabcabc...", поэтому S = "bca" и ответ 3.
		Пример: t = "ababab" - это префикс "abababab...", поэтому S = "ab" и ответ 2.
		Пример: t = "abc" - S может быть "abc", повторённым много раз, поэтому ответ 3.

		METHOD: Rolling hash (полиномиальный хеш)
		Полиномиальный хеш - это числовой отпечаток строки: каждый символ
		умножается на степень основания P и всё складывается по модулю M.
		Одинаковые строки дают одинаковый хеш, разные - почти всегда разный.
		Префиксные хеши позволяют получить хеш любой подстроки за O(1).
		Для каждого p сравниваем хеши подстрок t[:n-p] и t[p:].
		Их равенство означает, что t продолжается с периодом p.
		Первое подходящее p - ответ, если его нет, ответ равен n.

		TIME COMPLEXITY: O(n)
		- Построение префиксных хешей O(n), проверка n периодов по O(1).

		SPACE COMPLEXITY: O(n)
		- Массивы префиксных хешей и степеней для двух модулей.
	*/
	n := len(t)
	if n == 0 {
		return 0
	}

	first := newRollingHash(t, hashMod1, hashBase)
	second := newRollingHash(t, hashMod2, hashBase)

	for p := 1; p < n; p++ {
		length := n - p
		if first.substr(0, length) == first.substr(p, length) &&
			second.substr(0, length) == second.substr(p, length) {
			return p
		}
	}
	return n
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t string
	if _, err := fmt.Fscan(in, &t); err != nil {
		return
	}

	fmt.Fprintln(out, minBaseLengthV4(t))
}
