package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	seedOne uint64 = 0x9e3779b97f4a7c15
	seedTwo uint64 = 0xbf58476d1ce4e5b9
)

func countDistinctSetsV3(sets [][]int) int {
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

		METHOD: Set hash (хеш множества, не зависящий от порядка)
		Для каждого значения value вычисляются две псевдослучайные величины mix64(value + seedOne) и mix64(value + seedTwo).
		Суммы этих величин по всем элементам множества не зависят от порядка элементов.
		Поэтому одинаковые множества дают одинаковую пару сумм h1 и h2.
		Пара [2]uint64{h1, h2} используется как ключ в хеш-таблице seen.
		Разные множества могут случайно дать одинаковую пару, вероятность такого события около 2^-128.
		Ответ - это количество различных пар в seen после обработки всех множеств.

		TIME COMPLEXITY: O(N*K) в среднем
		Каждый из N*K элементов обрабатывается за O(1), вставка в хеш-таблицу в среднем за O(1).

		SPACE COMPLEXITY: O(N)
		Хеш-таблица seen хранит не больше N пар по 16 байт.
	*/
	seen := make(map[[2]uint64]struct{}, len(sets))
	for _, row := range sets {
		var h1, h2 uint64
		for _, value := range row {
			v := uint64(value)
			h1 += mix64(v + seedOne)
			h2 += mix64(v + seedTwo)
		}
		seen[[2]uint64{h1, h2}] = struct{}{}
	}
	return len(seen)
}

func mix64(x uint64) uint64 {
	x += 0x9e3779b97f4a7c15
	x = (x ^ (x >> 30)) * 0xbf58476d1ce4e5b9
	x = (x ^ (x >> 27)) * 0x94d049bb133111eb
	return x ^ (x >> 31)
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

	fmt.Fprintln(writer, countDistinctSetsV3(sets))
}
