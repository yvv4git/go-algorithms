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
	for i := range n {
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

func minCyclicShiftV2(x, y string) int {
	/*
		TASK: Проверить, является ли y циклическим сдвигом x вправо,
		и найти минимальный размер сдвига или -1, если сдвига нет.
		Циклический сдвиг вправо на k позиций - это перенос последних k символов
		в начало строки, а остальные символы сдвигаются вправо.
		Строку можно представить как кольцо символов, которое поворачивают на k позиций.
		Примеры для x = "abcde" (n = 5):
		k = 0: "abcde" - сдвига нет;
		k = 1: "eabcd" - последняя буква "e" уходит в начало;
		k = 2: "deabc" - последние "de" уходят в начало;
		k = 3: "cdeab" - последние "cde" уходят в начало;
		k = 5: "abcde" - полный оборот, снова исходная строка.
		Если y = "abced", то это не циклический сдвиг, ответ -1.

		METHOD: Rolling hash (полиномиальный хеш)
		Полиномиальный хеш - это числовой отпечаток строки: каждый символ
		умножается на степень основания P и всё складывается по модулю M.
		Одинаковые строки дают одинаковый хеш, разные - почти всегда разный.
		Префиксные хеши позволяют получить хеш любой подстроки за O(1).
		Строим два префиксных хеша для z = x + x с разными модулями.
		Затем сравниваем хеши окон длины n с хешем y.
		Идём по позициям pos от n вниз: первое совпадение даёт максимальную
		позицию, а значит минимальный сдвиг k = n - pos.

		TIME COMPLEXITY: O(n)
		- Построение префиксных хешей O(n).
		- Проверка n+1 окон по O(1) каждое.

		SPACE COMPLEXITY: O(n)
		- Массивы префиксных хешей и степеней для двух модулей.
	*/
	n := len(x)
	if len(y) != n {
		return -1
	}
	if n == 0 {
		return 0
	}

	z := x + x
	first := newRollingHash(z, hashMod1, hashBase)
	second := newRollingHash(z, hashMod2, hashBase)
	targetFirst := newRollingHash(y, hashMod1, hashBase)
	targetSecond := newRollingHash(y, hashMod2, hashBase)

	targetHash1 := targetFirst.substr(0, n)
	targetHash2 := targetSecond.substr(0, n)

	for pos := n; pos >= 0; pos-- {
		if first.substr(pos, n) == targetHash1 && second.substr(pos, n) == targetHash2 {
			return n - pos
		}
	}
	return -1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var x, y string
	if _, err := fmt.Fscan(in, &x, &y); err != nil {
		return
	}

	fmt.Fprintln(out, minCyclicShiftV2(x, y))
}
