package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func longestCommonPrefix(a, b string) int {
	limit := len(a)
	if len(b) < limit {
		limit = len(b)
	}
	i := 0
	for i < limit && a[i] == b[i] {
		i++
	}
	return i
}

func countDistinctSubstringsV3(s string) int {
	/*
		TASK: Посчитать количество различных подстрок строки S.
		Подстрока S[i..j] - это непрерывный кусок строки с i-го по j-й символ включительно.
		Одинаковые куски считаются один раз.
		Пример: S = "abc" - подстроки "a", "b", "c", "ab", "bc", "abc", ответ 6.
		Пример: S = "aaa" - различные подстроки "a", "aa", "aaa", ответ 3.
		Пример: S = "abab" - различные подстроки "a", "b", "ab", "ba", "aba", "bab", "abab", ответ 7.

		METHOD: Suffix array and LCP (суффиксный массив и LCP)
		Суффиксный массив - это номера всех суффиксов строки в лексикографическом порядке.
		LCP двух суффиксов - это длина их наибольшего общего префикса.
		Всего подстрок с учётом повторов n(n+1)/2.
		Каждый общий префикс соседних суффиксов уже был посчитан раньше,
		поэтому его нужно вычесть.
		Ответ равен n(n+1)/2 минус сумма LCP всех соседних пар в суффиксном массиве.

		TIME COMPLEXITY: O(n^2 log n)
		- Сортировка суффиксов сравнивает строки длины O(n), LCP считается за O(n^2).

		SPACE COMPLEXITY: O(n)
		- Хранятся только суффиксный массив и счётчики.
	*/
	n := len(s)
	if n == 0 {
		return 0
	}

	suffixes := make([]int, n)
	for i := range suffixes {
		suffixes[i] = i
	}
	sort.Slice(suffixes, func(i, j int) bool {
		return s[suffixes[i]:] < s[suffixes[j]:]
	})

	total := n * (n + 1) / 2
	for i := 1; i < n; i++ {
		total -= longestCommonPrefix(s[suffixes[i-1]:], s[suffixes[i]:])
	}
	return total
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	if _, err := fmt.Fscan(in, &s); err != nil {
		return
	}

	fmt.Fprintln(out, countDistinctSubstringsV3(s))
}
