package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func hasIndependentSetV2(vertices int, target int, edges [][2]int) bool {
	/*
		TASK: Проверить, есть ли в неориентированном графе независимое множество размера S.
		Независимое множество - это набор вершин, никакие две из которых не соединены ребром.
		Функция возвращает true, если такое множество существует, и false иначе.
		Пример: V = 5, S = 3, рёбра 1-2, 3-4, 4-5.
		Вершины 1, 3 и 5 не соединены между собой, поэтому ответ true.
		Пример: V = 5, S = 4 с теми же рёбрами, ответ false: больше трёх независимых вершин нет.
		Случай S = 1: подходит любая вершина, поэтому при V >= 1 ответ всегда true.
		Граница S = V: при отсутствии рёбер ответ true, а при наличии хотя бы одного ребра ответ false.
		Плохой пример: V = 3, S = 2, рёбра 1-2, 2-3 и 1-3, ответ false.

		METHOD: Backtracking with pruning (перебор с отсечениями)
		Строим независимое множество постепенно и храним маску available из вершин, которые ещё можно добавить.
		На каждом шаге выбираем доступную вершину v с наибольшим числом соседей внутри available.
		В первом варианте берём v: из available удаляем саму v и всех её соседей, а размер набора увеличиваем на 1.
		Во втором варианте не берём v: удаляем из available только саму v.
		Если chosen + количество доступных вершин меньше S, то даже взяв все доступные вершины мы не наберём S, поэтому ветку отбрасываем.
		Если chosen достигло S, возвращаем true.
		Выбор вершины с наибольшим числом соседей быстрее уменьшает available и ускоряет перебор.

		TIME COMPLEXITY: O(2^V)
		- В худшем случае каждая вершина независимо берётся или не берётся, а отсечения уменьшают число веток на практике.

		SPACE COMPLEXITY: O(V + E)
		- Храним маски смежности для всех вершин и рекурсию глубины не больше V.
	*/
	if target <= 0 {
		return true
	}
	if target > vertices {
		return false
	}

	adj := make([]uint64, vertices)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if a < 0 || a >= vertices || b < 0 || b >= vertices || a == b {
			continue
		}
		adj[a] |= uint64(1) << uint(b)
		adj[b] |= uint64(1) << uint(a)
	}

	all := (uint64(1) << uint(vertices)) - 1
	return searchBacktracking(all, 0, target, adj)
}

func searchBacktracking(available uint64, chosen int, target int, adj []uint64) bool {
	if chosen >= target {
		return true
	}
	if chosen+bits.OnesCount64(available) < target {
		return false
	}
	if available == 0 {
		return false
	}

	bestVertex := -1
	bestDegree := -1
	for rest := available; rest != 0; rest &= rest - 1 {
		v := bits.TrailingZeros64(rest)
		degree := bits.OnesCount64(adj[v] & available)
		if degree > bestDegree {
			bestDegree = degree
			bestVertex = v
		}
	}

	bit := uint64(1) << uint(bestVertex)
	if searchBacktracking(available&^(adj[bestVertex]|bit), chosen+1, target, adj) {
		return true
	}
	return searchBacktracking(available&^bit, chosen, target, adj)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var vertices, target int
	if _, err := fmt.Fscan(reader, &vertices, &target); err != nil {
		return
	}

	var edgeCount int
	if _, err := fmt.Fscan(reader, &edgeCount); err != nil {
		return
	}

	edges := make([][2]int, edgeCount)
	for i := 0; i < edgeCount; i++ {
		var a, b int
		if _, err := fmt.Fscan(reader, &a, &b); err != nil {
			return
		}
		edges[i] = [2]int{a - 1, b - 1}
	}

	if hasIndependentSetV2(vertices, target, edges) {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}
