package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countIsomorphicSubtreePairs(n int, edges [][2]int, values []int) int {
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

		METHOD: Bruteforce (попарное сравнение сигнатур)
		Сигнатура поддерева - это строка, которая однозначно описывает его структуру и значения.
		Дерево обходится от корня, чтобы получить список детей для каждой вершины.
		Вершины обрабатываются в обратном порядке обхода: дети раньше родителей.
		Для вершины v сигнатуры детей сортируются и склеиваются, а спереди добавляется значение v.
		После этого все пары сигнатур sigs[i] и sigs[j] при i < j сравниваются напрямую.
		Если сигнатуры равны, поддеревья изоморфны и счётчик count увеличивается.

		TIME COMPLEXITY: O(n^3) в худшем случае
		В цепочке суммарная длина сигнатур равна O(n^2), а сравнение всех пар добавляет O(n^2) сравнений строк.

		SPACE COMPLEXITY: O(n^2) в худшем случае
		Массив sigs хранит строку для каждой вершины, и в цепочке их суммарная длина равна O(n^2).
	*/
	children, order := buildTree(n, edges)

	sigs := make([]string, n+1)
	for i := len(order) - 1; i >= 0; i-- {
		v := order[i]
		parts := make([]string, len(children[v]))
		for j, child := range children[v] {
			parts[j] = sigs[child]
		}
		sort.Strings(parts)

		var b strings.Builder
		b.WriteString(strconv.Itoa(values[v-1]))
		b.WriteByte('(')
		b.WriteString(strings.Join(parts, ","))
		b.WriteByte(')')
		sigs[v] = b.String()
	}

	count := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if sigs[i] == sigs[j] {
				count++
			}
		}
	}
	return count
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

	fmt.Fprintln(writer, countIsomorphicSubtreePairs(n, edges, values))
}
