package main

import (
	"bufio"
	"fmt"
	"os"
)

type trieNode struct {
	next map[byte]int
}

func countDistinctSubstringsV2(s string) int {
	/*
		TASK: Посчитать количество различных подстрок строки S.
		Подстрока S[i..j] - это непрерывный кусок строки с i-го по j-й символ включительно.
		Одинаковые куски считаются один раз.
		Пример: S = "abc" - подстроки "a", "b", "c", "ab", "bc", "abc", ответ 6.
		Пример: S = "aaa" - различные подстроки "a", "aa", "aaa", ответ 3.
		Пример: S = "abab" - различные подстроки "a", "b", "ab", "ba", "aba", "bab", "abab", ответ 7.

		METHOD: Suffix trie (суффиксный бор)
		Бор (trie) - это дерево, где путь от корня задаёт строку,
		а каждое ребро помечено одним символом.
		Вставляем в бор все суффиксы S[i:].
		Разные подстроки - это разные пути от корня, поэтому каждая подстрока
		встречается в боре ровно один раз.
		Число различных подстрок равно числу созданных вершин без корня.

		TIME COMPLEXITY: O(n^2)
		- Суммарная длина всех суффиксов равна n(n+1)/2.

		SPACE COMPLEXITY: O(n^2)
		- В худшем случае бор содержит O(n^2) вершин.
	*/
	n := len(s)
	if n == 0 {
		return 0
	}

	nodes := []trieNode{{}}
	distinct := 0
	for start := 0; start < n; start++ {
		current := 0
		for i := start; i < n; i++ {
			c := s[i]
			if nodes[current].next == nil {
				nodes[current].next = make(map[byte]int)
			}
			next, ok := nodes[current].next[c]
			if !ok {
				next = len(nodes)
				nodes[current].next[c] = next
				nodes = append(nodes, trieNode{})
				distinct++
			}
			current = next
		}
	}
	return distinct
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	if _, err := fmt.Fscan(in, &s); err != nil {
		return
	}

	fmt.Fprintln(out, countDistinctSubstringsV2(s))
}
