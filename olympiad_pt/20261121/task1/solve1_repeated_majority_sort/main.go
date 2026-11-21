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
		Надёжное сравнение повторяет один и тот же запрос r = 2*k+1 раз.
		Если ответов YES больше, считаем pi(x) < pi(y), иначе pi(x) > pi(y).
		Жюри может соврать не больше k раз всего, поэтому абсолютное большинство из 2*k+1 ответов верно.
		Дальше сортируем позиции 1..N сортировкой слиянием, используя надёжное сравнение.
		Сортировка слиянием делит массив пополам, рекурсивно сортирует части и сливает их по порядку.
		После сортировки позиций по значению pi получаем порядок от меньшего значения к большему.
		По этому порядку восстанавливаем массив p, где p[pos-1] = значение pi(pos).

		TIME COMPLEXITY: O(r * N log N), где r = 2*k+1.
		Сортировка слиянием делает O(N log N) сравнений, каждое сравнение делает r запросов.

		SPACE COMPLEXITY: O(N).
		Храним массив позиций, временный буфер слияния и итоговую перестановку.
	*/
	repeats := max(2*k+1, 1)

	less := func(x, y int) bool {
		yes := 0
		no := 0
		for range repeats {
			if ask(x, y) {
				yes++
			} else {
				no++
			}
		}
		return yes > no
	}

	order := make([]int, n)
	for i := range n {
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
	_ = q

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
