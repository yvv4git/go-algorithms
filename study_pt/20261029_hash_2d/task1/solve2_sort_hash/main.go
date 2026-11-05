package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countDistinctSetsV2(sets [][]int) int {
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

		METHOD: Sort and hash (сортировка и хеш-таблица)
		Множество - это набор, где порядок элементов не важен.
		Каждое множество копируется и сортируется по возрастанию.
		Отсортированные числа записываются в строку key через запятую, например "1,2,3,".
		Такой ключ однозначно задаёт множество, потому что все числа неотрицательные.
		Хеш-таблица seen хранит уже встреченные ключи.
		Если ключ новый, он добавляется в seen, а одинаковые множества дают одинаковый ключ.
		Ответ - это количество ключей в seen после обработки всех множеств.

		TIME COMPLEXITY: O(N*K*log K)
		Сортировка каждого из N множеств длины K занимает O(K*log K).
		Построение ключей и вставка в хеш-таблицу занимают O(N*K) в среднем.

		SPACE COMPLEXITY: O(N*K)
		Хеш-таблица хранит N ключей, суммарная длина всех ключей равна O(N*K).
	*/
	seen := make(map[string]struct{}, len(sets))
	for _, row := range sets {
		cp := append([]int(nil), row...)
		sort.Ints(cp)

		var key strings.Builder
		for _, value := range cp {
			key.WriteString(strconv.Itoa(value))
			key.WriteByte(',')
		}
		seen[key.String()] = struct{}{}
	}
	return len(seen)
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

	fmt.Fprintln(writer, countDistinctSetsV2(sets))
}
