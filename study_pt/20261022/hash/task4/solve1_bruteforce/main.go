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

func countDistinctSubstrings(s string) int {
	/*
		TASK: Посчитать количество различных подстрок строки S.
		Подстрока S[i..j] - это непрерывный кусок строки с i-го по j-й символ включительно.
		Одинаковые куски считаются один раз.
		Пример: S = "abc" - подстроки "a", "b", "c", "ab", "bc", "abc", ответ 6.
		Пример: S = "aaa" - различные подстроки "a", "aa", "aaa", ответ 3.
		Пример: S = "abab" - различные подстроки "a", "b", "ab", "ba", "aba", "bab", "abab", ответ 7.

		METHOD: Bruteforce (перебор подстрок и хеш-множество)
		Полный перебор - это проверка всех возможных вариантов по очереди.
		Полиномиальный хеш - это числовой отпечаток строки: символы умножаются
		на степени основания P и складываются по модулю M.
		Префиксные хеши позволяют получить хеш подстроки S[i..j] за O(1).
		Перебираем начало i и конец j, считаем пару хешей с двумя модулями
		и кладём в множество.
		Размер множества - количество различных подстрок.

		TIME COMPLEXITY: O(n^2)
		- Всего O(n^2) подстрок, хеш каждой считается за O(1).

		SPACE COMPLEXITY: O(n^2)
		- Множество хранит до O(n^2) пар хешей.
	*/
	n := len(s)
	if n == 0 {
		return 0
	}

	first := newRollingHash(s, hashMod1, hashBase)
	second := newRollingHash(s, hashMod2, hashBase)
	seen := make(map[[2]int64]struct{})

	for left := 0; left < n; left++ {
		for right := left + 1; right <= n; right++ {
			length := right - left
			seen[[2]int64{first.substr(left, length), second.substr(left, length)}] = struct{}{}
		}
	}
	return len(seen)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	if _, err := fmt.Fscan(in, &s); err != nil {
		return
	}

	fmt.Fprintln(out, countDistinctSubstrings(s))
}
