package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	firstMod   int64 = 1_000_000_007
	secondMod  int64 = 1_000_000_009
	firstBase  int64 = 131
	secondBase int64 = 137
)

type rollingHash struct {
	mod  int64
	pow  []int64
	pref []int64
}

func newRollingHash(s string, mod, base int64) *rollingHash {
	n := len(s)
	h := &rollingHash{
		mod:  mod,
		pow:  make([]int64, n+1),
		pref: make([]int64, n+1),
	}
	h.pow[0] = 1
	for i := 0; i < n; i++ {
		h.pow[i+1] = h.pow[i] * base % mod
		h.pref[i+1] = (h.pref[i]*base + int64(s[i])) % mod
	}
	return h
}

func (h *rollingHash) substr(left, length int) int64 {
	return (h.pref[left+length] - h.pref[left]*h.pow[length]%h.mod + h.mod) % h.mod
}

func suffixGreater(s string, a, b int, first, second *rollingHash) bool {
	if a == b {
		return false
	}

	n := len(s)
	limit := n - max(a, b)
	lo, hi := 0, limit
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if first.substr(a, mid) == first.substr(b, mid) &&
			second.substr(a, mid) == second.substr(b, mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	lcp := lo
	if lcp == n-a {
		return false
	}
	if lcp == n-b {
		return true
	}

	return s[a+lcp] > s[b+lcp]
}

func maxSubstringV2(s string) string {
	/*
		TASK: Найти лексикографически максимальную непрерывную подстроку строки S.
		Подстрока - это непрерывный кусок строки.
		Лексикографический порядок - это порядок как в словаре: символы сравниваются слева направо, и больше та строка, у которой на первом различии символ больше.
		Если одна строка является префиксом другой, то более длинная строка больше.
		Пример: S = "abacaba", максимальная подстрока "caba".
		Пример: S = "babb", максимальная подстрока "bb", потому что "bb" длиннее "b" и имеет "b" префиксом.
		Пример: S = "abc", максимальная подстрока "c", потому что "c" больше "bc" и "abc".
		Граница: S = "a", ответ "a"; S = "az", ответ "z".
		Пустая строка недопустима по условию, но для S = "" функция возвращает "".
		Контрпример: S = "ab", максимальная подстрока "b", а не "ab", потому что на первом различии "b" > "a".
		Функция возвращает саму максимальную подстроку, а не её позицию.

		METHOD: Rolling hash and binary search (хеши и бинпоиск)
		Суффикс S[i:] - это конец строки начиная с позиции i.
		Максимальная подстрока всегда является суффиксом, поэтому ищем максимальный суффикс.
		Полиномиальный хеш - это числовой отпечаток строки: символы умножаются на степени основания base и складываются по модулю mod.
		Префиксные хеши h.pref позволяют получить хеш подстроки S[left..left+length-1] за O(1) через h.substr(left, length).
		Двойной хеш - это два хеша с разными mod и base; совпадение обоих сильно снижает вероятность коллизии.

		Длиннейший общий префикс (LCP) двух суффиксов - это количество совпадающих первых символов.
		В suffixGreater бинарным поиском ищем максимальную длину mid, на которой хеши суффиксов a и b совпадают.
		После этого сравниваем первые различающиеся символы s[a+lcp] и s[b+lcp].
		Если один суффикс закончился, то более длинный суффикс больше, потому что короткий является его префиксом.

		Перебираем i от 1 до n-1, сравниваем суффикс i с текущим best через suffixGreater и обновляем best.
		В конце возвращаем s[best:].

		TIME COMPLEXITY: O(n log n)
		- Перебираем n суффиксов, каждое сравнение с бинарным поиском занимает O(log n).

		SPACE COMPLEXITY: O(n)
		- Два массива префиксных хешей и два массива степеней, по O(n) каждый.
	*/

	n := len(s)
	if n == 0 {
		return ""
	}

	first := newRollingHash(s, firstMod, firstBase)
	second := newRollingHash(s, secondMod, secondBase)

	best := 0
	for i := 1; i < n; i++ {
		if suffixGreater(s, i, best, first, second) {
			best = i
		}
	}

	return s[best:]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	scanner.Scan()
	s := strings.TrimSpace(scanner.Text())

	fmt.Println(maxSubstringV2(s))
}
