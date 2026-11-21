package main

import (
	"bufio"
	"fmt"
	"os"
)

type asker func(x, y int) bool

func recoverPermutationV2(n, k, q int, ask asker) []int {
	/*
		TASK: Восстановить скрытую перестановку pi чисел от 1 до N по интерактивным сравнениям.
		Перестановка - это список, где каждое число от 1 до N встречается ровно один раз.
		Функция ask(x, y) возвращает true, если жюри утверждает, что pi(x) < pi(y).
		Жюри может соврать не более k раз за всё взаимодействие и выбирает ложные ответы адаптивно.
		Нужно вернуть массив p длины N, где p[i-1] = pi(i).
		Пример: N = 3, k = 1, pi = [2, 3, 1].
		Запрос ? 1 2 истинен: YES. Запрос ? 3 1 истинен: YES.
		Запрос ? 2 3 истинен: NO, но жюри может соврать и ответить YES.
		Повторный запрос ? 2 3 даёт NO, и ложь израсходована, ответ восстанавливается как [2, 3, 1].
		Граница: k = 0, ложных ответов нет, достаточно обычных сравнений.

		METHOD: Adaptive fault-tolerant sort (адаптивная устойчивая к ошибкам сортировка)
		Сравнение повторяется, пока перевес ответов не превышает число оставшихся возможных лжей.
		Счётчик liesLeft хранит, сколько ложных ответов жюри ещё может сделать.
		Для пары (x, y) считаем yes и no. Если yes-no > liesLeft, истина YES; если no-yes > liesLeft, истина NO.
		Ответы меньшинства в завершённом сравнении ложные, поэтому liesLeft уменьшается на их число.
		Когда liesLeft становится 0, сравнения делаются одним запросом.
		Позиции сортируются сортировкой слиянием, а сравнение использует адаптивную процедуру.
		При исчерпании бюджета запросов Q возвращается лучший текущий порядок.

		TIME COMPLEXITY: O(N log N + R), где R - число повторных запросов.
		Сортировка даёт O(N log N) сравнений, повторы зависят от стратегии жюри.

		SPACE COMPLEXITY: O(N).
		Храним массив позиций, временный буфер слияния и итоговую перестановку.
	*/
	liesLeft := k
	if liesLeft < 0 {
		liesLeft = 0
	}
	queriesUsed := 0
	maxQueries := q
	if maxQueries < 0 {
		maxQueries = 0
	}

	less := func(x, y int) bool {
		yes := 0
		no := 0
		for {
			if yes > liesLeft {
				liesLeft -= no
				return true
			}
			if no > liesLeft {
				liesLeft -= yes
				return false
			}
			if queriesUsed >= maxQueries {
				if yes == no {
					return x < y
				}
				return yes > no
			}
			queriesUsed++
			if ask(x, y) {
				yes++
			} else {
				no++
			}
		}
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

	p := recoverPermutationV2(n, k, q, ask)
	fmt.Fprint(writer, "!")
	for _, v := range p {
		fmt.Fprintf(writer, " %d", v)
	}
	fmt.Fprintln(writer)
}
