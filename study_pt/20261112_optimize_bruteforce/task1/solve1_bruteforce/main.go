package main

import (
	"bufio"
	"fmt"
	"os"
)

func hasIndependentSet(vertices int, target int, edges [][2]int) bool {
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

		METHOD: Bruteforce (полный перебор)
		Перебираем все наборы ровно из S вершин и проверяем в каждом отсутствие рёбер.
		Набор храним в маске chosen: бит с номером v равен 1, если вершина v входит в набор.
		Рекурсивная функция searchBruteforce добавляет вершины по возрастанию номера, поэтому каждый набор перебирается ровно один раз.
		Для готового набора функция isIndependent проверяет все пары вершин: если у какой-то вершины есть сосед внутри набора, набор не подходит.
		Если найден хотя бы один независимый набор, сразу возвращаем true.
		Число наборов равно C(V, S), то есть числу сочетаний из V по S.

		TIME COMPLEXITY: O(C(V, S) * S^2)
		- Для каждого набора из S вершин проверяем до S^2 пар вершин.

		SPACE COMPLEXITY: O(V + E)
		- Храним граф в виде массива масок смежности и рекурсию глубины не больше S.
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

	return searchBruteforce(0, target, 0, adj)
}

func searchBruteforce(start int, left int, chosen uint64, adj []uint64) bool {
	if left == 0 {
		return isIndependent(chosen, adj)
	}
	n := len(adj)
	for v := start; v <= n-left; v++ {
		next := chosen | (uint64(1) << uint(v))
		if searchBruteforce(v+1, left-1, next, adj) {
			return true
		}
	}
	return false
}

func isIndependent(mask uint64, adj []uint64) bool {
	for v := 0; v < len(adj); v++ {
		if mask&(uint64(1)<<uint(v)) == 0 {
			continue
		}
		if adj[v]&mask != 0 {
			return false
		}
	}
	return true
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

	if hasIndependentSet(vertices, target, edges) {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}
