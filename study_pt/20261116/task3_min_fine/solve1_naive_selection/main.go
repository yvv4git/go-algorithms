package main

import (
	"bufio"
	"fmt"
	"os"
)

func solveDeadlines(deadlines []int) []int {
	/*
		TASK: Назначить задачам дни 1..N так, чтобы суммарный штраф за просрочку был минимальным.
		Дедлайн - это номер дня, до которого включительно задачу нужно решить.
		Штраф за задачу с дедлайном t, решённую в день d, равен max(0, d - t).
		Функция возвращает массив answer, где answer[i] - день решения задачи с номером i.
		Пример: deadlines = [2, 3, 1], N = 3.
		Оптимальное расписание: задача 3 в день 1, задача 1 в день 2, задача 2 в день 3, ответ [2, 3, 1], штраф 0.
		Случай: deadlines = [1], N = 1, ответ [1], штраф 0.
		Случай: deadlines = [1, 1, 1], N = 3, дни 1, 2, 3 дают штраф 0 + 1 + 2 = 3.
		Граница: deadlines = [1000000, 1000000], N = 2, обе задачи решаются без штрафа, ответ [1, 2].
		Плохой пример: deadlines = [2, 1], порядок [1, 2] даёт штраф 0 для задачи 2 и 1 для задачи 1, всего 1.
		Оптимальный порядок [2, 1] даёт штраф 0.

		METHOD: Naive selection (наивный выбор минимального дедлайна)
		Каждый день среди нерешённых задач выбирается задача с наименьшим дедлайном.
		Дедлайн - это крайний день, когда задачу ещё можно решить без штрафа.
		Массив solved хранит, решена ли задача, а массив answer - назначенный день для каждой задачи.
		Такой выбор оптимален: задачу с меньшим дедлайном всегда выгодно поставить раньше.
		Например, две задачи с дедлайнами 5 и 2 на днях 2 и 3 дают штраф 1 в порядке 5, 2 и штраф 0 в порядке 2, 5.

		TIME COMPLEXITY: O(N^2)
		Каждый из N дней просматривает до N нерешённых задач.

		SPACE COMPLEXITY: O(N)
		Массивы solved и answer хранят по одному значению на задачу.
	*/
	n := len(deadlines)
	answer := make([]int, n)
	solved := make([]bool, n)

	for day := 1; day <= n; day++ {
		best := -1
		for i := 0; i < n; i++ {
			if solved[i] {
				continue
			}
			if best == -1 || deadlines[i] < deadlines[best] {
				best = i
			}
		}
		answer[best] = day
		solved[best] = true
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

	answer := solveDeadlines(deadlines)
	for i, day := range answer {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, day)
	}
	fmt.Fprintln(writer)
}
