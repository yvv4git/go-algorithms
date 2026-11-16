package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func solveDeadlinesV2(deadlines []int) []int {
	/*
		TASK: Назначить задачам дни 1..N так, чтобы суммарный штраф за просрочку был минимальным.
		Дедлайн - это номер дня, до которого включительно задачу нужно решить.
		Штраф за задачу с дедлайном t, решённую в день d, равен max(0, d - t).
		Функция возвращает массив answer, где answer[i] - день решения задачи с номером i.
		Пример: deadlines = [2, 3, 1], N = 3.
		После сортировки по дедлайну задачи идут так: 3 (t=1), 1 (t=2), 2 (t=3), ответ [2, 3, 1], штраф 0.
		Случай: deadlines = [1], N = 1, ответ [1], штраф 0.
		Случай: deadlines = [1, 1, 1], N = 3, дни 1, 2, 3 дают штраф 0 + 1 + 2 = 3.
		Граница: deadlines = [1000000, 1000000], N = 2, обе задачи решаются без штрафа, ответ [1, 2].
		Плохой пример: deadlines = [2, 1], порядок [1, 2] даёт штраф 0 для задачи 2 и 1 для задачи 1, всего 1.
		Оптимальный порядок [2, 1] даёт штраф 0.

		METHOD: Greedy EDD (сортировка по дедлайну)
		EDD расшифровывается как Earliest Due Date, то есть самый ранний дедлайн первым.
		Дедлайн - это крайний день, когда задачу ещё можно решить без штрафа.
		Массив order содержит номера задач и сортируется по возрастанию deadlines.
		Задача на позиции k в массиве order получает день k+1, а результат пишется в answer по её номеру.
		Почему это оптимально: пусть задача с большим дедлайном стоит раньше задачи с меньшим.
		Обмен этих двух задач не увеличивает штраф, значит сортировка по дедлайну даёт оптимум.

		TIME COMPLEXITY: O(N log N)
		Сортировка N задач занимает O(N log N), раздача дней занимает O(N).

		SPACE COMPLEXITY: O(N)
		Массивы order и answer хранят по одному значению на задачу.
	*/
	n := len(deadlines)
	order := make([]int, n)
	for i := range order {
		order[i] = i
	}

	slices.SortStableFunc(order, func(a, b int) int {
		switch {
		case deadlines[a] < deadlines[b]:
			return -1
		case deadlines[a] > deadlines[b]:
			return 1
		default:
			return 0
		}
	})

	answer := make([]int, n)
	for day, problem := range order {
		answer[problem] = day + 1
	}
	return answer
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return
	}

	deadlines := make([]int, n)
	for i := range deadlines {
		if _, err := fmt.Fscan(reader, &deadlines[i]); err != nil {
			return
		}
	}

	answer := solveDeadlinesV2(deadlines)
	for i, day := range answer {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, day)
	}
	fmt.Fprintln(writer)
}
