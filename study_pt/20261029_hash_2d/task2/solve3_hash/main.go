package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	seedOne uint64 = 0x9e3779b97f4a7c15
	seedTwo uint64 = 0xbf58476d1ce4e5b9
	baseOne uint64 = 0x100000001b3
	baseTwo uint64 = 0xcbf29ce484222325
)

type hashPair struct {
	h1 uint64
	h2 uint64
}

func countIsomorphicSubtreePairsV3(n int, edges [][2]int, values []int) int {
	/*
		TASK: Посчитать количество пар изоморфных корневых поддеревьев в дереве с корнем в вершине 1.
		Корневое поддерево вершины v - это сама вершина v и все её потомки.
		Изоморфизм - это взаимно однозначное соответствие вершин, сохраняющее значения, корень и рёбра.
		Порядок детей не важен, поэтому поддеревья считаются одинаковыми по структуре и значениям.
		Пример: n = 4, рёбра 1-2, 1-3, 3-4, значения [1, 2, 3, 2].
		Листья 2 и 4 имеют значение 2 и не имеют детей, поэтому образуют одну изоморфную пару, ответ 1.
		Случаи: n = 2 с равными значениями даёт ответ 1, с разными значениями - ответ 0.
		Граница: звезда с корнем 1 и n-1 листьями одного значения даёт (n-1)*(n-2)/2 пар.
		Если значения всех вершин различны, изоморфных пар нет, ответ 0.

		METHOD: Subtree hash (хеш поддерева)
		Хеш поддерева - это пара 64-битных чисел, которая описывает структуру и значения поддерева.
		Вершины обрабатываются в обратном порядке обхода: дети раньше родителей.
		Для вершины v пары хешей всех детей сортируются, чтобы порядок детей не влиял на результат.
		Начальные h1 и h2 получаются перемешиванием значения v функцией mix64 с разными seed.
		Затем каждый ребёнок добавляется по правилу h1 = h1*baseOne + childH1 и h2 = h2*baseTwo + childH2.
		Пары hashes[v] складываются в хеш-таблицу частот, и форма с частотой cnt даёт cnt*(cnt-1)/2 пар.

		TIME COMPLEXITY: O(n*log n) в среднем
		Сортировка списков детей по всем вершинам занимает O(n*log n), остальная обработка - O(n).

		SPACE COMPLEXITY: O(n)
		Массивы hashes и order и хеш-таблица частот хранят по O(n) элементов.
	*/
	children, order := buildTree(n, edges)

	hashes := make([]hashPair, n+1)
	for i := len(order) - 1; i >= 0; i-- {
		v := order[i]
		childHashes := make([]hashPair, len(children[v]))
		for j, child := range children[v] {
			childHashes[j] = hashes[child]
		}
		sort.Slice(childHashes, func(a, b int) bool {
			if childHashes[a].h1 != childHashes[b].h1 {
				return childHashes[a].h1 < childHashes[b].h1
			}
			return childHashes[a].h2 < childHashes[b].h2
		})

		h1 := mix64(uint64(values[v-1]) + seedOne)
		h2 := mix64(uint64(values[v-1]) + seedTwo)
		for _, childHash := range childHashes {
			h1 = h1*baseOne + childHash.h1
			h2 = h2*baseTwo + childHash.h2
		}
		hashes[v] = hashPair{h1: h1, h2: h2}
	}

	freq := make(map[hashPair]int)
	for v := 1; v <= n; v++ {
		freq[hashes[v]]++
	}

	total := 0
	for _, cnt := range freq {
		total += cnt * (cnt - 1) / 2
	}
	return total
}

func mix64(x uint64) uint64 {
	x += 0x9e3779b97f4a7c15
	x = (x ^ (x >> 30)) * 0xbf58476d1ce4e5b9
	x = (x ^ (x >> 27)) * 0x94d049bb133111eb
	return x ^ (x >> 31)
}

func buildTree(n int, edges [][2]int) ([][]int, []int) {
	adj := make([][]int, n+1)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	children := make([][]int, n+1)
	parent := make([]int, n+1)
	parent[1] = -1
	order := make([]int, 0, n)
	order = append(order, 1)
	for i := 0; i < len(order); i++ {
		v := order[i]
		for _, u := range adj[v] {
			if u == parent[v] {
				continue
			}
			parent[u] = v
			children[v] = append(children[v], u)
			order = append(order, u)
		}
	}
	return children, order
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	edges := make([][2]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}

	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &values[i])
	}

	fmt.Fprintln(writer, countIsomorphicSubtreePairsV3(n, edges, values))
}
