package main

import "fmt"

type TrieNode struct {
	children map[byte]*TrieNode
	fail     *TrieNode
	output   []int
}

func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

func buildTrie(patterns []string) *TrieNode {
	root := newTrieNode()
	for idx, pat := range patterns {
		node := root
		for i := 0; i < len(pat); i++ {
			ch := pat[i]
			if node.children[ch] == nil {
				node.children[ch] = newTrieNode()
			}
			node = node.children[ch]
		}
		node.output = append(node.output, idx)
	}
	return root
}

func buildFailureLinks(root *TrieNode) {
	/*
		METHOD: Aho-Corasick failure links (BFS)
		TIME COMPLEXITY: O(total length of all patterns)
		SPACE COMPLEXITY: O(total length of all patterns)

		Для каждого узла строится fail-ссылка — указатель на самый длинный собственный
		суффикс (в бору), который совпадает с префиксом какого-либо паттерна.
		Это аналог LPS-массива из KMP, но для бора (trie).
	*/
	queue := make([]*TrieNode, 0)

	for _, node := range root.children {
		node.fail = root
		queue = append(queue, node)
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for ch, child := range current.children {
			failNode := current.fail
			for failNode != nil && failNode.children[ch] == nil {
				failNode = failNode.fail
			}
			if failNode == nil {
				child.fail = root
			} else {
				child.fail = failNode.children[ch]
				child.output = append(child.output, child.fail.output...)
			}
			queue = append(queue, child)
		}
	}
}

func AhoCorasickSearch(text string, patterns []string) map[int][]int {
	/*
		METHOD: Алгоритм Ахо-Корасик
		TIME COMPLEXITY: O(n + m + k), где n — длина text, m — суммарная длина паттернов, k — число вхождений
		SPACE COMPLEXITY: O(m)

		Мультипаттернный поиск. Строит бор (trie) со fail-ссылками, затем проходит
		по тексту за O(n), находя все вхождения всех паттернов одновременно.
	*/
	if len(patterns) == 0 {
		return nil
	}

	root := buildTrie(patterns)
	buildFailureLinks(root)

	result := make(map[int][]int)
	node := root

	for i := 0; i < len(text); i++ {
		ch := text[i]

		for node != root && node.children[ch] == nil {
			node = node.fail
		}
		if node.children[ch] != nil {
			node = node.children[ch]
		}

		for _, patIdx := range node.output {
			pos := i - len(patterns[patIdx]) + 1
			result[patIdx] = append(result[patIdx], pos)
		}
	}

	return result
}

func main() {
	/*
		Пример использования алгоритма Ахо-Корасик для поиска нескольких паттернов в тексте.
	*/
	text := "ushershe"
	patterns := []string{"he", "she", "his", "hers"}

	matches := AhoCorasickSearch(text, patterns)

	fmt.Printf("Текст:    %s\n", text)
	fmt.Printf("Паттерны: %v\n\n", patterns)
	for idx, pat := range patterns {
		if positions, ok := matches[idx]; ok {
			for _, pos := range positions {
				fmt.Printf("«%s» найден на позиции %d: %s\n",
					pat, pos,
					text[:pos]+"["+pat+"]"+text[pos+len(pat):])
			}
		} else {
			fmt.Printf("«%s» не найден\n", pat)
		}
	}
}
