package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countIsomorphicSubtreePairsV2(n int, edges [][2]int, values []int) int {
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

		METHOD: Canonical IDs (канонические идентификаторы)
		Каноническая форма поддерева - это его описание, одинаковое для всех изоморфных поддеревьев.
		Вершины обрабатываются в обратном порядке обхода: дети раньше родителей.
		Для вершины v идентификаторы детей ids[child] сортируются и превращаются в строковый ключ вместе со значением v.
		Хеш-таблица typeMap выдаёт каждому новому ключу уникальный целый идентификатор.
		Одинаковые поддеревья получают одинаковый идентификатор, потому что их ключи совпадают.
		После обработки всех вершин частоты идентификаторов считаются в freq, и форма с частотой cnt даёт cnt*(cnt-1)/2 пар.

		TIME COMPLEXITY: O(n*log n) в худшем случае
		Сортировка списков детей по всем вершинам занимает O(n*log n).
		Суммарная длина ключей равна O(n), потому что каждое ребро добавляет один идентификатор ребёнка.

		SPACE COMPLEXITY: O(n)
		Массивы ids и order, а также хеш-таблицы typeMap и freq хранят по O(n) элементов.
	*/
	children, order := buildTree(n, edges)

	ids := make([]int, n+1)
	typeMap := make(map[string]int)
	for i := len(order) - 1; i >= 0; i-- {
		v := order[i]
		childIDs := make([]int, len(children[v]))
		for j, child := range children[v] {
			childIDs[j] = ids[child]
		}
		sort.Ints(childIDs)

		var b strings.Builder
		b.WriteString(strconv.Itoa(values[v-1]))
		b.WriteByte('|')
		for _, id := range childIDs {
			b.WriteString(strconv.Itoa(id))
			b.WriteByte(',')
		}
		key := b.String()

		id, ok := typeMap[key]
		if !ok {
			id = len(typeMap)
			typeMap[key] = id
		}
		ids[v] = id
	}

	freq := make(map[int]int)
	for v := 1; v <= n; v++ {
		freq[ids[v]]++
	}

	total := 0
	for _, cnt := range freq {
		total += cnt * (cnt - 1) / 2
	}
	return total
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

	fmt.Fprintln(writer, countIsomorphicSubtreePairsV2(n, edges, values))
}
