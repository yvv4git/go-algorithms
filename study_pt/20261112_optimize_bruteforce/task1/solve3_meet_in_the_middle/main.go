package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func hasIndependentSetV3(vertices int, target int, edges [][2]int) bool {
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

		METHOD: Meet-in-the-middle (встречный перебор)
		Делим вершины на левую половину 0..split-1 и правую половину split..V-1.
		Маска - это целое число, где бит с номером i равен 1, если вершина i входит в набор.
		Для правой половины функция best(mask) считает максимальный размер независимого множества среди вершин из mask.
		best работает по рекурсии: берём младшую вершину v из mask и выбираем максимум из двух вариантов.
		Не брать v: перейти к mask без бита v.
		Взять v: добавить 1 и перейти к mask без v и без всех соседей v.
		Ответы best сохраняются в массиве memo, поэтому каждая маска правой половины вычисляется не более одного раза.
		Затем перебираем все независимые наборы левой половины и для каждого считаем маску allowed разрешённых вершин справа.
		Если размер левого набора плюс best(allowed) не меньше S, то суммарно можно набрать S вершин, и ответ true.

		TIME COMPLEXITY: O(2^(V/2) + V + E)
		- Левых независимых наборов не больше 2^(V/2), правых масок тоже не больше 2^(V/2).

		SPACE COMPLEXITY: O(2^(V/2) + V + E)
		- Массив memo хранит по одному значению на каждую маску правой половины.
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

	split := vertices / 2
	rightCount := vertices - split

	cross := make([]uint64, split)
	for i := 0; i < split; i++ {
		rightNeighbors := adj[i] >> uint(split)
		for rest := rightNeighbors; rest != 0; rest &= rest - 1 {
			j := bits.TrailingZeros64(rest)
			cross[i] |= uint64(1) << uint(j)
		}
	}

	rightAdj := make([]uint64, rightCount)
	for j := 0; j < rightCount; j++ {
		rightAdj[j] = adj[split+j] >> uint(split)
	}

	allRight := (uint64(1) << uint(rightCount)) - 1
	memo := make([]int8, 1<<uint(rightCount))
	for i := range memo {
		memo[i] = -1
	}

	var best func(mask uint64) int
	best = func(mask uint64) int {
		if mask == 0 {
			return 0
		}
		if cached := memo[mask]; cached >= 0 {
			return int(cached)
		}
		v := bits.TrailingZeros64(mask)
		bit := uint64(1) << uint(v)
		without := best(mask &^ bit)
		with := 1 + best((mask&^bit)&^rightAdj[v])
		result := without
		if with > result {
			result = with
		}
		memo[mask] = int8(result)
		return result
	}

	var search func(available uint64, chosen int, allowedRight uint64) bool
	search = func(available uint64, chosen int, allowedRight uint64) bool {
		if chosen+bits.OnesCount64(available)+bits.OnesCount64(allowedRight) < target {
			return false
		}
		if chosen+best(allowedRight) >= target {
			return true
		}
		for available != 0 {
			v := bits.TrailingZeros64(available)
			bit := uint64(1) << uint(v)
			available &^= bit
			nextAvailable := available &^ adj[v]
			nextAllowed := allowedRight &^ cross[v]
			if search(nextAvailable, chosen+1, nextAllowed) {
				return true
			}
		}
		return false
	}

	leftAll := (uint64(1) << uint(split)) - 1
	return search(leftAll, 0, allRight)
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

	if hasIndependentSetV3(vertices, target, edges) {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}
