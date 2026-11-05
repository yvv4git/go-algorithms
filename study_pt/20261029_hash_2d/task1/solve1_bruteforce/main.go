package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func countDistinctSets(sets [][]int) int {
	/*
		TASK: Посчитать количество различных множеств среди N наборов по K чисел.
		Множество - это набор чисел, в котором порядок не важен и нет повторов.
		Например, наборы {1, 2, 3} и {3, 2, 1} задают одно и то же множество.
		Пример: N = 4, K = 3, наборы [1 2 3], [3 2 1], [11 12 13], [2 1 3].
		Первый, второй и четвёртый наборы дают множество {1, 2, 3}, третий даёт {11, 12, 13}.
		Ответ равен 2, потому что различных множеств два.
		Случаи по K: K = 1 - набор из одного числа; K = 5 - набор из пяти чисел.
		Граница: N = 1, K = 10^6 - всего одно множество, ответ 1.
		Пустой список наборов sets = [] не содержит множеств, ответ 0.

		METHOD: Bruteforce (полный перебор)
		Множество - это набор, где порядок элементов не важен.
		Сначала каждое множество копируется в sorted и сортируется по возрастанию.
		После сортировки одинаковые множества выглядят как одинаковые массивы.
		Каждый массив sorted[i] сравнивается со всеми предыдущими sorted[j], где j < i.
		Если найдено полное совпадение, множество повторное и ответ не увеличивается.
		Иначе счётчик count увеличивается, потому что такое множество встретилось впервые.

		TIME COMPLEXITY: O(N^2*K)
		Сортировка всех множеств занимает O(N*K*log K).
		Сравнение каждой пары отсортированных множеств занимает O(K), всего пар O(N^2).

		SPACE COMPLEXITY: O(N*K)
		Массив sorted хранит копии всех множеств суммарной длины N*K.
	*/
	sorted := make([][]int, len(sets))
	for i, row := range sets {
		cp := append([]int(nil), row...)
		sort.Ints(cp)
		sorted[i] = cp
	}

	count := 0
	for i := range sorted {
		seenBefore := false
		for j := range i {
			if equalSets(sorted[i], sorted[j]) {
				seenBefore = true
				break
			}
		}
		if !seenBefore {
			count++
		}
	}
	return count
}

func equalSets(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k int
	if _, err := fmt.Fscan(reader, &n, &k); err != nil {
		return
	}

	sets := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(reader, &row[j])
		}
		sets[i] = row
	}

	fmt.Fprintln(writer, countDistinctSets(sets))
}
