package main

import (
	"bufio"
	"fmt"
	"os"
)

type asker func(x, y int) bool

func recoverPermutation(n, k, q int, ask asker) []int {
	/*
		TASK: Восстановить скрытую перестановку pi чисел от 1 до N по ответам на сравнения.
		Перестановка - это список, где каждое число от 1 до N встречается ровно один раз.
		Запрос ask(x, y) возвращает true, если жюри утверждает, что pi(x) < pi(y).
		Жюри может соврать не более k раз за всё взаимодействие, а ложные ответы выбирает адаптивно.
		Нужно вернуть массив p длины N, где p[i-1] = pi(i).
		Пример: N = 3, k = 1, pi = [2, 3, 1].
		Сравнение (1,2): истина YES, pi(1)=2 < pi(2)=3.
		Сравнение (2,3): истина NO, pi(2)=3 > pi(3)=1, но жюри может один раз соврать и сказать YES.
		Сравнение (3,1): истина YES, pi(3)=1 < pi(1)=2.
		Итоговая перестановка: [2, 3, 1].
		Граница: k = 0, тогда каждый ответ правдивый и обычная сортировка восстанавливает pi.

		METHOD: Repeated majority sort (сортировка с повторными сравнениями большинством)
		Базовая идея - повторять один и тот же запрос фиксированное число раз и брать ответ большинства.
		Число повторов r выбирается как 2*k+1, но ограничивается бюджетом Q, чтобы не превысить лимит.
		При малом r большинство может ошибаться, поэтому это простой, но не всегда надёжный подход.
		Позиции 1..N сортируются сортировкой слиянием, где сравнение использует процедуру большинства.
		Сортировка слиянием делит массив пополам, сортирует части и сливает их по порядку.
		После сортировки позиций по значению pi восстанавливаем p, где p[pos-1] = значение pi(pos).

		TIME COMPLEXITY: O(r * N log N), где r - число повторов сравнения.
		Сортировка даёт O(N log N) сравнений, каждое сравнение делает до r запросов.

		SPACE COMPLEXITY: O(N).
		Храним массив позиций, временный буфер слияния и итоговую перестановку.
	*/
	queriesUsed := 0
	maxQueries := q
	if maxQueries < 0 {
		maxQueries = 0
	}

	repeats := 2*k + 1
	if repeats < 1 {
		repeats = 1
	}
	estComparisons := n
	log2 := 0
	for (1 << log2) < n {
		log2++
	}
	estComparisons = n * log2
	if estComparisons < 1 {
		estComparisons = 1
	}
	if repeats*estComparisons > maxQueries && maxQueries > 0 {
		repeats = maxQueries / estComparisons
		if repeats < 1 {
			repeats = 1
		}
	}

	less := func(x, y int) bool {
		yes := 0
		no := 0
		for i := 0; i < repeats; i++ {
			if queriesUsed >= maxQueries {
				break
			}
			queriesUsed++
			if ask(x, y) {
				yes++
			} else {
				no++
			}
			remaining := repeats - i - 1
			if yes > no+remaining {
				return true
			}
			if no > yes+remaining {
				return false
			}
		}
		if yes == no {
			return x < y
		}
		return yes > no
	}

	order := make([]int, n)
	for i := 0; i < n; i++ {
		order[i] = i + 1
	}

	var mergeSort func([]int) []int
	mergeSort = func(arr []int) []int {
		if len(arr) <= 1 {
			res := make([]int, len(arr))
			copy(res, arr)
			return res
		}
		mid := len(arr) / 2
		left := mergeSort(arr[:mid])
		right := mergeSort(arr[mid:])
		res := make([]int, 0, len(arr))
		i, j := 0, 0
		for i < len(left) && j < len(right) {
			if less(left[i], right[j]) {
				res = append(res, left[i])
				i++
			} else {
				res = append(res, right[j])
				j++
			}
		}
		res = append(res, left[i:]...)
		res = append(res, right[j:]...)
		return res
	}

	sortedPos := mergeSort(order)
	p := make([]int, n)
	for value, pos := range sortedPos {
		p[pos-1] = value + 1
	}
	return p
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k, q int
	if _, err := fmt.Fscan(reader, &n, &k, &q); err != nil {
		return
	}

	ask := func(x, y int) bool {
		fmt.Fprintf(writer, "? %d %d\n", x, y)
		writer.Flush()
		var resp string
		if _, err := fmt.Fscan(reader, &resp); err != nil {
			os.Exit(0)
		}
		if resp == "-1" {
			os.Exit(0)
		}
		return resp == "YES"
	}

	p := recoverPermutation(n, k, q, ask)
	fmt.Fprint(writer, "!")
	for _, v := range p {
		fmt.Fprintf(writer, " %d", v)
	}
	fmt.Fprintln(writer)
}
